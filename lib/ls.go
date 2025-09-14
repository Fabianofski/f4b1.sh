package lib

import (
	"html/template"
	"strings"
)

func getFilesInDirectory(path string, session *TerminalSession) []string {
	files := []string{}
	for _, v := range session.FileTree {
		relativeFile, ok := strings.CutPrefix(v, path)
		if ok {
			relativeFileTrimmed := strings.TrimSuffix(relativeFile, "/")
			pathParts := strings.Split(relativeFileTrimmed, "/")
			if relativeFile == "" || len(pathParts) > 1 {
				continue
			}
			files = append(files, relativeFile)
		}
	}
	return files
}

func ls(args []string, session *TerminalSession) error {
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
