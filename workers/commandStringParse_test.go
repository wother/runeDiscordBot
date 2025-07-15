package workers

import (
	"testing"
)

func TestParseMessage(t *testing.T) {
	// Test for help command
	helpResult := ParseMessage("!help")
	if helpResult["type"] != "text" {
		t.Error("Expected type 'text' for !help command")
	}

	// Test for single rune cast
	castOneResult := ParseMessage("!cast one")
	if castOneResult["type"] != "embed" {
		t.Error("Expected type 'embed' for !cast one command")
	}

	// Test for multiple rune cast
	castThreeResult := ParseMessage("!cast three")
	if castThreeResult["type"] != "runeArray" {
		t.Error("Expected type 'runeArray' for !cast three command")
	}

	// Test for info command with a valid rune name
	infoResult := ParseMessage("!info fehu")
	if infoResult["type"] != "info" {
		t.Error("Expected type 'info' for !info command with a valid rune name")
	}

	// Test for info command with "allrunes"
	infoAllResult := ParseMessage("!info allrunes")
	if infoAllResult["type"] != "allRunesLinks" {
		t.Error("Expected type 'allRunesLinks' for !info allrunes command")
	}

	// Test for uptime command
	uptimeResult := ParseMessage("!uptime")
	if uptimeResult["type"] != "text" {
		t.Error("Expected type 'text' for !uptime command")
	}

	// Test for a non-command message
	nonCommandResult := ParseMessage("not a command")
	if nonCommandResult != nil {
		t.Error("Expected nil for a non-command message")
	}
}

