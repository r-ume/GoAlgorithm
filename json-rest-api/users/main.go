package main

import(
  "fmt"
  "log"
  "net/http"
  "sync"
  "github.com/ant0ine/go-json-rest/rest"
)

func main(){
  users := Users{
    Store: map[string]*User{},
  }

  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)
  router, err := rest.MakeRouter(
    rest.Post("/users", users.PostUser),
  )

  if err != nil{
    log.Fatal(err)
  }

  api.SetApp(router)
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type User struct{
  Id   string
  Name string
}

type Users struct{
  sync.RWMutex
  Store map[string]*User
}

func (users *Users) PostUser(w rest.ResponseWriter, r *rest.Request){
  user := User{}
  err := r.DecodeJsonPayload(&user)
  if err != nil{
    rest.Error(w, err.Error(), http.StatusInternalServerError)
  }

  users.Lock()
  id := fmt.Sprintf("%d", len(users.Store))
  user.Id = id
  users.Store[id] = &user
  users.Unlock()
  w.WriteJson(&user)
}
