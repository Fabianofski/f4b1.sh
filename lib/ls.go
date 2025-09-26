package lib

import (
	"fmt"
	"github.com/Fabianofski/f4b1.sh/model"
	"html/template"
	"strings"
)

func pathToAbsolutePath(path string, session *model.TerminalSession) string {
	absPath := ""

	if strings.HasPrefix(path, ".") {
		absPath = strings.Replace(path, ".", session.Cwd, 1)
		return absPath
	}

	if !strings.HasPrefix(path, "/") {
		absPath = session.Cwd + path
	} else {
		absPath = path
	}

	if strings.HasPrefix(path, "~") {
		absPath = strings.Replace(path, "~", "/home/guest", 1)
	}

	if !strings.HasSuffix(absPath, "/") {
		absPath += "/"
	}
	fmt.Println(absPath)
	return absPath
}

func getFilesInDirectory(path string, session *model.TerminalSession) []string {
	dir, ok := session.Root[path]
	if !ok {
		return []string{fmt.Sprintf("ls: cannot access %s: No such file or directory", path)}
	}

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
		path := pathToAbsolutePath(args[0], session)
		files = getFilesInDirectory(path, session)
	}
	out := template.HTML(strings.Join(files, ", "))
	session.StdOut = append(session.StdOut, out)
	return nil
}
