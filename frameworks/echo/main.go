package main

import(
  "fmt"
  "net/http"
  "strconv"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"

   _ "github.com/go-sql-driver/mysql"
  "github.com/gocraft/dbr"
)

type(
  userinfo struct { 
    ID        int    `db:"id"`
    Email     string `db:"email"`
    FirstName string `db:"first_name"`
    LastName  string `db:"last_name"`
  }

  userinfoJSON struct{
    ID        int `json:"id"`
    Email     string `json:"email"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
  }

  responseData struct{
    Users []userinfo `json:"users"`
  }
)

var(
  tablename = "userinfo"
  seq = 1
  conn, _ = dbr.Open("mysql", "root:@tcp(127.0.0.1:3306)/gorm", nil)
  sess = conn.NewSession(nil)
)

func main(){
  echo := echo.New()

  echo.Use(middleware.Logger())
  echo.Use(middleware.Recover())

  echo.POST("/users/", insertUser)
  echo.GET("/users/", selectUsers)
  echo.GET("/users/:id", selectUser)
  echo.PUT("/users/", updateUser)
  echo.DELETE("/users/:id", deleteUser)

  echo.Logger.Fatal(echo.Start(":1323"))
}

// Handlers

func selectUsers(c echo.Context) error{
  var u []userinfo

  sess.Select("*").From(tablename).Load(&u)
  fmt.Println(&u)
  response := new(responseData)
  response.Users = u
  return c.JSON(http.StatusOK, response)
}

func selectUser(c echo.Context) error {
  var userinfo userinfo

  id := c.Param("id")
  sess.Select("*").From(tablename).Where("id = ?", id).Load(&userinfo)

  return c.JSON(http.StatusOK, userinfo)
}

func insertUser(c echo.Context) error {
  u := new(userinfoJSON)

  if err := c.Bind(u); err != nil{
    return err
  }

  sess.InsertInto(tablename).Columns("id", "email", "first_name", "last_name").Values(u.ID, u.Email, u.FirstName, u.LastName).Exec()

  return c.NoContent(http.StatusOK)
}

func updateUser(c echo.Context) error{
  u := new(userinfoJSON)

  if err := c.Bind(u); err != nil{
    return err
  }

  attrsMap := map[string]interface{}{"id": u.ID, "email": u.Email, "first_name": u.FirstName, "last_name": u.LastName}
  sess.Update(tablename).SetMap(attrsMap).Where("id = ?", u.ID).Exec()
  return c.NoContent(http.StatusOK)
}

func deleteUser(c echo.Context) error {
  id, _ := strconv.Atoi(c.Param("id"))

  sess.DeleteFrom(tablename).Where("id = ?", id).Exec()

  return c.NoContent(http.StatusNoContent)
}
