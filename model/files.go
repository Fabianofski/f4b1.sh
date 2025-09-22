package model

type File struct {
	Content string
}

type Directory struct {
	Path  string
	Files map[string]File
}
