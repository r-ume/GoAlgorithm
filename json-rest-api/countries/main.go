package main

import(
  "github.com/ant0ine/go-json-rest/rest"
  "log"
  "net/http"
  "sync"
)

var store = map[string]*Country{}
var lock = sync.RWMutex{}

func main(){
  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)

  router, err := rest.MakeRouter(
    rest.Get("/countries", GetAllCountries),
    rest.Get("/countries/:code", GetCountry),
    rest.Post("/countries", PostCountry),
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

func GetCountry(w rest.ResponseWriter, r *rest.Request){
  code := r.PathParam("code")

  lock.RLock()
  var country *Country

  if store[code] != nil{
    country = &Country{}
    *country = *store[code]
  }
  lock.RUnlock()

  if country == nil{
    rest.NotFound(w, r)
    return
  }

  w.WriteJson(country)
}

func PostCountry(w rest.ResponseWriter, r *rest.Request){
  country := Country{}
  err := r.DecodeJsonPayload(&country)

  if err != nil{
    rest.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if country.Code == "" {
    rest.Error(w, "country code required", 400)
  }

  if country.Name == "" {
    rest.Error(w, "country name required", 400)
  }

  lock.Lock()
  store[country.Code] = &country
  lock.Unlock()

  w.WriteJson(&country)
}
