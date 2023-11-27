package main

import (
	"fmt"

	"github.com/zhengyi13/go-ddns/domain"
	"github.com/zhengyi13/go-ddns/ipinfo"
	"github.com/zhengyi13/go-ddns/record"
)

func main() {

	ip, err := ipinfo.New().Get()
	if err != nil {
		fmt.Printf("ip get error: %v\n", err)
	}

	a := record.New("A", "www", domain.New("deprioritize.me"))
	fmt.Printf("Set record %v to IP: %s\n", a, ip)
}
