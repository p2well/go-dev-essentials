package mascot_test

import (
	"testing"

	"github.com/p2well/go-dev-essentials/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "R2-D2" {
		t.Fatal("Wrong mascot")
	}
}