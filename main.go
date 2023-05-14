package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func handleInput() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func selectOption() string {
	var sortOption string

	fmt.Println("\nSize sorting options:")
	fmt.Println("1. Descending")
	fmt.Println("2. Ascending")

	fmt.Println("\nEnter a sorting option:")
	sortOption = handleInput()

	switch sortOption {
	case "1":
		return sortOption
	case "2":
		return sortOption
	default:
		fmt.Println("\nWrong option")
		selectOption()

	}
	return ""
}

func hashFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	md5Hash := md5.New()
	// Copy the data from 'hello.txt' to the 'md5Hash' interface until reaching EOF:
	if _, err := io.Copy(md5Hash, file); err != nil {
		fmt.Println(err)
	}
	return md5Hash.Sum(nil)
}

func validateFiles(filesToValidate []string, files []string) bool {

	isValid := false
	for _, val := range filesToValidate {
		i, err := strconv.Atoi(val)
		if err != nil {
			return false
		}
		for k := range files {
			if i == k+1 {
				isValid = true
				break
			}
		}
	}
	return isValid
}

func selectFiles(files []string) []string {
	fmt.Println("\nEnter file numbers to delete:")
	fileNumbers := handleInput()
	splitFileNumbers := strings.Split(fileNumbers, " ")

	if len(splitFileNumbers) == 0 {
		fmt.Println("\nWrong format")
		return selectFiles(files)

	} else if len(splitFileNumbers) > len(files) {
		fmt.Println("\nWrong format")
		return selectFiles(files)

	} else {
		isValid := validateFiles(splitFileNumbers, files)

		if !isValid {
			fmt.Println("\nWrong format")
			return selectFiles(files)
		}
	}
	return splitFileNumbers
}

func deleteFiles(files []string) int64 {

	fmt.Println("\nDelete files?")
	input := handleInput()

	switch input {
	case "yes":
		var freedUpSize int64
		selectedFiles := selectFiles(files)

		for _, validFile := range selectedFiles {
			getIndex, _ := strconv.Atoi(validFile)

			fileInfo, err := os.Stat(files[getIndex-1])
			if err != nil {
				fmt.Println(err)
			}

			freedUpSize += fileInfo.Size()

			err = os.Remove(files[getIndex-1])
			if err != nil {
				fmt.Println(err)
			}
		}

		return freedUpSize
	case "no":
		return 0
	default:
		fmt.Println("\nWrong option")
		deleteFiles(files)
	}
	return 0
}

func checkDuplicates(files map[int64][]string, sortedKeys []int64) []string {
	fmt.Println("Check for duplicates?")
	input := handleInput()

	switch input {
	case "yes":
		counter := 1
		var controller []string

		for _, value := range sortedKeys {
			categorized := false
			duplicated := make(map[string][]string)

			if len(files[value]) > 1 {

				for _, filePath := range files[value] {
					hashPath := fmt.Sprintf("%x\n", hashFile(filePath))
					duplicated[hashPath] = append(duplicated[hashPath], filePath)
				}
				for key, hashValue := range duplicated {

					if len(hashValue) > 1 {
						if !categorized {
							fmt.Println("")
							fmt.Println(value, " bytes")
							categorized = true
						}
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

		return controller

	case "no":
		break
	default:
		fmt.Println("Wrong option")
		checkDuplicates(files, sortedKeys)
	}

	return []string{}

}

func main() {

	files := make(map[int64][]string)

	var keys []int64

	var freedSpace int64

	if len(os.Args) != 2 {
		fmt.Println("Directory is not specified")

	} else {
		fmt.Println("\nEnter file format:")
		fileFormat := handleInput()

		sorter := selectOption()
		fmt.Print("\n")

		err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Println(err)
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
			fmt.Println(err)
		}

		for key := range files {
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
			fmt.Print("\n")
		}

		controller := checkDuplicates(files, keys)

		if len(controller) != 0 {
			freedSpace = deleteFiles(controller)
			if freedSpace != 0 {
				fmt.Printf("\nTotal freed up space: %d bytes", freedSpace)
			}

		}

	}

}
