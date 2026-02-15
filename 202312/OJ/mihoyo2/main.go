package main

import "fmt"

const fight = float64(1) / float64(6)

func main() {
	cards, hp := 0, 0
	_, _ = fmt.Scanf("%d %d", &cards, &hp)
	skill, x := 0, 0
	coin := 0
	for i := 0; i < cards; i++ {
		_, _ = fmt.Scanf("%d %d", &skill, &x)
		if skill == 1 {
			coin += x
		} else {
			hp -= x
			if hp <= 0 {
				fmt.Println(1)
				//return
			} else {
				fmt.Println(fight)
				//return
			}
		}
	}
}
