package evaluator

import (
	"bytes"
	"context"
	"os"
	"testing"
	"time"
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
		{
			name:    "Invalid stack overflow",
			input:   infLoop,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			e := New(tt.input, os.Stdin, buf)
			if err := e.Evaluate(context.Background()); (err != nil) != tt.wantErr {
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

const infLoop = `wWwwWWwWWwwwwWwWWwwWwWWw`

func TestTimeout(t *testing.T) {
	buf := &bytes.Buffer{}
	e := New(infLoop, os.Stdin, buf)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := e.Evaluate(ctx)
	if err == nil {
		t.Fatalf("Evaluator.Evaluate() error = %v, wantErr %v", nil, true)
	}
	if err != context.DeadlineExceeded {
		t.Fatalf("Evaluator.Evaluate() error = %v, want %v", err, context.DeadlineExceeded)
	}
}
