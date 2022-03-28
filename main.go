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

	// debug mode
	if debugLogFile := system.CreateDebugLogFile(); debugLogFile != nil {
		log.SetFlags(log.Lshortfile)
		log.SetOutput(debugLogFile)
	}

	finish := make(chan bool)
	if len(os.Getenv("ENCRYPT_KEY")) != 32 {
		log.Fatal("Invalid ENCRYPT_KEY")
	}

	if os.Getenv("VPN_CLIENT_CONFIG_FILE") == "" {
		os.Setenv("VPN_CLIENT_CONFIG_FILE", "/tmp/myvpn-client-config")
	}

	setup, err := installer.CreateInstaller(os.Getenv("VPN_TYPE"), os.Getenv("VPN_OS"))
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		setup.Start()
	}()

	stopAutocertHttp := make(chan bool)

	http.HandleFunc("/", handler.HandleState(setup, stopAutocertHttp, os.Getenv("ENCRYPT_KEY")))

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

	go func() {
		log.Println(http.ListenAndServe(":8400", nil))
	}()

	go func() {
		ip, err := system.PublicIpAddr()
		if err != nil {
			log.Println("Failed to get public ip addr")
			return
		}
		domain := system.DomainName(ip)
		log.Println("Host:", domain)

		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(domain),
		}
		server := &http.Server{
			Addr: ":https",
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}

		go func() {
			certServe := &http.Server{Addr: ":80", Handler: certManager.HTTPHandler(nil)}

			go func() {
				if err = certServe.ListenAndServe(); err != nil {
					if err.Error() == "http: Server closed" {
						fmt.Println("Shutdown :80 success")
					} else {
						fmt.Printf("Start cert serve error: %s\n", err)
					}
				}
			}()

			for {
				select {
				case <-stopAutocertHttp:
					if certServe == nil {
						continue
					}
					if err = certServe.Shutdown(context.Background()); err != nil {
						fmt.Printf("Failed shutdown :80 error: %s\n", err.Error())
					}
				}
			}
		}()

		if err = server.ListenAndServeTLS("", ""); err != nil {
			log.Println("Failed to serve TLS", err)
		}
	}()

	<-finish
}
