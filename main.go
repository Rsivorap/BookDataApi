package main

import (
	"./csvtoDB"
	"./database"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func main() {

	// Get Credentials from YAML file
	dbLogin, err := ioutil.ReadFile("credentials.yaml")
	if err != nil {
		log.Fatalln("Could not open dbLogin credentials", err)
	}

	m := make(map[string]string)
	yaml.Unmarshal(dbLogin, &m)
	fmt.Println(m)

	// Generate array of structs from csv
	data := csvtoDB.GenerateDBData("resources/books.csv")
	fmt.Println("Testing Generate csv")
	fmt.Println("Length of Data: ", len(data))

	// Grab Credentials from Yaml

	db := database.GenerateDb(m["root"], m["pass"], m["address"], m["dbName"])

	for _, entry := range data {

		_ = db.Insert("BookData", entry)
	}
	//fmt.Println(str)

}
