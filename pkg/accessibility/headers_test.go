package accessibility

import (
	"testing"
)

func TestHeaderWithoutMain(t *testing.T) {
	type args struct {
		accessibilityChecks []AccessibilityCheck
		accessibilityCheck  AccessibilityCheck
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "When h1 is present.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					AccessibilityCheck{
						Element: "h1",
					},
					AccessibilityCheck{
						Element: "h2",
					},
					AccessibilityCheck{
						Element: "h3",
					},
				},
				accessibilityCheck: AccessibilityCheck{
					Element: "h2",
				},
			},
			want: false,
		},
		{
			name: "When do not have h1 headers.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					AccessibilityCheck{
						Element: "h2",
					},
					AccessibilityCheck{
						Element: "h2",
					},
					AccessibilityCheck{
						Element: "h3",
					},
				},
				accessibilityCheck: AccessibilityCheck{
					Element: "h4",
				},
			},
			want: true,
		},
		{
			name: "When h1 is not present and starts with h3.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					AccessibilityCheck{
						Element: "h3",
					},
					AccessibilityCheck{
						Element: "h2",
					},
					AccessibilityCheck{
						Element: "h4",
					},
				},
				accessibilityCheck: AccessibilityCheck{
					Element: "h2",
				},
			},
			want: true,
		},
		{
			name: "When h1 is added but starts with h4.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					AccessibilityCheck{
						Element: "h4",
					},
					AccessibilityCheck{
						Element: "h2",
					},
					AccessibilityCheck{
						Element: "h3",
					},
				},
				accessibilityCheck: AccessibilityCheck{
					Element: "h1",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeaderWithoutMain(tt.args.accessibilityChecks, tt.args.accessibilityCheck); got != tt.want {
				t.Errorf("HeaderWithoutMain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderUnavailable(t *testing.T) {
	type args struct {
		accessibilityChecks []AccessibilityCheck
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "When h1 is present.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					AccessibilityCheck{
						Element: "h1",
					},
					AccessibilityCheck{
						Element: "h2",
					},
					AccessibilityCheck{
						Element: "h3",
					},
				},
			},
			want: false,
		},
		{
			name: "When do not have h1 headers but have other headers.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					AccessibilityCheck{
						Element: "h2",
					},
					AccessibilityCheck{
						Element: "h2",
					},
					AccessibilityCheck{
						Element: "h3",
					},
				},
			},
			want: false,
		},
		{
			name: "When no headers is present and other elements is present.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					AccessibilityCheck{
						Element: "a",
					},
					AccessibilityCheck{
						Element: "input",
					},
					AccessibilityCheck{
						Element: "div",
					},
				},
			},
			want: true,
		},
		{
			name: "When accessibilitychecks is empty.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeaderUnavailable(tt.args.accessibilityChecks); got != tt.want {
				t.Errorf("HeaderUnavailable() = %v, want %v", got, tt.want)
			}
		})
	}
}
