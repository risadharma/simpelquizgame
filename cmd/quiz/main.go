package main

import (
	"flag"

	"github.com/risadharma/simplequiz/internal/quiz"
)

var (
	filepath = flag.String("problemsource", "problems/simplearitmathic.csv", "define where file with problem resource located")
)

func main() {
	flag.Parse()

	csvReader := quiz.NewCSVReader(*filepath)
	moderator := quiz.NewModerator(csvReader)

	moderator.Play()
}
