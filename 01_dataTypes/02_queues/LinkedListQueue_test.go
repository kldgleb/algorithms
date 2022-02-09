package queue

import "testing"

var testLinkedListQueue = NewLinkedListQueue()

func TestLinkedListQueue_Enqueue(t *testing.T) {
	Enqueue(testLinkedListQueue, "val1")
	Enqueue(testLinkedListQueue, "val2")
}

func TestLinkedListQueue_Dequeue(t *testing.T) {
	expectedDelVal := "val1"

	resultDelVal, err := Dequeue(testLinkedListQueue)
	if err != nil {
		t.Errorf("Fail while Dequeueuueue: %s", err)
	}

	if expectedDelVal != resultDelVal {
		t.Errorf("Dequeue test failed. Expected: %s. Get: %s", expectedDelVal, resultDelVal)
	}
}

func TestLinkedListQueue_IsEmpty(t *testing.T) {
	expected := false
	result := IsEmpty(testLinkedListQueue)
	if expected != result {
		t.Errorf("Queueu IsEmpty test failed. Expected: %t. Get: %t", expected, result)
	}
}

func TestLinkedListQueue_Size(t *testing.T) {
	expectedSize := 1
	resultSize := Size(testLinkedListQueue)
	if expectedSize != resultSize {
		t.Errorf("ArrayQueue Size test failed. Expected: %d. Get: %d", expectedSize, resultSize)
	}
}

func TestLinkedListQueue_Iterate(t *testing.T) {
	expectedOutput := "Queue index: 0 value: val2 \n"
	resultOutput := Iterate(testLinkedListQueue)
	if resultOutput != expectedOutput {
		t.Errorf("Incorect ArrayQueue output. \n Expected: %s. \n Get: %s", expectedOutput, resultOutput)
	}
}
