package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func parseCSV(filePath string) ([]match, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var matches []match
	for _, record := range records {
		score1, _ := strconv.Atoi(record[2])
		score2, _ := strconv.Atoi(record[3])
		water, _ := strconv.Atoi(record[4])
		matches = append(matches, match{
			team1:  team{record[0], record[1]},
			team2:  team{record[5], record[6]},
			score1: score1,
			score2: score2,
			water:  water,
		})
	}

	return matches, nil
}

func main() {
	filePath := "path/to/your/file.csv"
	matches, err := parseCSV(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	beater := Beater{matchList: matches}
	res := beater.calculate()
	for _, item := range res {
		fmt.Printf("%s 转给 %s %d 瓶水\n", item.from, item.to, item.water)
	}
}
