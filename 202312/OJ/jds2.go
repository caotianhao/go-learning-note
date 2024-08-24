package main

import "fmt"

// 30%
func main() {
	n, m := 0, 0
	_, _ = fmt.Scanf("%d %d", &n, &m)
	s, force := " ", 0
	// humanOrMonster 存储人和怪物，下标为对应 id
	// 若为人，则值为 force
	// 若为怪兽，则值为 -force
	isDied := make([]bool, n)
	humanOrMonster := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%s %d", &s, &force)
		if s == "human" {
			humanOrMonster[i] = force
		} else {
			humanOrMonster[i] = -force
		}
	}
	//fmt.Println("hom", humanOrMonster)
	against := make([][4]interface{}, 0)
	var id1, id2 int
	var b1, b2 byte
	for i := 0; i < m; i++ {
		_, _ = fmt.Scanf("%d %d %c %c\n", &id1, &id2, &b1, &b2)
		against = append(against, [4]interface{}{id1, id2, b1, b2})
		//fmt.Println("ag", i, against)
	}
	for _, v := range against {
		fight1, fight2 := v[0].(int), v[1].(int)
		//fmt.Println(fight1, fight2)
		// 如果一人一兽，且都没死
		if humanOrMonster[fight1-1]*humanOrMonster[fight2-1] < 0 && !isDied[fight1-1] && !isDied[fight2-1] {
			// 第一个是人，第二个是兽
			if humanOrMonster[fight1-1] > 0 {
				// 人状态是 y，不管兽是什么都会战斗
				if v[2].(uint8) == 89 {
					// 人大于兽，兽死
					if humanOrMonster[fight1-1] > -humanOrMonster[fight2-1] {
						isDied[fight2-1] = true
					} else if humanOrMonster[fight1-1] == -humanOrMonster[fight2-1] {
						// 人等于兽，同死
						isDied[fight1-1] = true
						isDied[fight2-1] = true
					} else {
						// 人小于兽，人死
						isDied[fight1-1] = true
					}
				} else {
					// 人状态是 n，兽 y
					if v[3].(uint8) == 89 {
						if humanOrMonster[fight1-1] > -humanOrMonster[fight2-1] {
							isDied[fight2-1] = true
						}
					}
				}
			} else {
				// 第一个是兽，第二个是人
				// 人状态是 y
				if v[3].(uint8) == 89 {
					if -humanOrMonster[fight1-1] > humanOrMonster[fight2-1] {
						isDied[fight2-1] = true
					} else if -humanOrMonster[fight1-1] == humanOrMonster[fight2-1] {
						isDied[fight1-1] = true
						isDied[fight2-1] = true
					} else {
						isDied[fight1-1] = true
					}
				} else {
					// 人状态是 n，兽 y
					if v[2].(uint8) == 89 {
						if -humanOrMonster[fight1-1] < humanOrMonster[fight2-1] {
							isDied[fight1-1] = true
						}
					}
				}
			}
		}
		//fmt.Println(isDied)
	}
	for _, v := range isDied {
		if v {
			fmt.Printf("%c", 'N')
		} else {
			fmt.Printf("%c", 'Y')
		}
	}
}
