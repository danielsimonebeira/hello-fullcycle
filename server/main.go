package main

import (
  "fmt"
  "net/http"
  )

func main()  {

  // Apresenta no console
  fmt.Printf("Hello, Full Cycle\n")

  // Acessa via navegador
  http.HandleFunc("/hello", HelloHandler)
  http.ListenAndServe(":8887", nil)


}

func HelloHandler(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(w, "Hello Full Cycle")
}
