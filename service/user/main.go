package main

import (
	"log"
	user "tiktok_micro/kitex_gen/user/userrpc"
)

func main() {
	svr := user.NewServer(new(UserRpcImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
