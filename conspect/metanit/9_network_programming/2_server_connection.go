package networkprogramming

import (
	"fmt"
	"io"
	"net"
	"os"
)

func ExampleServerConnectionProcess() {
	//server() // билдить и тестить отдельной программой
	//client() // билдить и тестить отдельной программой
}

func server() {
	message := "Hello, I am a server" // отправляемое сообщение
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(message))
		conn.Close()
	}
}

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	io.Copy(os.Stdout, conn)
	fmt.Println("\nDone")
}
