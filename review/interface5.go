package main

import "fmt"

type JSONable interface {
	JSON() string
}

type User struct {
	Id   int
	Name string
}

func (s *User) JSON() string {
	return fmt.Sprintf(`{ "Id": %d, "Name": "%s" }`, s.Id, s.Name)
}

type AdminUser struct {
	User
	Admin bool
}

func (s *AdminUser) JSON() scdtring {
	return fmt.Sprintf(`{ "Id": %d, "Name": "%s", "Admin": "%v" }`, s.Id, s.Name, s.Admin)
}

func main() {
	var user JSONable = &User{1, "r-ume"}
	fmt.Println(user.JSON())

	var adminUser JSONable = &AdminUser{User{0, "admin"}, true}
	fmt.Println(adminUser.JSON())

	jsonable, ok := adminUser.(JSONable)
	if ok {
		// 実装してる場合
		fmt.Printf("JSON(): %s\n", jsonable.JSON())
	} else {
		// 実装してない場合
		fmt.Printf("Not JSONable\n")
	}
}
