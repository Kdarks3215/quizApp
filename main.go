package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := "ghCapitals.csv"
	timeLimit := 90

	file, err := os.Open(csvFilename)
	if err != nil {
		fmt.Printf("Failed to open the CSV file :%s\n ", csvFilename)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	
	
	correct := 0
	for i, p := range problems{
		fmt.Printf("Problem #%d %s:  \n", i+1, p.q)
		answerCh := make(chan string)
		go func(){
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			answer := strings.Title(scanner.Text())
			answerCh <- answer 
		}()
		select {
		case <-timer.C:
			fmt.Println("\nYou ran out of time!")
			fmt.Printf("You scored %d out of %d. \n", correct, len(problems))
			return
		case answer := <- answerCh:
			if answer == p.a {
				correct++
	}
	}
		}
		
	fmt.Printf("You scored %d out of %d. \n", correct, len(problems))
	time.Sleep(5 * time.Second)
}
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines{
		ret[i] =  problem{
		q: line[0],
		a: line[1],
	}
}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}