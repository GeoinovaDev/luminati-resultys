package main

import (
	"fmt"

	"github.com/GeoinovaDev/luminati-resultys/luminati"
)

func main() {
	p := luminati.CreatePool()

	url1 := luminati.CreateURL("user1", "pass1", "host1", "port1")
	url2 := luminati.CreateURL("user2", "pass2", "host2", "port2")
	url3 := luminati.CreateURL("user3", "pass3", "host3", "port3")
	url4 := luminati.CreateURL("user4", "pass4", "host4", "port4")

	p.Add(url1)
	p.Add(url2)
	p.Add(url3)
	p.Add(url4)

	p.RemoveTemporary(url2, 5)

	fmt.Scanln()
}
