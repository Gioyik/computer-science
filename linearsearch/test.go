package main

import "fmt"

func main() {
  var x = 5
  var z [7]int

  z[0] = 1
  z[1] = 2
  z[2] = 3
  z[3] = 4
  z[4] = 5
  z[5] = 6
  z[6] = 7

  for i:=0; i<x; i++ {
    if x == z[i] {
      fmt.Println(i)
    }
  }
}
