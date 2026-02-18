package selectRandomly

func selectRandomly(ch chan int) {
	go func() {
		for {
			select {
			case ch <- 1:
			case ch <- 2:
			case ch <- 3:
			}
		}
	}()
}
