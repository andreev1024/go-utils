package utils_test

import (
	"fmt"
	"reflect"
	"testing"

	u "github.com/andreev1024/go-utils"
)

func TestSliceDiff(t *testing.T) {

	var intTests = []struct {
		a        []int
		b        []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 20, 3, 40, 5}, []int{2, 4}},
		{[]int{1, 2, 3, 4, 5}, []int{}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, []int{}},
	}

	for i := range intTests {
		t.Run(fmt.Sprintf("int test %d", i), func(t *testing.T) {
			if !reflect.DeepEqual(u.SliceDiff(intTests[i].a, intTests[i].b), intTests[i].expected) {
				t.Error("Test failed")
			}
		})
	}

	var stringTests = []struct {
		a        []string
		b        []string
		expected []string
	}{
		{[]string{"a", "b", "c"}, []string{"a", "x", "c"}, []string{"b"}},
	}

	for i := range stringTests {
		t.Run(fmt.Sprintf("string test %d", i), func(t *testing.T) {
			if !reflect.DeepEqual(u.SliceDiff(stringTests[i].a, stringTests[i].b), stringTests[i].expected) {
				t.Error("Test failed")
			}
		})
	}
}
