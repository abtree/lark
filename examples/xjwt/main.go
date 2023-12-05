package main

import (
	"encoding/json"
	"fmt"
	"lark/com/pkgs/xjwt"
	"strings"
)

func main() {
	t, err := xjwt.CreateToken([]byte("123"), true, 3600)
	if err != nil {
		fmt.Println("Create JWT Token Error", err.Error())
		return
	}
	fmt.Println("Create JWT Token success", t.Token)

	dt, err := xjwt.Decode(strings.TrimPrefix(t.Token, xjwt.JWT_PREFIX))
	if err != nil {
		fmt.Println("Decode JWT Token Error", err.Error())
		return
	}
	fmt.Println("Decode JWT Token Success")
	byts, _ := json.Marshal(dt)
	fmt.Println(string(byts))
}
