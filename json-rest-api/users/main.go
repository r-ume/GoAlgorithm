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
    rest.Get("/users", users.GetAllUsers),
    rest.Get("/users/:id", users.GetUser),
    rest.Post("/users", users.PostUser),
    rest.Put("/users/:id", users.PutUser),
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

func (users *Users) GetAllUsers (w rest.ResponseWriter, r *rest.Request){
  users.RLock()
  all_users := make([]User, len(users.Store))
  index := 0
  for _, user := range users.Store{
    all_users[index] = *user
    index++
  }
  users.RUnlock()
  w.WriteJson(&all_users)
}

func (users *Users) GetUser(w rest.ResponseWriter, r *rest.Request){
  id := r.PathParam("id")
  users.RLock()
  var user *User
  if users.Store[id] != nil{
    user = &User{}
    *user = *users.Store[id]
  }
  users.RUnlock()
  if user == nil{
    rest.NotFound(w, r)
    return
  }

  w.WriteJson(user)
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

func (users *Users) PutUser(w rest.ResponseWriter, r *rest.Request){
  id := r.PathParam("id")
  users.Lock()

  fmt.Println(users.Store[id])

  if users.Store[id] == nil{
    rest.NotFound(w, r)
    users.Unlock()
    return
  }

  user := User{}
  err := r.DecodeJsonPayload(&user)
  if err != nil{
    rest.Error(w, err.Error(), http.StatusInternalServerError)
    users.Unlock()
    return
  }

  user.Id = id
  users.Store[id] = &user
  users.Unlock()
  w.WriteJson(&user)
}
