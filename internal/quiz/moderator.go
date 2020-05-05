package quiz

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Problem represent quiz problem.
// It will imported by reading file that has multiple questions and answers
type Problem struct {
	Question string
	Answer   string
}

// Moderator represent quiz moderator that generate problems and human answer
type Moderator struct {
	fileReader Reader

	playerScore int
}

// NewModerator return new Moderator
func NewModerator(reader Reader) *Moderator {
	return &Moderator{
		fileReader: reader,
	}
}

// Play playing quiz
func (m *Moderator) Play() {
	problems, err := m.fileReader.Read()
	if err != nil {
		panic(err)
	}

	totalProblems := len(problems)
	reader := bufio.NewReader(os.Stdin)

	for i, problem := range problems {
		fmt.Printf("question %d of %d. %s: ", i+1, totalProblems, problem.Question)
		answer, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		if strings.Trim(answer, "\n") == problem.Answer {
			m.playerScore++
		}
	}

	fmt.Printf("Congratulation, you've answer %d questions.\n", m.playerScore)
}
