package main

import(
  "github.com/ant0ine/go-json-rest/rest"
  "log"
  "net/http"
)

func main(){
  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)
  api.Use(&rest.AuthBasicMiddleware{
    Realm: "test zone",
    Authenticator: func(userId string, password string) bool{
      if userId == "admin" && password == "admin"{
        return true
      }

      return false
    },
  })

  api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request){
    w.WriteJson(map[string]string{"Body": "Hello Basic Auth!"})  
  }))

  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
