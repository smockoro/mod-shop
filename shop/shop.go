package shop

import (
	"fmt"

	u "github.com/smockoro/mod-user/v2/user"
)

func ShopHello() {
	fmt.Println("mod Shop v1.0.2")
	user := u.User{Name: "Mod Shop Gopher", Age: 20}
	user.UserHello()
}
