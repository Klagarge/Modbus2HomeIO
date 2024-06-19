package main

import (
	"Modbus2HomeIO/homeio"
	"crypto/x509"
	"github.com/simonvetter/modbus"
	"time"
)

func main() {

	// Create the Home I/O client instance.
	home, err := homeio.New("http://10.211.55.3:9797")
	if err != nil {
		panic(err)
	}

	// Create the Modbus handler instance.
	handler := NewModbusHandler(home)

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

	for {
		time.Sleep(1 * time.Second)
	}
}
