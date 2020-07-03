package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// ExampleAdd serves as an example for godoc
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
}
