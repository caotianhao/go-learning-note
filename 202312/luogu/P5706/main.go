package main

import "fmt"

func main() {
	var total float64
	var stu int
	_, _ = fmt.Scanf("%f %d", &total, &stu)
	fmt.Printf("%.3f\n%d", total/float64(stu), stu*2)
}
