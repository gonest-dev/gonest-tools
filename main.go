package main

import (
	"fmt"
	"os"

	"github.com/gonest-dev/gonest-tools/badge"
	"github.com/gonest-dev/gonest-tools/clean"
	"github.com/gonest-dev/gonest-tools/mkdir"
	"github.com/gonest-dev/gonest-tools/modules"
	"github.com/gonest-dev/gonest-tools/tag"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gonest-tools <command> [args]")
		os.Exit(1)
	}

	command := os.Args[1]
	os.Args = append([]string{os.Args[0]}, os.Args[2:]...)

	switch command {
	case "badge":
		badge.ExecuteBadge()
	case "clean":
		clean.ExecuteClean()
	case "mkdir":
		mkdir.ExecuteMkdir()
	case "modules":
		modules.ExecuteModules()
	case "tag":
		tag.ExecuteTag()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
