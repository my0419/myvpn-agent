package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/my0419/myvpn-agent/handler"
	"github.com/my0419/myvpn-agent/installer"
	"github.com/my0419/myvpn-agent/system"
	"golang.org/x/crypto/acme/autocert"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	log.Println("Welcome to MyVPN, your reliable service for installing and configuring VPN connections.")

	// debug mode
	if debugLogFile := system.CreateDebugLogFile(); debugLogFile != nil {
		log.SetFlags(log.Lshortfile)
		log.SetOutput(debugLogFile)
	}

	if len(os.Getenv("ENCRYPT_KEY")) != 32 {
		log.Fatal("Invalid ENCRYPT_KEY")
	}

	if os.Getenv("VPN_CLIENT_CONFIG_FILE") == "" {
		os.Setenv("VPN_CLIENT_CONFIG_FILE", "/tmp/myvpn-client-config")
	}

	setup, err := installer.CreateInstaller(os.Getenv("VPN_TYPE"))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		setup.RunPreStage()
	}()

	stopHttp := make(chan bool)

	http.HandleFunc("/", handler.HandleState(setup, stopHttp, os.Getenv("ENCRYPT_KEY")))

	http.HandleFunc("/debug", func(writer http.ResponseWriter, request *http.Request) {
		if false == system.DebugEnabled() {
			writer.WriteHeader(404)
			return
		}
		b, err := ioutil.ReadFile(system.DEBUG_FILE)
		if err != nil {
			log.Println(err)
			return
		}
		writer.Write(b)
	})

	ip, err := system.PublicIpAddr()
	if err != nil {
		log.Println("Failed to get public ip addr")
		return
	}

	domain := system.DomainName(ip)

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain),
	}

	serveTLS := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	serverCert := &http.Server{
		Addr:    ":80",
		Handler: certManager.HTTPHandler(nil),
	}

	serverHttp := &http.Server{
		Addr:    ":8400",
		Handler: nil,
	}

	// serve :80 - ACME challenge
	go func() {
		if err = serverCert.ListenAndServe(); err != nil {
			if err.Error() == "http: Server closed" {
				fmt.Println("Shutdown :80 success")
			} else {
				fmt.Printf("Start cert serve error: %s\n", err)
			}
		}
	}()

	// serve :8400 - API
	go func() {
		if err = serverHttp.ListenAndServe(); err != nil {
			if err.Error() == "http: Server closed" {
				fmt.Println("Shutdown :8400 success")
			} else {
				fmt.Printf("Start cert serve error: %s\n", err)
			}
		}
	}()

	// serve :443 - API
	go func() {
		if err = serveTLS.ListenAndServeTLS("", ""); err != nil {
			if err.Error() == "http: Server closed" {
				fmt.Println("Shutdown :443 success")
			} else {
				fmt.Printf("Start TLS error: %s\n", err)
			}
		}
	}()

	// wait finish
	for {
		select {
		case <-stopHttp:

			if err = serverCert.Shutdown(context.Background()); err != nil {
				fmt.Printf("Failed shutdown :80 error: %s\n", err.Error())
			}

			if err = serveTLS.Shutdown(context.Background()); err != nil {
				fmt.Printf("Failed shutdown :443 error: %s\n", err.Error())
			}

			if err = serverHttp.Shutdown(context.Background()); err != nil {
				fmt.Printf("Failed shutdown :8400 error: %s\n", err.Error())
			}

			setup.RunPostStage()

			os.Exit(0) // turn off agent response are delivered
		}
	}

}
