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

func getFilesInDirectory(path string, session *model.TerminalSession, showHidden bool) []string {
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
		if !strings.HasPrefix(k, ".") || showHidden {
			keys = append(keys, k)
		}
	}
	return keys
}

func ls(args []string, session *model.TerminalSession) error {
	files := []string{}
	showHidden := false

	dirs := []string{}
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			if arg == "-a" {
				showHidden = true
			}
		} else {
			dirs = append(dirs, arg)
		}
	}

	if len(dirs) == 0 {
		files = getFilesInDirectory(session.Cwd, session, showHidden)
	} else {
		for _, v := range dirs {
			files = getFilesInDirectory(v, session, showHidden)
		}
	}

	out := strings.Join(files, " ")
	session.StdOut = append(session.StdOut, template.HTML(out))
	return nil
}
