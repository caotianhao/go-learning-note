package main

/*
#include <stdio.h>
int main1() {
    printf("test for cgo\n");
	return 0;
}
*/
import "C"

func main() {
	C.main1()
}
