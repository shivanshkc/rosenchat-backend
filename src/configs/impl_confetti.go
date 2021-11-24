package configs

import (
	"github.com/shivanshkc/confetti"
)

// loadWithConfetti loads configs using shivanshkc/confetti package.
func loadWithConfetti(target interface{}) {
	if err := confetti.GetLoader().Load(target); err != nil {
		panic(err)
	}
}
