/*
  variables son de tipo estatico

    tipos de declaracion:
      - var nombre_variable tipo_variable = valor
      - nombre_variable := valor

  EN GO SI DECLARAS UNA VARIABLE DEBES USARLA OBLIGATORIAMENTE
*/
package main
import (
  "fmt"
)

var (
  nombre = "Bryan Ariel"
  apellido = "Sanchez Anariba"
  edad = 22
)
func main () {

  //example one
  fmt.Println()
  fmt.Println("Declaracion de una variable String: ")
  var language string = "Spanish"
  fmt.Println("Aprendiendo El Lenguaje -> " + language)
  fmt.Println()

  //example two
  fmt.Println("Suma de Dos Numeros: ")
  var number_1 int = 10;
  var number_2 int = 10;
  var res int = (number_1 + number_2)
  fmt.Println(number_1 , " + " , number_2 , " = " , res)
  fmt.Println()

  //example three using var(nombre , apellido , edad)
  fmt.Println("===================Datos de una Persona Declarada Como Objeto================")
  fmt.Println("Nombre -> " , nombre , "\nApellido -> " , apellido , "\nEdad -> " , edad)
  fmt.Println()

  //DECLARACION CORTA -> VALIDA SOLO DENTRO DE UNA FUNCION DE LO CONTRARIO NO FUNCIONA
  fmt.Println("Declaracion Corta de Variables , valido solo dentro de una funcion de lo contrario tira error")
  nombrePersona := "Bryan Ariel Sanchez Anariba"
  fmt.Println(nombrePersona)
  nombrePersona1 , apellidoPersona := "Maria Fernanda" ,"Sanchez Anariba"
  fmt.Println(nombrePersona1 , apellidoPersona)
  fmt.Println()

  //CONSTANTES -> IMPORTANTA

  const pi = 3.141516
  fmt.Println()
  fmt.Println("Imprimiendo una constante -> " , pi)
  fmt.Println()
}
