package main

import ("net/http")

func main () {
  //routes -> example -> localhost:4500/route_name

  /*
    estructura

    http.HandleFunc("nombre de la ruta" ,)
  */
  http.HandleFunc("/" , homeHandler)
  http.HandleFunc("/contact" , contactHandler)
  //escucha y envia una request iniciando el servidor
  http.ListenAndServe(":4500" , nil)
}


//funcion pedida desde arriba
func homeHandler(response http.ResponseWriter , request *http.Request) {
  //ojo aqui por alguna razon es arreglo de bytes sera por serializacion o no se
  response.Write([]byte("Home page Using Golang how backend......"))
}

func contactHandler(response http.ResponseWriter , request *http.Request) {
  response.Write([]byte("Contact Page Using Go How Backend......!"))
}
