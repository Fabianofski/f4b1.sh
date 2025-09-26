package main

import (
	"github.com/Fabianofski/f4b1.sh/model"
)

var DefaultFileTree = map[string]model.Directory{
	"/": {
		Files: map[string]model.File{},
	},
	"/home/": {
		Files: map[string]model.File{},
	},
	"/home/guest/": {
		Files: map[string]model.File{
			"about": {
				Content: "Hello World!",
			},
		},
	},
}
