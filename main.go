package main

import (
	"fmt"
	"github.com/zhengyi13/go-ddns/ipinfo"
)

func main() {

	ip, err := ipinfo.New().Get()
	if err != nil {
		fmt.Printf("ip get error: %v\n", err)
	}

	fmt.Printf("%s\n", ip)

	
}
