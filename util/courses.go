package util

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Course struct {
	ID     int     `json:"id"`
	Course string  `json:"course"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	URL    string  `json:"URL"`
}

func Courses() {
	// Open the file
	getFile, err := excelize.OpenFile("./files/go-materials.xlsx")
	if err != nil {
		fmt.Println("Unable to open the file:", err)
		return
	}

	defer func() {
		// Close the file
		if err := getFile.Close(); err != nil {
			fmt.Println("Unable to close the file:", err)
		}
	}()

	// Get all the rows in the Courses sheet
	rows, err := getFile.GetRows("Courses")
	if err != nil {
		fmt.Println("Unable to get the rows from Courses:", err)
		return
	}

	var courses []Course
	for i, _ := range rows {
		// This statement avoid to get the row 0 with cells like "Author, Price, URL"
		if i == 0 {
			continue
		}

		courseTitle, _ := getFile.GetCellValue("Courses", fmt.Sprintf("A%d", i+1))
		courseAuthor, _ := getFile.GetCellValue("Courses", fmt.Sprintf("B%d", i+1))
		coursePriceString, _ := getFile.GetCellValue("Courses", fmt.Sprintf("C%d", i+1))
		courseURL, _ := getFile.GetCellValue("Courses", fmt.Sprintf("D%d", i+1))

		// Convert coursePriceString from string to float64
		coursePrice, err := strconv.ParseFloat(coursePriceString, 64)
		if err != nil {
			fmt.Println("Unable to concert string to float64:", err)
			return
		}

		course := Course{
			ID:     i,
			Course: courseTitle,
			Author: courseAuthor,
			Price:  coursePrice,
			URL:    courseURL,
		}

		courses = append(courses, course)
	}

	fmt.Println(courses)

}
