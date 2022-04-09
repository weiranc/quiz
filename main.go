package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	total := 0
	correct := 0

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		total++

		fmt.Printf("Problem #%d: %s = \n", total, rec[0])
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == rec[1] {
			correct++
		}

	}
	fmt.Printf("You scored %d out of %d.\n", correct, total)
}
