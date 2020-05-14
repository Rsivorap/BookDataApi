package csvtoDB

import (
	//"database/sql"
	"encoding/csv"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"os"
	"strconv"
)

//CSV format
// BookId, Title, Authors, Average Rating, ISBN, ISBN13, Language Cde, Num Pages, Ratings, Reviews
type BookData struct {
	BookID  string
	Title   string
	Authors string
	Rating  int
	ISBN    int
	ISBN13  int
	Lang    string
	Pages   int
	Ratings int
	Reviews int
}

func csvToBookData(record []string) *BookData {
	temp := new(BookData)
	temp.BookID = record[0]
	temp.Title = record[1]
	temp.Authors = record[2]
	t, _ := strconv.Atoi(record[3])
	temp.Rating = t
	t, _ = strconv.Atoi(record[4])
	temp.ISBN = t
	t, _ = strconv.Atoi(record[5])
	temp.ISBN13 = t
	temp.Lang = record[6]
	t, _ = strconv.Atoi(record[7])
	temp.Pages = t
	t, _ = strconv.Atoi(record[8])
	temp.Ratings = t
	t, _ = strconv.Atoi(record[9])
	temp.Reviews = t

	return temp
}

func GenerateDBData(path string) []BookData {
	csv_file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open csv", err)
		return nil
	}
	r := csv.NewReader(csv_file)
	var data []BookData
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		temp := csvToBookData(record)
		data = append(data, *temp)

	}

	return data
}
