package helloworld

import "testing"

var tests = map[string]struct {
	name     string
	language string
	result   string
}{
	"should say hello to people": {
		name:     "Hippo",
		language: "",
		result:   "Hello, Hippo",
	},
	"should say Hello, Hippo when an empty string is input": {
		name:     "",
		language: "",
		result:   "Hello, Hippo",
	},
	"should say Hello in Spanish": {
		name:     "Hippo",
		language: "Spanish",
		result:   "Hola, Hippo",
	},
	"should say Hello in French": {
		name:     "Hippo",
		language: "French",
		result:   "Bonjour, Hippo",
	},
}

func TestHello(t *testing.T) {
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := Hello(tt.name, tt.language)
			want := tt.result
			if got != want {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}
