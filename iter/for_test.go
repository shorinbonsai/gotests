package iter

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	looped := Loopy("a", 6)
	expected := "aaaaaa"
	if looped != expected {
		t.Errorf("expected %q but got %q", expected, looped)
	}
}

func BenchmarkLoopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loopy("a", 5)
	}
}

func ExampleLoopy() {
	thingy := Loopy("a", 4)
	fmt.Println(thingy)
	// Output: aaaa
}
