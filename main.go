package main

import (
	"Modbus2HomeIO/homeio"
	"Modbus2HomeIO/homeiosim"
	"Modbus2HomeIO/registers"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"time"

	modbus "github.com/Klagarge/modbusGo"
)

var CA = getCA(false)
var CERT = getServerCertificate(false)

func getCA(safeCA bool) *x509.CertPool {
	if safeCA {
		CA, err := modbus.LoadCertPool("CA-OT-Security.crt")
		if err != nil {
			fmt.Printf("failed to load server certificate authority (CA): %v\n", err)
			os.Exit(1)
		}
		return CA
	} else {
		return &x509.CertPool{}
	}
}

func getServerCertificate(safeCert bool) tls.Certificate {
	if safeCert {
		serverCert, err := tls.LoadX509KeyPair("HomeIoServerTLS.crt", "key.pem")
		if err != nil {
			fmt.Printf("failed to load server key pair: %v\n", err)
			os.Exit(1)
		}
		return serverCert
	} else {
		serverCert, err := GenerateSelfSignedCertificate("DebugCertificate")
		if err != nil {
			fmt.Printf("failed to generate server self certificate: %v\n", err)
			os.Exit(1)
		}
		return *serverCert
	}
}

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

	// Create the Modbus TCP+TLS server instance.
	tls, err := modbus.NewServer(&modbus.ServerConfiguration{
		URL:           "tcp+tls://0.0.0.0:5802",
		Timeout:       30 * time.Second,
		MaxClients:    10,
		TLSServerCert: &CERT,
		TLSClientCAs:  CA,
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
