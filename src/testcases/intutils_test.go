package main

import "testing"

// TestIntMinBasic basic test case for IntMin
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// TestIntMinBasic basic test case for IntMin
func TestIntMinMultipleData(t *testing.T) {
	testData := []struct {
		no1, no2, want int
	}{
		{
			2, -2, -2,
		}, {
			-4, -2, -4,
		}, {
			-2, 2, -2,
		}, {
			0, 5, 0,
		}, {
			1, 10, 1,
		},
	}

	for _, data := range testData {

		ans := IntMin(data.no1, data.no2)
		if ans != data.want {
			t.Errorf("IntMin(%d, %d) = %d; want %d", data.no1, data.no2, ans, data.want)
		}
	}

	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}
