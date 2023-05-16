package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func handleInput() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

type Applicant struct {
	Name, Surname string
	Gpa           float64
}

func main() {

	totalApplicants, _ := strconv.Atoi(handleInput())

	vacancies, _ := strconv.Atoi(handleInput())

	var applicants []Applicant

	for i := 0; i < totalApplicants; i++ {
		person := handleInput()
		splittedPerson := strings.Split(person, " ")
		gpa, err := strconv.ParseFloat(splittedPerson[2], 64)
		if err != nil {
			log.Fatal("GPA is not a number")
		}
		applicants = append(applicants, Applicant{Name: splittedPerson[0], Surname: splittedPerson[1], Gpa: gpa})
	}

}
