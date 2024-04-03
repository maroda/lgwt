package integers

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(4, 20)
	fmt.Println(sum)
	// Output: 24
}
