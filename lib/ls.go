package lib

import (
	"github.com/Fabianofski/f4b1.sh/model"
	"html/template"
	"strings"
)

func getFilesInDirectory(path string, session *model.TerminalSession) []string {
	dir := session.Root[path]

	keys := make([]string, 0, len(dir.Files))
	for k := range dir.Files {
		keys = append(keys, k)
	}
	return keys
}

func ls(args []string, session *model.TerminalSession) error {
	files := []string{}
	if len(args) == 0 {
		files = getFilesInDirectory(session.Cwd, session)
	} else {
		path := ""
		if !strings.HasPrefix(args[0], "/") {
			path = session.Cwd + args[0]
		} else {
			path = args[0]
		}

		if !strings.HasSuffix(path, "/") {
			path += "/"
		}

		files = getFilesInDirectory(path, session)
	}
	out := template.HTML(strings.Join(files, ", "))
	session.StdOut = append(session.StdOut, out)
	return nil
}
