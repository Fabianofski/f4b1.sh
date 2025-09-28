package model

import "html/template"

type TerminalSession struct {
	StdOut       []template.HTML
	InputAllowed bool
	Cwd          string
	CwdShort     string
	HomeDir      string
	Root         map[string]Directory
	Id           string
}

type Message struct {
	Input string
}
