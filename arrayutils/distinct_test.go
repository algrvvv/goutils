package arrayutils

import (
	"reflect"
	"testing"
)

func TestDistinct(t *testing.T) {
	inputData := []int{1, 2, 3, 3, 4, 5, 5, 6, 7, 8, 8, 9, 9, 10}
	actual := Distinct(inputData)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v\nexpected %v", actual, expected)
	}
}
