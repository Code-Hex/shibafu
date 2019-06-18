package evaluator

import (
	"bytes"
	"testing"
)

const helloworld = `
wwwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wwWwwWWWwWwwWwwWwwWwwWwwWwwWwwWwwwwWwWWWwWW
WWWWWwwwwwwWwwWwwWwwWwwWwwWwwWw
wwWwwWWWwWwwWwwWwwWwwwwWwWWWwWW
WWWwWwWWwwwWwwWwwWwwWwwWwwWwwWw
WWwwWWwwwWwwWwwWwWWwwwwWwwWwWWWwWWwww
wWwwWwwWwwWwwWwwWwwWwwWw
wwWwwWWWwWwwWwwWwwWwwwwWwWWWwWW
WWWWWwwwwwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wwWwwWWWwWwwWwwWwwWwwWwwwwWwWWWwWW
WWWWWwwwwwwWwwWwwWwwWwwWwwWwwWwwWw
wwWwwWWWwWwwWwwWwwwwWwWWWwWW
WWWWWwwwWwwWwwWwWWww
WwWWwWWwWWwWWwWWwWWWww
WwWWwWWwWWwWWwWWwWWwWWwWWWww
wwWwwWwWWWwWWwwwwWwwWwwWwwWwwWwwWwwWwwWw
wwWwwWWWwWwwWwwWwwWwwwwWwWWWwWW
WWWwWwWWwwwwWwwWwWWWwWW
wWwwWwwWwwWwwWwwWwwWwwWwwWwwWwWWww
`

const add = `
wWwwwwwWwwWwwwwWWWWWW
wwwwwWwwWwWWWWwWwwwwWWwWWWWW
wWwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wWwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wWwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wWwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wWwwWwwWwwWwwWwwWwwWwwWwWWww
`

func TestEvaluator_Evaluate(t *testing.T) {
	buf := &bytes.Buffer{}
	stdout = buf // switching output dest

	tests := []struct {
		name    string
		input   string
		wantOut string
		wantErr bool
	}{
		{
			name:    "Hello World!",
			input:   helloworld,
			wantOut: "Hello World!\n",
			wantErr: false,
		},
		{
			name:    "1 + 2 = 3",
			input:   add,
			wantOut: "3",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer buf.Reset()
			e := New(tt.input)
			if err := e.Evaluate(); (err != nil) != tt.wantErr {
				t.Errorf("Evaluator.Evaluate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if buf.String() != tt.wantOut {
					t.Errorf("got `%s` but want out `%s`", buf.String(), tt.wantOut)
				}
			}
		})
	}
}
