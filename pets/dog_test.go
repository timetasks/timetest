package pets

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestThatADogTalks(t *testing.T) {

	newDog := Dog{"Piero", 14}
	newDog.Talk()

	Convey("Make it talks", t, func() {
		So(newDog.Talk(), ShouldNotBeEmpty)
	})

}

func TestNewDog(t *testing.T) {

	newDog := Dog{"Piero", 14}
	newDog.Talk()

	Convey("Make Sure it's old enought", t, func() {
		So(newDog.Age, ShouldEqual, 14)
	})

	Convey("Make Sure it has the right name", t, func() {
		So(newDog.Name, ShouldEqual, "Piero")
	})

}

func TestThatItTalksIfCalledByName(t *testing.T) {

	newDog := Dog{"Piero", 14}

	Convey("If called by name it Should Talk", t, func() {
		So(newDog.AskToTheDog("Are you Piero?"), ShouldNotBeEmpty)

	})
}

func TestThatItTalksIfAskedAboutFood(t *testing.T) {

	newDog := Dog{"Piero", 14}

	Convey("If called by name it Should Talk", t, func() {
		So(newDog.AskToTheDog("Do you want some food?"), ShouldNotBeEmpty)

	})
}

func TestThatItsReallyHappy(t *testing.T) {

	newDog := Dog{"Piero", 14}

	Convey("Make Sure that happines is for real", t, func() {
		So(newDog.BeHappy(), ShouldEqual, "yeeeeh!")
	})

}
