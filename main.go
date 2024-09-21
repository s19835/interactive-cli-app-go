package main

import (
	"github.com/s19835/interactive-cli-app-go/cmd"
	"github.com/s19835/interactive-cli-app-go/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}
