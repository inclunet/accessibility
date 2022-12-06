# accessibility
A single command line tool for  accessibility check developed in golang


## How to run this program

After build, an executable file"main.exe" are available in   project directory and can by executed has following commands on your command line at the project folder:

### Linux

./main --url https://<your-url-here>/

### Windows

main.exe --url https://<your-url-here>/

### Aditional options to evaluate pages

* --url: URL to evaluate page ande generate a report file.
* --lang: A language to evaluate page. default is pt-br;
* --report: Outpute filename to generate a report file (Report files is only in html format and default filename is "report.html").

### Output file

The outpute file is a single html file with evaluation results. A summary result table is available on top of page followed by error results encountered after evaluation process.