package lib

import (
	"fmt"
	"github.com/Fabianofski/f4b1.sh/model"
	"html/template"
	"strings"
)

func pathToAbsolutePath(path string, session *model.TerminalSession) string {
	absPath := path

	if strings.HasPrefix(absPath, "..") {
		parts := strings.Split(session.Cwd, "/")
		absPath = strings.Join(parts[:len(parts)-2], "/") + "/"
	}

	if strings.HasPrefix(absPath, ".") {
		cwd := strings.TrimSuffix(session.Cwd, "/")
		absPath = strings.Replace(absPath, ".", cwd, 1)
	}

	if strings.HasPrefix(absPath, "~") {
		absPath = strings.Replace(absPath, "~", session.HomeDir, 1)
	}

	if !strings.HasPrefix(absPath, "/") {
		absPath = session.Cwd + absPath
	}

	if !strings.HasSuffix(absPath, "/") {
		absPath += "/"
	}
	return absPath
}

func getFilesInDirectory(path string, session *model.TerminalSession) []string {
	absPath := pathToAbsolutePath(path, session)

	dir, ok := session.Root[absPath]
	if !ok {
		return []string{fmt.Sprintf("ls: cannot access %s: No such file or directory", path)}
	}

	keys := []string{}
	for k := range session.Root {
		if trimmed, ok := strings.CutPrefix(k, absPath); ok {
			parts := strings.Split(trimmed, "/")
			if len(parts) == 2 {
				keys = append(keys, parts[len(parts)-2])
			}
		}
	}

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
		files = getFilesInDirectory(args[0], session)
	}
	out := template.HTML(strings.Join(files, ", "))
	session.StdOut = append(session.StdOut, out)
	return nil
}
