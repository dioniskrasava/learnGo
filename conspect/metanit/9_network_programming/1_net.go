package networkprogramming

import (
	"fmt"
	"io"
	"net"
	"os"
)

//-----------------------------------------------------------
//-----------------------------------------------------------
//
// О Т П Р А В К А    З А П Р О С О В
//
//-----------------------------------------------------------
//-----------------------------------------------------------

// Для отправки запросов к ресурсам в сети применяется функция net.Dial():
//
// func Dial(network, address string) (Conn, error)
//
//
// Эта функция принимает два параметра: network - тип протокола и address - адрес ресурса.
//
//
//----------------------------------------------------------------------------
// Есть следующие типы протоколов
//
//
//
// - tcp, tcp4, tcp6:
// протокол TCP. tcp по умолчанию представляет tcp4, цифра в конце указывает,
// какой тип адресов будет использоваться: IPv4 или IPv6
//
// - udp, udp4, udp6:
// протокол UDP. udp по умолчанию представляет udp4
//
// - ip, ip4, ip6:
// протокол IP. ip по умолчанию представляет ip4
//
// - unix, unixgram, unixpacket:
// сокеты Unix
//
//----------------------------------------------------------------------------
//
// Второй параметр представляет сетевой адрес ресурса (для адресов в сети интернет это домен).
// Это может быть числовой сетевой адрес, например, "127.0.0.1".
// Он может включать указание порта, например, "127.0.0.1:80".
// Это также может быть адрес в формате IPv6, например, "::1" или "[2516:b7f0:3421:b16::71]:80".
//
//---------------------------------------------------------------------------
//
// Функция возращает объект, который реализует интерфейс net.Conn. Этот интерфейс, в свою очередь,
// применяет интерфейсы  io.Reader и io.Writer, то есть может использоваться как поток для чтения
// и записи. Пакет net предоставляет базовые реализации этого интерфейса в виде типов IPConn,
// UDPConn, TCPConn. В зависимости от используемого протокола возвращается соответствующий тип.
//----------------------------------------------------------------------------
//
// Таким образом, используя данную функцию, мы можем отправлять запросы по протоколу TCP и UDP. Например:

func ex1() {
	httpRequest := "GET / HTTP/1.1\n" +
		"Host: golang.org\n\n"
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, conn)
	fmt.Println("Done")
}

/*
conn.Write([]byte(httpRequest)) посылает данные, которые здесь представлены переменной httpRequest.
Так как метод Write отправляет срез байтов, то любые данные надо преобразовать в срез байтов.
*/

//----------------------------------------------------------------------------
//
//
//----------------------------------------------------------------------------
//
//
//----------------------------------------------------------------------------
//
//
//----------------------------------------------------------------------------
//
//
//----------------------------------------------------------------------------
//
//
//----------------------------------------------------------------------------
//
//
//----------------------------------------------------------------------------
//
//
//----------------------------------------------------------------------------

func ExampleNet() {
	ex1()
}
