package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func handleInput() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func selectOption() string {
	var sortOption string
	for {
		fmt.Println("Size sorting options:")
		fmt.Println("1. Descending")
		fmt.Println("2. Ascending")

		sortOption = handleInput()
		if sortOption == "1" || sortOption == "2" {
			return sortOption
		}

		fmt.Println("Wrong option")
	}

}

func main() {

	files := make(map[int64][]string)

	var keys []int64

	if len(os.Args) != 2 {
		fmt.Println("Directory is not specified")
	} else {
		fmt.Println("Enter file format:")
		fileFormat := handleInput()

		//sorter := selectOption()

		err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatal(err)
			}

			if !info.IsDir() {
				if fileFormat == "" {
					files[info.Size()] = append(files[info.Size()], path)
				} else if filepath.Ext(path) == fileFormat {
					files[info.Size()] = append(files[info.Size()], path)
				}

			}

			return nil
		})

		for key, _ := range files {
			keys = append(keys, key)
		}

		fmt.Println(files)

		if err != nil {
			log.Fatal(err)
		}
	}

}
