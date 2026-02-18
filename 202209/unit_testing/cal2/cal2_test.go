package cal2

import "testing"

func TestProductCal(t *testing.T) {
	res := productCal(10)
	if res != 3628800 {
		t.Fatalf("Err!Err!Err!")
	}
	t.Logf("Yes")
}
