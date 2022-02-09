package stack

import (
	"testing"
)

var s = NewArrayStack()

func TestArrayStack_Push(t *testing.T) {
	//Act
	Push(s, "val1")
}

func TestArrayStack_Size(t *testing.T) {
	//Arrange
	expectedSize := 1

	//Act
	resultSize := Size(s)

	//Assert
	if resultSize != expectedSize {
		t.Errorf("Incorrect stack size. Expected: %d. Get: %d", expectedSize, resultSize)
	}
}

func TestArrayStack_Iterate(t *testing.T) {
	//Arrange
	expectedOutput := "Stack index: 0 value: val1 \n"

	//Act
	resultOutput := Iterate(s)

	//Assert
	if resultOutput != expectedOutput {
		t.Errorf("Incorect stack output. \n Expected: %s. \n Get: %s", expectedOutput, resultOutput)
	}
}

func TestArrayStack_Top(t *testing.T) {
	//Arrange
	expectedTop := "val1"

	//Act
	resultTop, err := Top(s)
	if err != nil {
		t.Errorf("Error while getting top value: %s", err)
	}

	//Assert
	if resultTop != expectedTop {
		t.Errorf("Incorrect stack size. Expected: %s. Get: %s", expectedTop, resultTop)
	}
}

func TestArrayStack_Pop(t *testing.T) {
	//Arrange
	expectedDelVal := "val1"

	//Act
	resultDeletedValue, err := Pop(s)
	if err != nil {
		t.Error("Got error while deleting: ", err)
	}

	//Assert
	if expectedDelVal != resultDeletedValue {
		t.Errorf("Incorrect stack size. Expected: %s. Get: %s", expectedDelVal, resultDeletedValue)
	}
}

func TestArrayStack_IsEmpty(t *testing.T) {
	//Arrange
	expected := false

	//Act
	result := IsEmpty(s)

	//Assert
	if result != expected {
		t.Errorf("Incorect IsEmpty bevave, Expected: %t. Get: %t", expected, result)
	}

}
