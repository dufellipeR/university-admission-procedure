package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
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

func hashFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	md5Hash := md5.New()
	// Copy the data from 'hello.txt' to the 'md5Hash' interface until reaching EOF:
	if _, err := io.Copy(md5Hash, file); err != nil {
		log.Fatal(err)
	}
	return md5Hash.Sum(nil)
}

func checkDuplicates(files map[int64][]string, sortedKeys []int64) {
	fmt.Println("Check for duplicates?")
	input := handleInput()

	switch input {
	case "yes":
		counter := 1
		var controller []string

		for _, value := range sortedKeys {
			duplicated := make(map[string][]string)

			if len(files[value]) > 1 {
				fmt.Println("")
				fmt.Println(value, " bytes")

				for _, filePath := range files[value] {
					hashPath := fmt.Sprintf("%x\n", hashFile(filePath))
					duplicated[hashPath] = append(duplicated[hashPath], filePath)
				}
				for key, hashValue := range duplicated {
					if len(hashValue) > 1 {
						fmt.Printf("Hash: %s", key)
						for index := range hashValue {
							fmt.Printf("%d. %s \n", counter, hashValue[index])
							controller = append(controller, hashValue[index])
							counter++
						}

					}
				}

			}
		}
	case "no":
		break
	default:
		fmt.Println("Wrong option")
		checkDuplicates(files, sortedKeys)
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

		sorter := selectOption()

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
		if err != nil {
			log.Fatal(err)
		}

		for key, _ := range files {
			keys = append(keys, key)
		}

		sort.Slice(keys, func(i, j int) bool {
			if sorter == "1" {
				return keys[i] > keys[j]
			} else {
				return keys[i] < keys[j]
			}
		})

		for index := range keys {
			fmt.Println(keys[index], " bytes")
			for _, value := range files[keys[index]] {
				fmt.Println(value)
			}
			fmt.Println("")
		}

		checkDuplicates(files, keys)

	}

}
