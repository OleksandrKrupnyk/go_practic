/**
List 1.3
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	_, err = fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println(status)
}
