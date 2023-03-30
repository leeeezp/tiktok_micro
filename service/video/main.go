package main

import (
	"log"
	video "tiktok_micro/kitex_gen/video/videorpc"
)

func main() {
	svr := video.NewServer(new(VideoRpcImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
