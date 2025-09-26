package model

type File struct {
	Content string
}

type Directory struct {
	Files map[string]File
}
