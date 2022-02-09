package queue

import "testing"

var testArrayQueue = NewArrayQueue()

func TestArrayQueue_Enqueue(t *testing.T) {
	Enqueue(testArrayQueue, "val1")
	Enqueue(testArrayQueue, "val2")
}

func TestArrayQueue_Dequeue(t *testing.T) {
	expectedDelVal := "val1"

	resultDelVal, err := Dequeue(testArrayQueue)
	if err != nil {
		t.Errorf("Fail while Dequeueuueue: %s", err)
	}

	if expectedDelVal != resultDelVal {
		t.Errorf("Dequeue test failed. Expected: %s. Get: %s", expectedDelVal, resultDelVal)
	}
}

func TestArrayQueue_IsEmpty(t *testing.T) {
	expected := false
	result := IsEmpty(testArrayQueue)
	if expected != result {
		t.Errorf("Queueu IsEmpty test failed. Expected: %t. Get: %t", expected, result)
	}
}

func TestArrayQueue_Size(t *testing.T) {
	expectedSize := 1
	resultSize := Size(testArrayQueue)
	if expectedSize != resultSize {
		t.Errorf("ArrayQueue Size test failed. Expected: %d. Get: %d", expectedSize, resultSize)
	}
}

func TestArrayQueue_Iterate(t *testing.T) {
	expectedOutput := "Queue index: 0 value: val2 \n"
	resultOutput := Iterate(testArrayQueue)
	if resultOutput != expectedOutput {
		t.Errorf("Incorect ArrayQueue output. \n Expected: %s. \n Get: %s", expectedOutput, resultOutput)
	}
}
