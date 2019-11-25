package main

import (
  "fmt"
  "math"
)

func main () {
  fmt.Print("=======================Menu De Calculos=====================")
  fmt.Println("LAS OPCIONES DE CALCULO SON -> " ,
  "\nCalculo Hipotenusa -> 1 \nCalculo Raiz Cuadrada -> 2 \nCalculo Distancia Entre Dos Puntos -> 3")
  
  var opcion int
  fmt.Println("Digite una Opcion de Calculo -> ")
  fmt.Scanln(&opcion)
  if (opcion == 1) {
    var a float64
    var b float64
    var c float64
    fmt.Println("Please Digit the a value -> ")
    fmt.Scanln(&a)
    fmt.Println("Please Digit the b value -> ")
    fmt.Scanln(&b)
    c = math.Sqrt(a*a + b*b)
    fmt.Println("The Hipotenuse Or The C Value is -> " , c)
  } else {
    if (opcion == 2) {
      var number float64
      fmt.Println("Please Digit a number -> ")
      fmt.Scanln(&number)
      root := math.Sqrt(number)
      fmt.Println("The Root of this number is -> " , root)
    } else {
      var x1 float64
      var x2 float64
      var y2 float64
      var y1 float64
      var point_distance float64
      fmt.Println("Please digit the first coordinate of axis x -> ")
      fmt.Scanln(&x1)
      fmt.Println("Please Digit the second coordinate of axis x -> ")
      fmt.Scanln(&x2)
      fmt.Println("Please digit the first coordinate of axis y -> ")
      fmt.Scanln(&y1)
      fmt.Println("Please digit the second coordinate of axis y -> ")
      fmt.Scanln(&y2)
      point_distance = math.Sqrt(math.Pow( (x2 - x1) , 2) + math.Pow( (y2 - y1),2))
      fmt.Println("The Distance of Points is -> " , point_distance)
      /*
        x2 -> -7
        x1 -> -3
        y2 -> 1
        y1 -> 5
      */
    }
  }
}
