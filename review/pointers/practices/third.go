package practices

import "fmt"

type User struct {
	Name string
}

func third() {
	user = User{Name: "BeforeMyFunc"}
	fmt.Println("user.Nameの値(呼び出し前)：", user.Name)

	myFunc(&user)
	fmt.Println("user.Nameの値(呼び出し後)：", user.Name)
}

func myFunc(user *User) {
	user.Name = "pointer"
	fmt.Println("user.Nameの値(関数内)：", user.Name)
}
