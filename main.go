package main

import (
	"dojogpt/container"
	"os"
)

func main() {
	prompt := os.Args[1]
	awnser := container.Ask().MakeQuestion(prompt)
	println(awnser)
}
