package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Directory is not specified")
	} else {
		err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatal(err)
			}

			if !info.IsDir() {
				fmt.Println(path)
			}

			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	}

}
