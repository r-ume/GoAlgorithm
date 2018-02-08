// GORM is a simple ORM library for GO.

package main

import(
  "github.com/ant0ine/go-json-rest/rest"
  _ "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
  "log"
  "net/http"
  "time"
)

func main(){
  i := Impl{}
  i.InitDB()
  i.InitSchema()

  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)
  router, err := rest.MakeRouter(
    rest.Post("/reminders", i.PostReminder),
  )

  if err != nil{
    log.Fatal(err)
  }

  api.SetApp(router)
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type Reminder struct{
  Id        int64     `json:"id"`
  Message   string    `sql:"size:1024" json:"message"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
  DeleteAt  time.Time `json:"-"`
}

type Impl struct{
  DB *gorm.DB
}

func (i *Impl) InitDB(){
  var err error
  i.DB, err = gorm.Open("mysql", "root:@/gorm?charset=utf8&parseTime=true")

  if err != nil{
    log.Fatalf("Got error when connect database, the error is '%v'", err)
  }
  i.DB.LogMode(true)
}

func (i *Impl) InitSchema(){
  i.DB.AutoMigrate(&Reminder{})
}

func (i *Impl) PostReminder(w rest.ResponseWriter, r *rest.Request){
  reminder := Reminder{}
  if err := r.DecodeJsonPayload(&reminder); err != nil{
    rest.Error(w, err.Error(), http.StatusInternalServerError)
  }

  if err := i.DB.Save(&reminder).Error; err != nil{
    rest.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.WriteJson(&reminder)
}


