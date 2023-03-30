package main

import (
	"log"
	relation "tiktok_micro/kitex_gen/relation/userrpc"
)

func main() {
	svr := relation.NewServer(new(UserRpcImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
