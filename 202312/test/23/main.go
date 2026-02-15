package main

/*
#include <stdio.h>

int rr(){
	printf("%d\n",22222);
}
*/
import "C"

func main() {
	C.rr()
}
