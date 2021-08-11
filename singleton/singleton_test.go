package pattern

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("A new connection object must have been made.")
	}

	expectedCounter := counter1

	currentCount := counter1.AddOne()

	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be equal to 1 but it is %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		t.Errorf("Singleton instance must be different")
	}

	currentCount = counter2.AddOne()

	if currentCount != 2 {
		t.Errorf("After calling for the second time to count, the count must be equal to 2 but it is %d\n", currentCount)
	}
}
