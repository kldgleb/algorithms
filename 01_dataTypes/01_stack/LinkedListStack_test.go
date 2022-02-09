package stack

import (
	"testing"
)

var testLinkedListStack = NewLinkedListStack()

func TestLinkedListStack_Push(t *testing.T) {
	//Act
	Push(testLinkedListStack, "val1")
}

func TestLinkedListStack_Size(t *testing.T) {
	//Arrange
	expectedSize := 1

	//Act
	resultSize := Size(testLinkedListStack)

	//Assert
	if resultSize != expectedSize {
		t.Errorf("Incorrect stack size. Expected: %d. Get: %d", expectedSize, resultSize)
	}
}

func TestLinkedListStack_Iterate(t *testing.T) {
	//Arrange
	expectedOutput := "Stack index: 0 value: val1 \n"

	//Act
	resultOutput := Iterate(testLinkedListStack)

	//Assert
	if resultOutput != expectedOutput {
		t.Errorf("Incorect stack output. \n Expected: %s. \n Get: %s", expectedOutput, resultOutput)
	}
}

func TestLinkedListStack_Top(t *testing.T) {
	//Arrange
	expectedTop := "val1"

	//Act
	resultTop, err := Top(testLinkedListStack)
	if err != nil {
		t.Errorf("Error while getting top value: %s", err)
	}

	//Assert
	if resultTop != expectedTop {
		t.Errorf("Incorrect stack size. Expected: %s. Get: %s", expectedTop, resultTop)
	}
}

func TestLinkedListStack_Pop(t *testing.T) {
	//Arrange
	expectedDelVal := "val1"

	//Act
	resultDeletedValue, err := Pop(testLinkedListStack)
	if err != nil {
		t.Error("Got error while deleting: ", err)
	}

	//Assert
	if expectedDelVal != resultDeletedValue {
		t.Errorf("Incorrect stack size. Expected: %s. Get: %s", expectedDelVal, resultDeletedValue)
	}
}

func TestLinkedListStack_IsEmpty(t *testing.T) {
	//Arrange
	expected := false

	//Act
	result := IsEmpty(s)

	//Assert
	if result != expected {
		t.Errorf("Incorect IsEmpty bevave, Expected: %t. Get: %t", expected, result)
	}

}
