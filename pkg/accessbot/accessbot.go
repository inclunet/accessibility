package accessbot

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/inclunet/accessibility/pkg/checker"
	"gopkg.in/telebot.v3"
)

type AccessBot struct {
	Bot        *telebot.Bot
	ReportPath string
	Users      map[int64]*telebot.User
}

func (a *AccessBot) Handle() {
	a.Bot.Handle("/start", a.HelpAndler)
	a.Bot.Handle(telebot.OnText, a.MessageHandler)
	a.Bot.Start()
}

func (a *AccessBot) HelpAndler(c telebot.Context) error {
	a.SaveUsers(c.Sender())
	return c.Send("Seja bem-vindo! para testar a acessibilidade de uma página envie apenas a URL (Endereço ou Link) da página,  assim que eu terminar a avaliação  envio um relatório completo para você!")
}

func (a *AccessBot) LoadUsers() error {
	file, err := os.ReadFile(a.ReportPath + "/users.json")

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &a.Users)

	if err != nil {
		return err
	}
	return nil
}

func (a *AccessBot) MessageHandler(c telebot.Context) error {
	url := c.Text()
	if strings.HasPrefix(url, "http") {
		evaluator, err := checker.NewChecker(checker.AccessibilityChecker{
			Lang:       c.Sender().LanguageCode,
			ReportPath: a.ReportPath,
		})

		if err != nil {
			log.Println(err)
		}

		evaluator.GetDomainName(url)
		evaluator.AddCheckListItem(url, "pagina")
		evaluator.CheckAllList()
		evaluator.SaveAllReports()
		for _, accessibilityReport := range evaluator.Reports {
			reportFile := &telebot.Document{
				File:     telebot.FromDisk(accessibilityReport.HtmlReportPath),
				MIME:     "text/html",
				FileName: accessibilityReport.ReportFile + ".html",
				Caption:  "Segue a avaliação de acessibilidade da página " + accessibilityReport.Title,
			}
			err := c.Send(reportFile)

			if err != nil {
				log.Println(err)
			}

		}
	} else {
		err := c.Send("Você não enviou um link válido envie um link válido.")

		if err != nil {
			log.Println(err)
		}
	}

	a.SaveUsers(c.Sender())
	return nil
}

func (a *AccessBot) SaveUsers(user *telebot.User) error {
	_, ok := a.Users[user.ID]

	if !ok {
		a.Users[user.ID] = user

		file, err := os.Create(a.ReportPath + "users.json")

		if err != nil {
			return err
		}

		defer file.Close()

		jsonContent, err2 := json.Marshal(a.Users)

		if err2 != nil {
			return err
		}

		_, err = file.Write(jsonContent)

		if err != nil {
			return err
		}
	}

	return nil
}

func New(token string, reportPath string) (AccessBot, error) {
	newAccessBot := AccessBot{}

	config := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(config)

	if err != nil {
		return newAccessBot, err
	}

	newAccessBot.Bot = bot
	newAccessBot.ReportPath = reportPath
	newAccessBot.Users = make(map[int64]*telebot.User)
	err = newAccessBot.LoadUsers()

	if err != nil {
		return newAccessBot, err
	}

	return newAccessBot, nil
}
