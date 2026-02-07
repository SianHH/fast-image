package main

import (
	"fast-image/cmd"
	_ "fast-image/router"
	_ "fast-image/task"
	_ "fast-image/todo"
	"log"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
