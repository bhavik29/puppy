package puppy

import (
	"github.com/bhavik29/dog"
)

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof! woof! woof!"
}

func WhenGrownUp() string {
	return dog.WhenGrownUp(Bark())
}

func WhenGrownUps() string {
	return dog.WhenGrownUp(Barks())
}