func Test_isCommandToBot(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid command",
			args: args{message: "!help"},
			want: true,
		},
		{
			name: "valid command with leading space",
			args: args{message: " !help"},
			want: true,
		},
		{
			name: "not a command",
			args: args{message: "hello"},
			want: false,
		},
		{
			name: "empty string",
			args: args{message: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCommandToBot(tt.args.message); got != tt.want {
				t.Errorf("isCommandToBot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseVerb(t *testing.T) {
	type args struct {
		inputStringArr []string
	}
	tests := []struct {
		name string
		args args
		expectType string
		expectContentCheck func(t *testing.T, content interface{})
	}{
		{
			name: "help command",
			args: args{inputStringArr: []string{"help"}},
			expectType: "text",
			expectContentCheck: func(t *testing.T, content interface{}) {
				expected := `The Rune Secrets Bot will draw runes for you from the Elder Futhark.

**Commands are:**

!help for this help text
!cast or !cast one for a single rune casting
!cast three for a three rune casting
!cast five for a five rune casting (careful...)
!info allrunes or names or all or list for a list of all the rune names.
!info [runeName] for information on a specific Rune.`
				if content != expected {
					t.Errorf("Expected help text, got %v", content)
				}
			},
		},
		{
			name: "? command",
			args: args{inputStringArr: []string{"?"}},
			expectType: "text",
			expectContentCheck: func(t *testing.T, content interface{}) {
				expected := `The Rune Secrets Bot will draw runes for you from the Elder Futhark.

**Commands are:**

!help for this help text
!cast or !cast one for a single rune casting
!cast three for a three rune casting
!cast five for a five rune casting (careful...)
!info allrunes or names or all or list for a list of all the rune names.
!info [runeName] for information on a specific Rune.`
				if content != expected {
					t.Errorf("Expected help text, got %v", content)
				}
			},
		},
		{
			name: "info command with no subject",
			args: args{inputStringArr: []string{"info"}},
			expectType: "text",
			expectContentCheck: func(t *testing.T, content interface{}) {
				expected := `The Rune Secrets Bot will draw runes for you from the Elder Futhark.

**Commands are:**

!help for this help text
!cast or !cast one for a single rune casting
!cast three for a three rune casting
!cast five for a five rune casting (careful...)
!info allrunes or names or all or list for a list of all the rune names.
!info [runeName] for information on a specific Rune.`
				if content != expected {
					t.Errorf("Expected help text, got %v", content)
				}
			},
		},
		{
			name: "cast one command",
			args: args{inputStringArr: []string{"cast", "one"}},
			expectType: "embed",
			expectContentCheck: func(t *testing.T, content interface{}) {
				runeObj, ok := content.(Rune)
				if !ok {
					t.Errorf("Expected content to be of type Rune, got %T", content)
				}
				if runeObj.Name == "" {
					t.Error("Expected rune name to not be empty")
				}
				if runeObj.ImgFile != genImgPath(runeObj.Name) {
					t.Errorf("Expected ImgFile to be %s, got %s", genImgPath(runeObj.Name), runeObj.ImgFile)
				}
			},
		},
		{
			name: "cast three command",
			args: args{inputStringArr: []string{"cast", "three"}},
			expectType: "runeArray",
			expectContentCheck: func(t *testing.T, content interface{}) {
				runeArray, ok := content.([]Rune)
				if !ok {
					t.Errorf("Expected content to be of type []Rune, got %T", content)
				}
				if len(runeArray) != 3 {
					t.Errorf("Expected 3 runes, got %d", len(runeArray))
				}
				for _, runeObj := range runeArray {
					if runeObj.Name == "" {
						t.Error("Expected rune name to not be empty")
					}
					if runeObj.ImgFile != genImgPath(runeObj.Name) {
						t.Errorf("Expected ImgFile to be %s, got %s", genImgPath(runeObj.Name), runeObj.ImgFile)
					}
				}
			},
		},
		{
			name: "cast command with no number",
			args: args{inputStringArr: []string{"cast"}},
			expectType: "embed",
			expectContentCheck: func(t *testing.T, content interface{}) {
				runeObj, ok := content.(Rune)
				if !ok {
					t.Errorf("Expected content to be of type Rune, got %T", content)
				}
				if runeObj.Name == "" {
					t.Error("Expected rune name to not be empty")
				}
				if runeObj.ImgFile != genImgPath(runeObj.Name) {
					t.Errorf("Expected ImgFile to be %s, got %s", genImgPath(runeObj.Name), runeObj.ImgFile)
				}
			},
		},
		{
			name: "info allrunes command",
			args: args{inputStringArr: []string{"info", "allrunes"}},
			expectType: "allRunesLinks",
			expectContentCheck: func(t *testing.T, content interface{}) {
				runeNames, ok := content.([]string)
				if !ok {
					t.Errorf("Expected content to be of type []string, got %T", content)
				}
				if len(runeNames) != len(futharkNames) {
					t.Errorf("Expected %d rune names, got %d", len(futharkNames), len(runeNames))
				}
			},
		},
		{
			name: "info fehu command",
			args: args{inputStringArr: []string{"info", "fehu"}},
			expectType: "info",
			expectContentCheck: func(t *testing.T, content interface{}) {
				runeObj, ok := content.(Rune)
				if !ok {
					t.Errorf("Expected content to be of type Rune, got %T", content)
				}
				if runeObj.Name != "fehu" {
					t.Errorf("Expected rune name to be fehu, got %s", runeObj.Name)
				}
				if runeObj.ImgFile != genImgPath(runeObj.Name) {
					t.Errorf("Expected ImgFile to be %s, got %s", genImgPath(runeObj.Name), runeObj.ImgFile)
				}
			},
		},
		{
			name: "uptime command",
			args: args{inputStringArr: []string{"uptime"}},
			expectType: "text",
			expectContentCheck: func(t *testing.T, content interface{}) {
				expected := "All you need to know is I am online. Fer realsies."
				if content != expected {
					t.Errorf("Expected uptime text, got %v", content)
				}
			},
		},
		{
					{
			name: "unknown command",
			args: args{inputStringArr: []string{"unknown"}},
			expectType: "",
			expectContentCheck: func(t *testing.T, content interface{}) {
				if content == nil || len(content.(map[string]interface{})) != 0 {
					t.Errorf("Expected empty map content for unknown command, got %v", content)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseVerb(tt.args.inputStringArr)

			if tt.expectType == "" {
				if got != nil && len(got) != 0 { // For unknown commands, parseVerb returns an empty map
					t.Errorf("parseVerb() = %v, want empty map", got)
				}
			} else {
				if got["type"] != tt.expectType {
					t.Errorf("parseVerb() type = %v, want %v", got["type"], tt.expectType)
				}
				if tt.expectContentCheck != nil {
					tt.expectContentCheck(t, got["content"])
				}
			}
		})
	}
}

func Test_getRune(t *testing.T) {
	type args struct {
		inputString string
		infoBoolean bool
	}
	tests := []struct {
		name string
		args args
		expectType string
		expectContentCheck func(t *testing.T, content interface{})
	}{
		{
			name: "single rune from number string",
			args: args{inputString: "one", infoBoolean: false},
			expectType: "embed",
			expectContentCheck: func(t *testing.T, content interface{}) {
				runeObj, ok := content.(Rune)
				if !ok {
					t.Errorf("Expected content to be of type Rune, got %T", content)
				}
				if runeObj.Name == "" {
					t.Error("Expected rune name to not be empty")
				}
				if runeObj.ImgFile != genImgPath(runeObj.Name) {
					t.Errorf("Expected ImgFile to be %s, got %s", genImgPath(runeObj.Name), runeObj.ImgFile)
				}
			},
		},
		{
			name: "multiple runes from number string",
			args: args{inputString: "three", infoBoolean: false},
			expectType: "runeArray",
			expectContentCheck: func(t *testing.T, content interface{}) {
				runeArray, ok := content.([]Rune)
				if !ok {
					t.Errorf("Expected content to be of type []Rune, got %T", content)
				}
				if len(runeArray) != 3 {
					t.Errorf("Expected 3 runes, got %d", len(runeArray))
				}
				for _, runeObj := range runeArray {
					if runeObj.Name == "" {
						t.Error("Expected rune name to not be empty")
					}
					if runeObj.ImgFile != genImgPath(runeObj.Name) {
						t.Errorf("Expected ImgFile to be %s, got %s", genImgPath(runeObj.Name), runeObj.ImgFile)
					}
				}
			},
		},
		{
			name: "valid rune name, infoBoolean false",
			args: args{inputString: "fehu", infoBoolean: false},
			expectType: "embed",
			expectContentCheck: func(t *testing.T, content interface{}) {
				runeObj, ok := content.(Rune)
				if !ok {
					t.Errorf("Expected content to be of type Rune, got %T", content)
				}
				if runeObj.Name != "fehu" {
					t.Errorf("Expected rune name to be fehu, got %s", runeObj.Name)
				}
				if runeObj.ImgFile != genImgPath(runeObj.Name) {
					t.Errorf("Expected ImgFile to be %s, got %s", genImgPath(runeObj.Name), runeObj.ImgFile)
				}
			},
		},
		{
			{
			name: "valid rune name, infoBoolean true",
			args: args{inputString: "fehu", infoBoolean: true},
			expectType: "",
			expectContentCheck: func(t *testing.T, content interface{}) {
				if content == nil || len(content.(map[string]interface{})) != 0 {
					t.Errorf("Expected empty map content for infoBoolean true, got %v", content)
				}
			},
		},
		{
			name: "list command",
			args: args{inputString: "allrunes", infoBoolean: false},
			expectType: "allRunesLinks",
			expectContentCheck: func(t *testing.T, content interface{}) {
				runeNames, ok := content.([]string)
				if !ok {
					t.Errorf("Expected content to be of type []string, got %T", content)
				}
				if len(runeNames) != len(futharkNames) {
					t.Errorf("Expected %d rune names, got %d", len(futharkNames), len(runeNames))
				}
			},
		},
		{
			name: "invalid input",
			args: args{inputString: "invalid", infoBoolean: false},
			expectType: "",
			expectContentCheck: func(t *testing.T, content interface{}) {
				if content == nil || len(content.(map[string]interface{})) != 0 {
					t.Errorf("Expected empty map content for invalid input, got %v", content)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRune(tt.args.inputString, tt.args.infoBoolean)

			if tt.expectType == "" {
				if got != nil && len(got) != 0 { // For unknown commands, getRune returns an empty map
					t.Errorf("getRune() = %v, want empty map", got)
				}
			} else {
				if got["type"] != tt.expectType {
					t.Errorf("getRune() type = %v, want %v", got["type"], tt.expectType)
				}
				if tt.expectContentCheck != nil {
					tt.expectContentCheck(t, got["content"])
				}
			}
		})
	}
}

func Test_listCommand(t *testing.T) {
	type args struct {
		listTestString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "allrunes",
			args: args{listTestString: "allrunes"},
			want: true,
		},
		{
			name: "names",
			args: args{listTestString: "names"},
			want: true,
		},
		{
			name: "all",
			args: args{listTestString: "all"},
			want: true,
		},
		{
			name: "list",
			args: args{listTestString: "list"},
			want: true,
		},
		{
			name: "not a list command",
			args: args{listTestString: "fehu"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listCommand(tt.args.listTestString); got != tt.want {
				t.Errorf("listCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringInSlice(t *testing.T) {
	type args struct {
		a    string
		list []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "string in slice",
			args: args{a: "apple", list: []string{"apple", "banana", "orange"}},
			want: true,
		},
		{
			name: "string not in slice",
			args: args{a: "grape", list: []string{"apple", "banana", "orange"}},
			want: false,
		},
		{
			name: "empty slice",
			args: args{a: "apple", list: []string{}},
			want: false,
		},
		{
			name: "empty string in slice",
			args: args{a: "", list: []string{"", "banana"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringInSlice(tt.args.a, tt.args.list); got != tt.want {
				t.Errorf("stringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
