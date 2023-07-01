package main

import (
	"fmt"
)

func main() {
	
	var n int
	fmt.Scanf("%d", &n)
	var arr [200] int
  
	for i := 0; i < n; i++ {
	  fmt.Printf("\nEnter %d:", i)
	  fmt.Scanf("%d", arr[i])
    }
}