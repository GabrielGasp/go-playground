package main

import (
	"bench/forwarded"
	"fmt"
)

func main() {
	realIP := forwarded.GetIpWithoutSplit(`for=192.168.0.1, for=192.168.0.2;host=api.homolog.boongcloud.net;proto=https`)
	fmt.Println(realIP)
}
