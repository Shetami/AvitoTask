package repository

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Report(userId int, description string, price int) {
	var path = "report/data.csv"
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	userId_str := strconv.Itoa(userId)
	price_str := strconv.Itoa(price)
	var data [][]string
	data = append(data, []string{userId_str, description, price_str})
	w := csv.NewWriter(f)
	w.WriteAll(data)
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Appending succed")
}
