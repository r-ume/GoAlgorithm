package main

import(
  "github.com/ant0ine/go-json-rest/rest"
  "log"
  "net/http"
  "sync"
  "fmt"
)

var store = map[string]*Country{}
var lock = sync.RWMutex{}

func main(){
  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)

  router, err := rest.MakeRouter(
    rest.Get("/countries", GetAllCountries),
  )

  if err != nil{
    log.Fatal(err)
  }

  api.SetApp(router)
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type Country struct{
  Code string
  Name string
}

func GetAllCountries(w rest.ResponseWriter, r *rest.Request){
  lock.RLock()
  countries := make([]Country, len(store))
  index := 0

  for _, country := range store{
    countries[index] = *country
    index++
  }

  lock.RUnlock()
  w.WriteJson(&countries)
}
