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
					{
						Element: "h1",
					},
					{
						Element: "h2",
					},
					{
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
					{
						Element: "h2",
					},
					{
						Element: "h2",
					},
					{
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
					{
						Element: "h3",
					},
					{
						Element: "h2",
					},
					{
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
					{
						Element: "h4",
					},
					{
						Element: "h2",
					},
					{
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
					{
						Element: "h1",
					},
					{
						Element: "h2",
					},
					{
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
					{
						Element: "h2",
					},
					{
						Element: "h2",
					},
					{
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
					{
						Element: "a",
					},
					{
						Element: "input",
					},
					{
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

func TestHeaderMainCount(t *testing.T) {
	type args struct {
		accessibilityChecks []AccessibilityCheck
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "When h1 is present.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{
						Element: "h1",
					},
					{
						Element: "h2",
					},
					{
						Element: "h3",
					},
				},
			},
			want: 1,
		},

		{
			name: "When two h1 is present.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{
						Element: "h1",
					},
					{
						Element: "h2",
					},
					{
						Element: "h3",
					},
					{
						Element: "h1",
					},
				},
			},
			want: 2,
		},
		{
			name: "When do not have h1 headers but have other headers.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{
						Element: "h2",
					},
					{
						Element: "h2",
					},
					{
						Element: "h3",
					},
				},
			},
			want: 0,
		},
		{
			name: "When no headers is present and other elements is present.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{
						Element: "a",
					},
					{
						Element: "input",
					},
					{
						Element: "div",
					},
				},
			},
			want: 0,
		},
		{
			name: "When accessibilitychecks is empty.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeaderMainCount(tt.args.accessibilityChecks); got != tt.want {
				t.Errorf("HeaderMainCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderLevel(t *testing.T) {
	type args struct {
		accessibilityCheck AccessibilityCheck
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{

		{
			name: "When h1 is present.",
			args: args{
				accessibilityCheck: AccessibilityCheck{
					Element: "h1",
				},
			},
			want:  1,
			want1: true,
		},
		{
			name: "When h2 is present.",
			args: args{
				accessibilityCheck: AccessibilityCheck{
					Element: "h2",
				},
			},
			want:  2,
			want1: true,
		},
		{
			name: "When h3 is present.",
			args: args{
				accessibilityCheck: AccessibilityCheck{
					Element: "h3",
				},
			},
			want:  3,
			want1: true,
		},
		{
			name: "When h4 is present.",
			args: args{
				accessibilityCheck: AccessibilityCheck{
					Element: "h4",
				},
			},
			want:  4,
			want1: true,
		},
		{
			name: "When h5 is present.",
			args: args{
				accessibilityCheck: AccessibilityCheck{
					Element: "h5",
				},
			},
			want:  5,
			want1: true,
		},
		{
			name: "When h6 is present.",
			args: args{
				accessibilityCheck: AccessibilityCheck{
					Element: "h6",
				},
			},
			want:  6,
			want1: true,
		},
		{
			name: "When no headers is present and other elements is present.",
			args: args{
				accessibilityCheck: AccessibilityCheck{
					Element: "a",
				},
			},
			want:  0,
			want1: false,
		},
		{
			name: "When accessibilitychecks is empty.",
			args: args{
				accessibilityCheck: AccessibilityCheck{},
			},
			want:  0,
			want1: false,
		},
		{
			name: "When accessibilitychecks is nil.",
			args: args{
				accessibilityCheck: AccessibilityCheck{},
			},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := HeaderLevel(tt.args.accessibilityCheck)
			if got != tt.want {
				t.Errorf("HeaderLevel() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HeaderLevel() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestHeaderInvalidOrdenation(t *testing.T) {
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
					{
						Element: "h1",
					},
					{
						Element: "h2",
					},
					{
						Element: "h3",
					},
				},
			},
			want: false,
		},
		{
			name: "When h1 is not present and starts with h3.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{
						Element: "h3",
					},
					{
						Element: "h2",
					},
					{
						Element: "h4",
					},
				},
			},
			want: true,
		},
		{
			name: "When h1 is added but starts with h4.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{
						Element: "h4",
					},
					{
						Element: "h2",
					},
					{
						Element: "h3",
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
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeaderInvalidOrdenation(tt.args.accessibilityChecks); got != tt.want {
				t.Errorf("HeaderInvalidOrdenation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderInvalidHierarchy(t *testing.T) {
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
			name: "When is correct hierarchi.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{Element: "h1"},
					{Element: "h2"},
					{Element: "h3"},
				},
				accessibilityCheck: AccessibilityCheck{Element: "h4"},
			},
			want: false,
		},
		{
			name: "When starts with h2.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{Element: "h2"},
					{Element: "h2"},
					{Element: "h3"},
				},
				accessibilityCheck: AccessibilityCheck{Element: "h2"},
			},
			want: true,
		},
		{
			name: "When starts with h1 and jump to h3.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{Element: "h1"},
					{Element: "h3"},
					{Element: "h4"},
				},
				accessibilityCheck: AccessibilityCheck{Element: "h2"},
			},
			want: true,
		},
		{
			name: "When order is correct and adding h4 after h2.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{Element: "h1"},
					{Element: "h2"},
					{Element: "h2"}},
				accessibilityCheck: AccessibilityCheck{Element: "h4"}},
			want: true,
		},
		{
			name: "When correct hierarchy but repeating h3.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{Element: "h1"},
					{Element: "h2"},
					{Element: "h3"},
				},
				accessibilityCheck: AccessibilityCheck{Element: "h3"},
			},
			want: false,
		},
		{
			name: "When correct hierarchy but adding h2.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{Element: "h1"},
					{Element: "h2"},
					{Element: "h3"},
				},
				accessibilityCheck: AccessibilityCheck{Element: "h2"},
			},
			want: false,
		},
		{
			name: "When correct hierarchy but adding h1.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{Element: "h1"},
					{Element: "h2"},
					{Element: "h3"},
				},
				accessibilityCheck: AccessibilityCheck{Element: "h1"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeaderInvalidHierarchy(tt.args.accessibilityChecks, tt.args.accessibilityCheck); got != tt.want {
				t.Errorf("HeaderInvalidHierarchy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderCount(t *testing.T) {
	type args struct {
		accessibilityChecks []AccessibilityCheck
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "When 1 header is present.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{Element: "h1"},
					{Element: "a"},
					{Element: "script"},
				},
			},
			want: 1,
		},
		{
			name: "When two headers is present.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{Element: "h1"},
					{Element: "a"},
					{Element: "h3"},
					{Element: "table"}},
			},
			want: 2,
		},
		{
			name: "When do not have headers.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{Element: "td"},
					{Element: "th"},
					{Element: "script"}},
			},
			want: 0,
		},
		{
			name: "When three headers is present.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{
					{Element: "h1"},
					{Element: "h2"},
					{Element: "h3"}},
			},
			want: 3,
		},
		{
			name: "When accessibilitychecks is empty.",
			args: args{
				accessibilityChecks: []AccessibilityCheck{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeaderCount(tt.args.accessibilityChecks); got != tt.want {
				t.Errorf("HeaderCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderCheck(t *testing.T) {
	type args struct {
		accessibilityCheck AccessibilityCheck
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{name: "When h1.",
			args: args{
				accessibilityCheck: AccessibilityCheck{Element: "h1"},
			},
			want: true,
		},

		{name: "When h2.",
			args: args{
				accessibilityCheck: AccessibilityCheck{Element: "h2"},
			},
			want: true,
		},

		{name: "When h3.",
			args: args{
				accessibilityCheck: AccessibilityCheck{Element: "h3"},
			},
			want: true,
		},

		{name: "When h4.",
			args: args{
				accessibilityCheck: AccessibilityCheck{Element: "h4"},
			},
			want: true,
		},

		{name: "When h5 .",
			args: args{
				accessibilityCheck: AccessibilityCheck{Element: "h5"},
			},
			want: true,
		},

		{name: "When h6.",
			args: args{
				accessibilityCheck: AccessibilityCheck{Element: "h6"},
			},
			want: true,
		},

		{name: "When h7 return false.",
			args: args{
				accessibilityCheck: AccessibilityCheck{Element: "h7"},
			},
			want: false,
		},

		{name: "When h0 return false .", args: args{
			accessibilityCheck: AccessibilityCheck{Element: "h0"},
		},
			want: false,
		},

		{name: "When accessibilitychecks is nil.",
			args: args{
				accessibilityCheck: AccessibilityCheck{},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeaderCheck(tt.args.accessibilityCheck); got != tt.want {
				t.Errorf("HeaderCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}
