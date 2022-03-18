package priorityQueue

import (
	"testing"
)

var h = NewHeap()

func TestHeap_Insert(t *testing.T) {
	heapExpectedOutput := []int{0, 8, 6, 2, 1, 4}

	Insert(h, 1)
	Insert(h, 4)
	Insert(h, 2)
	Insert(h, 8)
	Insert(h, 6)
	heapResultOutput := h.slice

	for i := 0; i < len(heapResultOutput); i++ {
		if heapExpectedOutput[i] != heapResultOutput[i] {
			t.Errorf("TestHeap_Insert Fail. Iterattion: %d Expected: %d. Get: %d",
				i, heapExpectedOutput[i], heapResultOutput[i])
		}
	}
}

func TestHeap_DelMax(t *testing.T) {
	expectedDelVal1 := 8
	//expectedDelVal2 := 6
	//expectedDelVal3 := 4

	resultDelVal1, err := DelMax(h)
	if err != nil {
		t.Errorf("test Heap err while DelMax: %s", err)
	}
	//resultDelVal2, err := DelMax(h)
	//if err != nil {
	//	t.Errorf("test Heap err while DelMax: %s", err)
	//}
	//resultDelVal3, err := DelMax(h)
	//if err != nil {
	//	t.Errorf("test Heap err while DelMax: %s", err)
	//}

	if expectedDelVal1 != resultDelVal1 {
		t.Errorf("TestHeap_DelMax Fail. Expected: %d. Get: %d", expectedDelVal1, resultDelVal1)
	}
	//if expectedDelVal2 != resultDelVal2 {
	//	t.Errorf("TestHeap_DelMax Fail. Expected: %d. Get: %d", expectedDelVal2, resultDelVal2)
	//}
	//if expectedDelVal3 != resultDelVal3 {
	//	t.Errorf("TestHeap_DelMax Fail. Expected: %d. Get: %d", expectedDelVal3, resultDelVal3)
	//}
}

func TestHeap_IsEmpty(t *testing.T) {
	resultBool := IsEmpty(h)
	if false != resultBool {
		t.Errorf("TestHeap_DelMax Fail. Expected: %t. Get: %t", false, resultBool)
	}
}

func TestHeap_Size(t *testing.T) {
	expectedSize := 2
	resultSize := Size(h)
	if expectedSize != resultSize {
		t.Errorf("TestHeap_Size Fail. Expected: %d. Get: %d", expectedSize, resultSize)
	}
}

func TestHeap_Max(t *testing.T) {
	expectedMax := 2

	resultMax, err := Max(h)
	if err != nil {
		t.Errorf("TestHeap_Max error")
	}
	if expectedMax != resultMax {
		t.Errorf("TestHeap_DelMax Fail. Expected: %d. Get: %d", expectedMax, resultMax)
	}
}
