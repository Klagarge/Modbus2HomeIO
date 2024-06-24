package main

import (
	"Modbus2HomeIO/homeio"
	"Modbus2HomeIO/homeiosim"
	"Modbus2HomeIO/registers"
	"crypto/x509"
	"fmt"
	"github.com/simonvetter/modbus"
	"time"
)

func main() {
	// Create the Home I/O REST client instance.
	home, err := homeio.New("http://10.211.55.3:9797")
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

	// Create the Modbus TCP+TLS server instance.
	cert, err := GenerateSelfSignedCertificate("Modbus2HomeIO")
	if err != nil {
		panic(err)
	}

	tls, err := modbus.NewServer(&modbus.ServerConfiguration{
		URL:           "tcp+tls://0.0.0.0:5802",
		Timeout:       30 * time.Second,
		MaxClients:    10,
		TLSServerCert: cert,
		TLSClientCAs:  &x509.CertPool{},
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
