package pets

import (
	"strings"
)

type Dog struct {
	Name string
	Age  int
}

func (p *Dog) Talk() string {
	return "Wuorf"
}

func (d *Dog) BeHappy() string {
	return "yeeeeh!"

}

func (d *Dog) AskToTheDog(question string) string {

	answer := ""

	if strings.Contains(question, d.Name) {
		answer = d.Talk()
	}

	if strings.Contains(question, "food") {
		answer = d.Talk()
	}

	return answer
}
