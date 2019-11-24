package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main () {
  fmt.Print("=======================Menu De Calculos=====================")
  fmt.Println("LAS OPCIONES DE CALCULO SON -> " ,
  "\nCalculo Hipotenusa -> 1 \nCalculo Raiz Cuadrada -> 2 \nCalculo Distancia Entre Dos Puntos -> 3 \nCalculo Formula Cuadratica -> 4")

  //para entrada de datos
  reader := bufio.NewReader(os.Stdin)
  entrada , _ := reader.ReadString('\n')
  opcion := strings.TrimRight(entrada , "\r\n")

  if (opcion == 1) {
    fmt.Println("You are selected the first option that is Hipotenuse Calculus.")
  }


}
