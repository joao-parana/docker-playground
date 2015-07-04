package main

import (
  "fmt"
  "log"
  "io/ioutil"
  "net/http"
  "os"
  "time" 
)

type Hello struct{}

func main() {
  fmt.Println("Seja bemvindo aluno da Linux Solutions.")
  fmt.Println("Este é um curso de Docker.")
  fmt.Println("Visite o Blog do Instrutor em http://joao-parana.com.br")
  
  time.AfterFunc(3 * time.Second, getLocalContent)

  time.AfterFunc(9 * time.Second, getBlogContent)
  
  var h Hello
  wserr := http.ListenAndServe("localhost:4500", h)
  if wserr != nil {
    log.Fatal(wserr)
  }
}

func getBlogContent() {
  getContent(os.Getenv("blog_addr"))
}

func getLocalContent() {
  getContent("http://localhost:4500")
}

func getContent(url string) {
  resp, err := http.Get(url)
  check(err)
  body, err := ioutil.ReadAll(resp.Body)
  check(err)
  fmt.Println(len(body), " bytes lidos em ", url)
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
