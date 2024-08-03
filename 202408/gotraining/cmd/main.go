package main

import (
	"fmt"
	cal "gotraining/internal/calculator"
	"gotraining/pkg/utils"
)

func main() {
	c := cal.Calculator{}
	fmt.Println(c.Add(1, 2))
	fmt.Println(c.Subtract(3, 2))
	fmt.Println(utils.IsEven(3))

	//r := gin.Default()
	//
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "GET",
	//	})
	//})
	//
	//err := r.Run(":9090")
	//if err != nil {
	//	return
	//}
}
