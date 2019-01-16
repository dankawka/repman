package main

import (
	commands "github.com/dankawka/repman/internal/app/repman"
)

func main() {
	commands.RegisterCommands()
	commands.Execute()
}
