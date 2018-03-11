package unitt

import "testing"
import "math/rand"
import "time"

func random() int {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)

	return rand.Intn(100)
}

func TestPow(t *testing.T) {
	x := random()

	t.Logf("Testing Pow with %d", x)

	if powerOfX := Pow(x); powerOfX != x*x {
		t.Errorf("Expected the power of x to equal %d", powerOfX)
	}
}
