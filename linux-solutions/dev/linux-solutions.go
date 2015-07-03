package main

import (
  "fmt"
  "log"
  "io/ioutil"
  "net/http"
  "os"
)

type Hello struct{}

func main() {
  fmt.Println("Seja bemvindo aluno da Linux Solutions.")
  fmt.Println("Este é um curso de Docker.")
  fmt.Println("Visite o Blog do Instrutor em http://joao-parana.com.br")
  resp, err := http.Get("http://joao-parana.com.br")
  check(err)
  body, err := ioutil.ReadAll(resp.Body)
  check(err)
  fmt.Println(len(body), " bytes lidos em http://joao-parana.com.br")
  var h Hello
  wserr := http.ListenAndServe("localhost:4500", h)
  if wserr != nil {
    log.Fatal(wserr)
  }
}

func check(err error) {
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func (h Hello) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request) {
  fmt.Fprint(w, "Hello!")
}
