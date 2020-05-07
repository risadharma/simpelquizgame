package quiz

import (
	"fmt"
	"strings"
	"time"
)

// Problem represent quiz problem.
// It will imported by reading file that has multiple questions and answers
type Problem struct {
	question string
	answer   string
}

// Quiz is represents quiz game
type Quiz struct {
	problems      []*Problem
	correctAnswer int
}

// NewQuiz return new Quiz
func NewQuiz(problems []*Problem) *Quiz {
	return &Quiz{
		problems:      problems,
		correctAnswer: 0,
	}
}

// Run represents to run quiz game
func (q *Quiz) Run(duration time.Duration) {
	var answer string
	ticker := time.NewTicker(duration * time.Second)
	defer ticker.Stop()

	for i, problem := range q.problems {
		fmt.Printf("Question #%d: %s= ", i+1, problem.question)

		answerChannel := make(chan string)
		go func() {
			fmt.Scanf("%s\n", &answer)
			answerChannel <- answer
		}()

		select {
		case <-ticker.C:
			fmt.Printf("\nCongratulation, you've answer %d of %d questions.\n", q.correctAnswer, len(q.problems))
			return
		case answer := <-answerChannel:
			if strings.TrimSpace(answer) == problem.answer {
				q.correctAnswer++
			}
		}
	}

}
