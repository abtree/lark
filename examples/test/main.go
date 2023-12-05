package main

import "lark/pb"

func main() {
	data := &pb.UserInfo{
		Id: 1,
	}
	bs := []byte(data)
}
