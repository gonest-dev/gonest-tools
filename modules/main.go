package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type WorkJSON struct {
	Use []struct {
		DiskPath string
	}
}

func main() {
	cmd := exec.Command("go", "work", "edit", "-json")
	out, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var work WorkJSON
	if err := json.Unmarshal(out, &work); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling: %v\n", err)
		os.Exit(1)
	}

	arg := ""
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	var results []string
	for _, u := range work.Use {
		path := u.DiskPath
		if !strings.HasPrefix(path, "./") && !strings.HasPrefix(path, "../") {
			path = "./" + path
		}

		switch arg {
		case "--packages":
			results = append(results, path+"/...")
		case "--dirs":
			results = append(results, path)
		default:
			results = append(results, path)
		}
	}

	fmt.Print(strings.Join(results, " "))
}
