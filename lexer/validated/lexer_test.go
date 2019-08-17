package validated

import (
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "let x = fn() {}"
	_, tokens := Lex("yeet", input)

	for t := range tokens {
		fmt.Printf("%s: %s\n", t.Type, t.Literal)
	}
}
