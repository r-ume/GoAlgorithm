package main

import(
  "github.com/ant0ine/go-json-rest/rest"
  "log"
  "net/http"
)

func main() {
  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)
  router, err := rest.MakeRouter(
    rest.Post("/hello", postHello),
  )

  if err != nil{
    log.Fatal(err)
  }

  log.Printf("Server started")
  api.SetApp(router)
  http.ListenAndServe(":9999", api.MakeHandler())
}

type postHelloInput struct{
  Name string
}

type postHelloOutput struct{
  Result string
}

func postHello(w rest.ResponseWriter, req *rest.Request){
  input := postHelloInput{}
  err := req.DecodeJsonPayload(&input)

  if err != nil{
    rest.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if input.Name == ""{
    rest.Error(w, "Name is required", 400)
    return
  }

  log.Printf("%#v", input)

  w.WriteJson(&postHelloOutput{
    "Hello" + input.Name,
  })
}
