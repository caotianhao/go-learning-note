package park

import (
	"testing"
	"time"
)

func TestPark(t *testing.T) {
	p := NewPark(parkCap)

	stopCh := make(chan struct{})
	defer close(stopCh)

	go RandomPark(p, stopCh)

	// ctx 学的不太好，暂时这样
	time.Sleep(11 * time.Second)

	t.Logf("\n求求了，多停会啊，一共才赚了 %.2f 元\n", p.GetTotalMoney())
}
