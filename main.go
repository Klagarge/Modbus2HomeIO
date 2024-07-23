package main

import (
	"Modbus2HomeIO/homeio"
	"Modbus2HomeIO/homeiosim"
	"Modbus2HomeIO/registers"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/simonvetter/modbus"
)

func main() {
	// Create the Home I/O REST client instance.
	home, err := homeio.New("http://127.0.0.1:9797")
	if err != nil {
		panic(err)
	}

	// Create the energy simulation.
	sim := homeiosim.New(home)

	// Create the Modbus handler instance.
	handler := registers.NewHandler(home, sim)

	// Create the Modbus TCP server instance.
	plain, err := modbus.NewServer(&modbus.ServerConfiguration{
		URL:        "tcp://0.0.0.0:1502",
		Timeout:    60 * time.Second,
		MaxClients: 10,
	}, handler)
	if err != nil {
		panic(err)
	}

	err = plain.Start()
	if err != nil {
		panic(err)
	}

	serverCertPool, _ := modbus.LoadCertPool("CA-OT-Security.crt")

	// Create the Modbus TCP+TLS server instance.
	// cert, err := GenerateSelfSignedCertificate("Modbus2HomeIO")
	cert, err := tls.LoadX509KeyPair("HomeIoServerTLS.crt", "key.pem")

	if err != nil {
		panic(err)
	}

	err = os.WriteFile("cert.pem", cert.Certificate[0], 0644)
	if err != nil {
		log.Fatal(err)
	}

	tls, err := modbus.NewServer(&modbus.ServerConfiguration{
		URL:           "tcp+tls://0.0.0.0:5802",
		Timeout:       30 * time.Second,
		MaxClients:    10,
		TLSServerCert: &cert,
		TLSClientCAs:  serverCertPool,
		//TLSClientCAs: &x509.CertPool{},
	}, handler)
	if err != nil {
		panic(err)
	}

	err = tls.Start()
	if err != nil {
		panic(err)
	}

	// Poll the Home I/O REST API and run the simulation every 25ms.
	for {
		time.Sleep(25 * time.Millisecond)
		err := home.Poll()
		if err != nil {
			fmt.Printf("Error polling Home I/O: %v\n", err)
		}
		sim.Step()
	}
}
