package observer

import (
	"fmt"
	"testing"
)

// TestObserver test observer
func TestObserver(t *testing.T) {
	random := NewRandomNumberGenerator()

	observer1 := &DigitObserver{random}
	observer2 := &DigitObserver{random}

	random.AddObserver(observer1)
	random.AddObserver(observer2)

	result := random.Execute()

	fmt.Println(result)

	for _, r := range result {
		if len(result) != 2 && r >= 50 {
			t.Errorf("Expect result to equal random int array")
		}
	}
}
