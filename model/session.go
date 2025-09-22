package model

import "html/template"

type TerminalSession struct {
	StdOut       []template.HTML
	InputAllowed bool
	Cwd          string
	Root         map[string]Directory
}

type Message struct {
	Input string
}
