package pets

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func setupCat() Cat {

	return Cat{"Piero", 14, 9}

}

func TestThatCatsMeows(t *testing.T) {

	newCat := setupCat()

	Convey("Make it meow", t, func() {
		So(newCat.Talk(), ShouldEqual, "meow")
	})

}

func TestThatIsARealCat(t *testing.T) {

	newCat := setupCat()

	Convey("Make it sgrut", t, func() {
		So(newCat.Sgrut(), ShouldEqual, "sgrunt")
	})

}

func TestThatItUsuallyIgnoresYou(t *testing.T) {

	newCat := setupCat()

	Convey("Try it ignores you..", t, func() {
		So(newCat.AskTheCat("Hey Cat!"), ShouldBeEmpty)
	})

	Convey("... Unless it's about food", t, func() {
		So(newCat.AskTheCat("Do you want some food?"), ShouldNotBeEmpty)
	})

}

func TestThaItCanLostLifesWithoutDying(t *testing.T) {

	newCat := setupCat()

	Convey("Check it has enough lifes", t, func() {
		So(newCat.Lifes, ShouldBeGreaterThan, 1)
	})
	lifes := newCat.Lifes
	Convey("Kill the cat", t, func() {
		So(newCat.Die(), ShouldBeTrue)
	})

	Convey("Check lost one life", t, func() {
		So(newCat.Lifes, ShouldEqual, lifes-1)
	})

}

func TestThatICanDie(t *testing.T) {

	newCat := setupCat()

	Convey("Check it has enough lifes", t, func() {
		So(newCat.Lifes, ShouldBeGreaterThan, 1)
	})

	lifes := newCat.Lifes
	for i := 0; i < lifes; i++ {
		Convey("Kill the cat", t, func() {
			So(newCat.Die(), ShouldBeTrue)
		})
	}

	Convey("Check that it's dead", t, func() {
		So(newCat.Lifes, ShouldEqual, 0)
	})

}
