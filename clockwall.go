package main

import (
	"os"
	"strings"
	"net"
	"log"
	"fmt"
)

func main() {
	//port := flag.String("Tokyo", "", "port number")
	//flag.Parse()
	//fmt.Println(*port)

	args := os.Args
	for _, arg := range args {
		serverInfo := strings.Split(arg, "=")
		if len(serverInfo) == 2 {
			name := serverInfo[0]
			address := serverInfo[1]
			c := make(chan string)
			go NetCat(name, address, c)
			fmt.Println(<-c)
		}

	}


}


func NetCat(name, address string, c chan string) {
	fmt.Println("server name: " + name)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c<-"123"
}