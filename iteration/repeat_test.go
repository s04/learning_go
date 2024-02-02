package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat5(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeat999(t *testing.T) {
	repeated := Repeat("a", 99)
	expected := ""

	for i := 0; i < 99; i++ {
		expected += "a"
	}

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleTest() {
	var r = Repeat("a", 10)

	fmt.Println(r)
	// Output: aaaaaaaaaa
}

// func BenchmarkRepeat(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Repeat("a")
// 	}
// }
