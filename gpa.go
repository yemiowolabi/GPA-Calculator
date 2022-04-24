package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/////FOR A 5 POINT SCALE
var grades = map[string]float64{
	"A": 5,
	"B": 4,
	"C": 3,
	"D": 2,
	"E": 1,
	"F": 0,
}
var reader = bufio.NewReader(os.Stdin)

func main() {
	var totalunits float64
	var totalpoints float64
	var noOfCourses int

	fmt.Println("GPA Calculator v1.0.0")
	fmt.Println(strings.Repeat("=", 40))
	// n, err := fmt.Scan(&noOfCourses)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for {
		fmt.Print("Enter number of Courses offered: ")
		coursenum, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		coursenum = strings.Trim(coursenum, "\n\r")
		var err2 error
		noOfCourses, err2 = strconv.Atoi(coursenum)
		if err2 != nil {
			log.Fatal(err2)
		}
		if err2 == nil && err == nil {
			if noOfCourses > 0 && noOfCourses <= 20 {
				break
			} else {
				err := errors.New("you must enter at least a Course and not more than 20")
				log.Fatal(err)
			}
		}
	}
	fmt.Println(strings.Repeat("-", 40))
	for i := 0; i < noOfCourses; i++ {
		var grade float64

		for {
			var score int
			var letterGrade string
			fmt.Print("Enter Score: ")

			// _, err := fmt.Scan(&score)
			// if err != nil {
			// 	log.Fatal(err)
			// }

			scoreInput, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			scoreInput = strings.Trim(scoreInput, "\n\r")
			score, err2 := strconv.Atoi(scoreInput)
			if err2 != nil {
				log.Fatal(err2)
			} else {
				if score <= 100 && score >= 70 {
					letterGrade = "A"
				} else if score < 70 && score >= 60 {
					letterGrade = "B"
				} else if score < 60 && score >= 50 {
					letterGrade = "C"
				} else if score < 50 && score >= 45 {
					letterGrade = "D"
				} else if score < 45 && score >= 40 {
					letterGrade = "E"
				} else if score < 40 && score >= 0 {
					letterGrade = "F"
				} else if score < 0 || score > 100 {
					err := errors.New("score cannot be less than 0 nor more than 100")
					log.Fatal(err)
				}

				if err == nil && err2 == nil {
					var ok bool
					grade, ok = grades[letterGrade]
					if ok {
						break
					}
				}
			}
		}

		var units float64
		for {
			fmt.Print("Enter Number of Units: ")
			unitsInput, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			unitsInput = strings.Trim(unitsInput, "\n\r")
			var err2 error
			units, err2 = strconv.ParseFloat(unitsInput, 64)
			if err2 != nil {
				log.Fatal(err2)
			}

			if err2 == nil && err == nil {
				if units > 0. && units <= 5. {
					break
				} else {
					err := errors.New("course units cannot be less than 1 nor exceed 5")
					log.Fatal(err)
				}
			}
		}
		fmt.Println(strings.Repeat("=", 40))

		totalunits += units
		totalpoints += units * grade
	}

	var gpa float64
	if totalunits != 0.0 {
		gpa = totalpoints / totalunits
	}
	fmt.Printf("GPA: %.3f", gpa)
}
