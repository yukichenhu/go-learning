package main

import (
	"encoding/json"
	"fmt"
)

/**
json序列化测试
*/

type Role struct {
	RoleId   int
	RoleName string
}

type Position struct {
	PositionId   int
	PositionName string
}

type Employee struct {
	Name     string
	Age      int
	Role     Role
	Position Position
}

func main() {
	employee := Employee{
		Name: "kitty",
		Age:  18,
		Role: Role{
			RoleId:   1,
			RoleName: "测试",
		},
		Position: Position{
			PositionId:   1,
			PositionName: "中级",
		},
	}

	jsonData, _ := json.Marshal(employee)
	println(len(jsonData))
	fmt.Printf("%+v\n", employee)
	println(string(jsonData))
}
