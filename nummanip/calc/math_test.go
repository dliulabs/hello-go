package calc

import "testing"

func testMathAdd(t *testing.T) {
	_, result = Add(1, 2, 3)
	if result != 6 {
		t.Errorf("Add(1,2,3) failed, expected %v but got %v", 6, result)
	} else {
		t.Logf("Add(1,2,3) PASSED, expected %v, and got %v", 6, results)
	}
}
