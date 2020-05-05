package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const nameList string = "big_name_list.tsv"

var names []string

var rows int
var moneyMin int
var showDate bool
var moneyMax int
var output string
var flagError = false

//var UseCents bool

func init() {
	flag.IntVar(&rows, "rows", -1, "Number of rows in fake data")
	flag.IntVar(&moneyMax, "max", -1, "Maximum amount of money for a donation")
	flag.IntVar(&moneyMin, "min", -1, "Minimum amount of money for a donation")
	flag.BoolVar(&showDate, "date", false, "Show a random date.")
	flag.StringVar(&output, "output", "-", "Output path")
	flag.Parse()

	if rows <= 0 {
		fmt.Println(fmt.Errorf("Set --rows to a positive number"))
		flagError = true
	}
	if !(0 <= moneyMin && moneyMin < moneyMax) {
		fmt.Println(fmt.Errorf("Set 0 <= --min <= --max"))
		flagError = true
	}
}

func randomMoney() int {
	return rand.Intn(moneyMax-moneyMin) + moneyMin
}

func randomName() string {
	return names[rand.Intn(len(names))]
}

func randomDate() time.Time {
	return time.Now().Add(time.Duration(-rand.Intn(1000)) * time.Hour)
}

func openOutput() (io.Writer, bool) {
	if output == "-" {
		return os.Stdout, true
	}
	outf, err := os.Create(output)
	if err != nil {
		log.Println("could not open", output)
		return nil, false
	}
	return outf, true
}

func parseNameList() {

	nl, err := os.Open(nameList)
	if err != nil {
		fmt.Println("could not open name list")
		return
	}
	r := csv.NewReader(nl)
	switch nameList[len(nameList)-4:] {
	case ".tsv":
		r.Comma = '\t'
	}

	ns, err := r.ReadAll()
	if err != nil {
		fmt.Println("could not parse name table")
		return
	}
	for _, row := range ns {
		n := row[0] + " " + row[1]
		names = append(names, n)
	}
}

func main() {
	if flagError {
		return
	}

	parseNameList()
	if len(names) == 0 {
		fmt.Println("need at least one name")
		return
	}

	outf, ok := openOutput()
	if !ok {
		fmt.Println(fmt.Errorf("Could not open output: %s", output))
		return
	}

	w := csv.NewWriter(outf)
	for i := 0; i < rows; i++ {
		r := []string{}
		if showDate {
			r = append(r, randomDate().String()[:10])
		}
		r = append(r, randomName(), strconv.Itoa(randomMoney()))
		w.Write(r)
	}
	w.Flush()

}
