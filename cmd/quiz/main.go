package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/risadharma/simplequiz/internal/quiz"
)

var (
	filepath = flag.String("problemsource", "problems/simplearitmathic.csv", "define where file with problem resource located")
	duration = flag.Int("duration", 30, "define game duration in second")
)

func main() {
	flag.Parse()

	csvReader := quiz.NewCSVReader()

	file, err := os.Open(*filepath)
	if err != nil {
		exit(err)
	}

	quiz := quiz.NewQuiz(csvReader.Read(file))

	var key string
	fmt.Print("Press enter to start the quiz!")
	fmt.Scanf("%s\n", &key)

	quiz.Run(time.Duration(*duration))
}

func exit(err error) {
	fmt.Printf("error occured: %v\n", err)
	os.Exit(1)
}
