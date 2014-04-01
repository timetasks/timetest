package pets

import (
	"strings"
)

type Cat struct {
	Name  string
	Age   int
	Lifes int
}

func (p *Cat) Talk() string {
	return "meow"
}

func (c *Cat) Sgrut() string {
	return "sgrunt"

}

func (c *Cat) AskTheCat(question string) string {

	answer := ""
	if strings.Contains(question, "food") {
		answer = c.Talk()
	}

	return answer
}

func (c *Cat) Die() bool {
	c.Lifes--
	return true
}
