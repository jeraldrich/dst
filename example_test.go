package dst_test

import (
	"net"

	"github.com/calmh/dst"
)

func ExampleMux_Dial() {
	// Create an underlying UDP socket on a random local port.
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{})
	if err != nil {
		panic(err)
	}

	// Create a UDT mux around the packet connection with the default max
	// packet size.
	mux := dst.NewMux(udpConn, 0)

	// Dial a UDT connection. The address is that of a remote UDT mux.
	conn, err := mux.Dial("dst", "192.0.2.42:23458")
	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("Hello via UDT!"))
	if err != nil {
		panic(err)
	}
}

func ExampleMux_Accept() {
	// Create an underlying UDP socket on a specified local port.
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 23458})
	if err != nil {
		panic(err)
	}

	// Create a UDT mux around the packet connection with the default max
	// packet size.
	mux := dst.NewMux(udpConn, 0)

	// Accept new UDT connections and handle them in a separate routine.
	for {
		conn, err := mux.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(conn)
	}
}

func handleConn(net.Conn) {}
