package quiz

import (
	"fmt"
	"strings"
	"time"
)

// Problem represent quiz problem.
// It will imported by reading file that has multiple questions and answers
type Problem struct {
	Question string
	Answer   string
}

// Moderator represent quiz moderator that generate problems and human answer
type Moderator struct {
	problems   []*Problem
	score      int
	fileReader Reader
}

// NewModerator return new Moderator
func NewModerator(reader Reader) *Moderator {
	return &Moderator{
		fileReader: reader,
	}
}

// Start start quiz game
func (m *Moderator) Start(duration time.Duration) error {
	if err := m.prepareProblems(); err != nil {
		return err
	}

	func() {
		var answer string
		ticker := time.NewTicker(duration * time.Second)
		defer ticker.Stop()

		for i, problem := range m.problems {
			fmt.Printf("Question #%d: %s= ", i+1, problem.Question)

			answerChannel := make(chan string)
			go func() {
				fmt.Scanf("%s\n", &answer)
				answerChannel <- answer
			}()

			select {
			case <-ticker.C:
				fmt.Printf("\nCongratulation, you've answer %d of %d questions.\n", m.score, len(m.problems))
				return
			case answer := <-answerChannel:
				if strings.TrimSpace(answer) == problem.Answer {
					m.score++
				}
			}
		}
	}()

	return nil
}

func (m *Moderator) prepareProblems() error {
	problems, err := m.fileReader.Read()
	if err != nil {
		return err
	}
	m.problems = problems

	return nil
}
