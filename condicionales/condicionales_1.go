package main
import (
  "fmt"
)

func main () {
  //mayor de tres numeros
  var numberOne = 500
  var numberTwo = 500
  var numberThree = 500

  if (numberOne > numberTwo && numberOne > numberThree) {
    fmt.Println("El mayor de los numers es -> " , numberOne)
  }
  if (numberTwo > numberOne && numberTwo > numberThree) {
    fmt.Println("El mayor de los numers es -> " , numberTwo)
  }
  if (numberThree > numberOne && numberThree > numberTwo) {
    fmt.Println("El mayor de los numers es -> " , numberThree)
  }
  if (numberOne == numberTwo && numberOne == numberThree) {
    fmt.Println("The Numbers are the same")
  }
}
