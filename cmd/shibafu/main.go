package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Code-Hex/shibafu/evaluator"
	"github.com/Code-Hex/shibafu/lexer"
	"github.com/Code-Hex/shibafu/token"
)

var version string

var (
	onelinerFlag    string
	syntaxCheckFlag string
	versionFlag     bool
	helpFlag        bool
)

func main() {
	flag.StringVar(&onelinerFlag, "e", "", "one line of program")
	flag.StringVar(&syntaxCheckFlag, "c", "", "check syntax only")
	flag.BoolVar(&versionFlag, "v", false, "print version number")
	flag.BoolVar(&helpFlag, "h", false, "print help message")
	flag.Parse()
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute: %+v", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 1 || helpFlag {
		flag.Usage()
		return nil
	}

	if versionFlag {
		fmt.Printf("shibafu %s", version)
		return nil
	}

	if onelinerFlag != "" {
		return evaluator.New(onelinerFlag, os.Stdin, os.Stdout).Evaluate(context.Background())
	}

	if syntaxCheckFlag != "" {
		input, err := readFile(syntaxCheckFlag)
		if err != nil {
			return err
		}
		syntaxCheck(input)
		return nil
	}

	input, err := readFile(args[0])
	if err != nil {
		return err
	}

	return evaluator.New(input, os.Stdin, os.Stdout).Evaluate(context.Background())
}

func readFile(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	return string(b), err
}

func syntaxCheck(input string) {
	var fail bool
	lexer := lexer.New(input)
LOOP:
	for {
		t := lexer.NextToken()
		switch t.Type {
		case token.EOF:
			break LOOP
		case token.ILLEGAL:
			fail = true
			fmt.Fprintf(os.Stderr, "line: %d, col: %d, literal: %s\n", t.Line, t.Col, t.Literal)
		}
	}
	if fail {
		os.Exit(1)
	}
	fmt.Println("Syntax OK!")
}
