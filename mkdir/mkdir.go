package mkdir

import (
	"fmt"
	"os"
)

func ExecuteMkdir() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gonest-tools mkdir <dir1> [dir2] ...")
		os.Exit(1)
	}

	for _, dir := range os.Args[1:] {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			os.Exit(1)
		}
		fmt.Printf("Created directory: %s\n", dir)
	}
}
