package main

import (
	"github.com/Fabianofski/f4b1.sh/model"
)

var DefaultFileTree = map[string]model.Directory{
	"/": {
		Files: map[string]model.File{
			"README.txt": {
				Content: "Welcome to this terminal! Feel free to look around :)",
			},
		},
	},
	"/home/": {
		Files: map[string]model.File{},
	},
	"/home/guest/": {
		Files: map[string]model.File{
			"about": {
				Content: "Hello World!",
			},
			".hidden": {
				Content: "Congrats! You found a hidden file!",
			},
			"todo.txt": {
				Content: "- learn vi\n- build something cool\n- find all the Easter Eggs",
			},
			"joke.txt": {
				Content: "Why do programmers prefer dark mode?\nBecause light attracts bugs.",
			},
		},
	},
	"/usr/share/": {
		Files: map[string]model.File{
			"motd": {
				Content: "Message of the day: Stay curious!",
			},
		},
	},
	"/etc/": {
		Files: map[string]model.File{
			"secrets.conf": {
				Content: "42",
			},
		},
	},
	"/games/": {
		Files: map[string]model.File{
			"snake": {
				Content: "Sorry, this game is not implemented yet ;)",
			},
			"ascii_art.txt": {
				Content: "¯\\_(ツ)_/¯",
			},
		},
	},
	"/var/log/": {
		Files: map[string]model.File{
			"system.log": {
				Content: "Nothing to see here... or is there?",
			},
			"hack.log": {
				Content: "ALERT: Guest attempted sudo at 03:14AM",
			},
		},
	},
}
