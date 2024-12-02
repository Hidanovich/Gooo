package main

import (
 "fmt"
 "math"
)

func main() 
 var a, b, c float64
 fmt.Scanln(   a,    b,    c)

 if a == 0 
  fmt.Println
  return
 

 d := b*b - 4*a*c
 switch 
 case d > 0:
  x1, x2 := (-b+math.Sqrt(d))/(2*a), (-b-math.Sqrt(d))/(2*a)
  fmt.Printf("Корни уравнения: x1 =
 case d == 0:
  x := -b / (2 * a)
  fmt.Printf("Уравнение имеет один корень: x =
 default:
  fmt.Println
 )