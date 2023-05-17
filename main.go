package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func handleInput() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

type Applicant struct {
	Name, Surname, First, Second, Third string
	Gpa                                 float64
}

func sortApplicants(applicants *[]Applicant) {
	sort.Slice(*applicants, func(i, j int) bool {
		if (*applicants)[i].Gpa != (*applicants)[j].Gpa {
			return (*applicants)[i].Gpa > (*applicants)[j].Gpa
		}
		if (*applicants)[i].Name != (*applicants)[j].Name {
			return (*applicants)[i].Name < (*applicants)[j].Name
		}

		return (*applicants)[i].Surname < (*applicants)[j].Surname
	})
}

func readApplicants(applicants *[]Applicant) {
	// Read applicants on txt and assign to applicants

	file, err := os.Open("applicants.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // this line closes the file before exiting the program

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		applicant := new(Applicant)
		applicant.Name = words[0]
		applicant.Surname = words[1]
		applicant.Gpa, _ = strconv.ParseFloat(words[2], 64)
		applicant.First = words[3]
		applicant.Second = words[4]
		applicant.Third = words[5]
		*applicants = append(*applicants, *applicant)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	maxVacancies, _ := strconv.Atoi(handleInput())
	var applicants []Applicant

	subjects := map[string][]Applicant{}

	readApplicants(&applicants)
	sortApplicants(&applicants)

	for _, applicant := range applicants {
		if len(subjects[applicant.First]) < maxVacancies {
			subjects[applicant.First] = append(subjects[applicant.First], applicant)
			continue
		}
		if len(subjects[applicant.Second]) < maxVacancies {
			subjects[applicant.Second] = append(subjects[applicant.Second], applicant)
			continue
		}
		if len(subjects[applicant.Third]) < maxVacancies {
			subjects[applicant.Third] = append(subjects[applicant.Third], applicant)
			continue
		}

	}

	for key, applicants := range subjects {
		fmt.Println(key)
		for _, applicant := range applicants {
			fmt.Println(applicant.Name, applicant.Gpa)
		}
		fmt.Println("\n")
	}
}
