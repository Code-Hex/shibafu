package evaluator

import (
	"bytes"
	"testing"
)

const helloworld = `
wwwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wWWWWWwWwwWwwWwwWwwWwwWwwWwwWwwwwWwWWWw
WWWWwwwwwwWwwWwwWwwWwwWwwWwwWw
wWWWWWwWwwWwwWwwWwwwwWwWWWw
WWWwWwWwwwWwwWwwWwwWwwWwwWwwWwWwwWwwwWwwWwwWwWww
wWWWwWWWwwwwwWwwWwwWwwWwwWwwWwwWwwWw
wWWWWWwWwwWwwWwwWwwwwWwWWWwWWWWwwwwwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wWWWWWwWwwWwwWwwWwwWwwwwWwWWWwWWWWwwwwwwWwwWwwWwwWwwWwwWwwWwwWw
wWWWWWwWwwWwwWwwwwWwWWWwWWWWwwwWwwWwwWwWwwWwWWwWWwWWwWWwWWwW
WwwWwWWwWWwWWwWWwWWwWWwWWwWWwwwWWWwWWWwwww
wWwwWwwWwwWwwWwwWwwWwwWw
wWWWWWwWwwWwwWwwWwwwwWwWWWw
WWWwWwWww
wWWWwWWWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWwWww
`

const add = `
wWwwwwwWwwWwwwwWWWWWW
wwwwWWWwWWWWwWwwwwWWwWWW <- Point
wWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWwWww
`

const mul = `
wWwwWwwWwwWwwwwwWwwWwwwwWWWWWW
wWWWwW 	                          <- Point 1
wwwwWWWwWwwwwwwwWwWWWWWWWWw       <- Point 2
wwwwwwwWWWwWWWWwWwWWWwWwwwwwwwWWw <- Point 3
WWWWWWWWW                         <- Point 4
WWwwwwwww                         <- Point 5
wWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWw
wWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWwwWw
Www
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
		{
			name:    "4 * 2 = 8",
			input:   mul,
			wantOut: "8",
			wantErr: false,
		},
		{
			name:    "Invalid",
			input:   `w`,
			wantErr: true,
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
