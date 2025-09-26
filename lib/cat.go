package lib

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/Fabianofski/f4b1.sh/model"
)

func readContentOfFile(path string, session *model.TerminalSession) string {
	pathParts := strings.Split(path, "/")
	fileName := pathParts[len(pathParts)-1]

	absDir := pathToAbsolutePath(strings.Join(pathParts[:len(pathParts)-1], "/"), session)

	dir, ok := session.Root[absDir]
	if !ok {
		return fmt.Sprintf("cat: cannot access %s: No such file or directory", path)
	}

	file, ok := dir.Files[fileName]
	if !ok {
		return fmt.Sprintf("cat: cannot access %s: No such file or directory", path)
	}

	return file.Content
}

func cat(args []string, session *model.TerminalSession) error {
	if len(args) == 0 {
		errorMsg := "cat: No argument supplied"
		session.StdOut = append(session.StdOut, template.HTML(errorMsg))
	} else {
		for _, path := range args {
			out := template.HTML(readContentOfFile(path, session))
			session.StdOut = append(session.StdOut, out)
		}
	}
	return nil
}
