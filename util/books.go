package util

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func Books() {
	// Open the file
	getFile, err := excelize.OpenFile("./files/go-materials.xlsx")
	if err != nil {
		fmt.Println("Unable to open the excel file:", err)
		return
	}

	defer func() {
		// Close the file
		if err := getFile.Close(); err != nil {
			fmt.Println("Unable to close the excel file:", err)
		}
	}()

	// Get all the rows in the Books sheet
	rows, err := getFile.GetRows("Books")
	if err != nil {
		fmt.Println("Unable to get the Books rows:", err)
		return
	}

	for i, _ := range rows {
		// This statement avoid to print the row 0 that contains cells like "Title, Author, Year, Rate"
		if i == 0 {
			continue
		}

		bookTitle, _ := getFile.GetCellValue("Books", fmt.Sprintf("A%d", i+1))
		bookAuthor, _ := getFile.GetCellValue("Books", fmt.Sprintf("B%d", i+1))
		bookYear, _ := getFile.GetCellValue("Books", fmt.Sprintf("C%d", i+1))
		bookRate, _ := getFile.GetCellValue("Books", fmt.Sprintf("D%d", i+1))

		fmt.Printf("{title: %+v, author: %+v, year: %+v, rate: %+v},\n", bookTitle, bookAuthor, bookYear, bookRate)
	}
}
