package analyze

import (
	"fmt"

	"github.com/cdipaolo/sentiment"
)

var sentimentModel sentiment.Models

func init() {
	model, err := sentiment.Restore()

	if err != nil {
		panic(fmt.Sprintf("Could not restore model!\n\t%v\n", err))
	}

	sentimentModel = model
}
