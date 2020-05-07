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

	csvReader := quiz.NewCSVReader(*filepath)
	moderator := quiz.NewModerator(csvReader)

	var key string
	fmt.Print("Press enter to start the quiz!")
	fmt.Scanf("%s\n", &key)

	if err := moderator.Start(time.Duration(*duration)); err != nil {
		exitWithError(err)
	}

	os.Exit(1)
}

func exitWithError(err error) {
	fmt.Printf("error occured: %v\n", err)
	os.Exit(1)
}
