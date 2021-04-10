package main

import (
	"crypto/tls"
	"github.com/my0419/myvpn-agent/handler"
	"github.com/my0419/myvpn-agent/installer"
	"github.com/my0419/myvpn-agent/system"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
	"os"
)

func main() {

	finish := make(chan bool)
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
		setup.Start()
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HandleState(setup, os.Getenv("ENCRYPT_KEY")))

	go func() {
		log.Println(http.ListenAndServe(":8400", mux))
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
			Cache: 		autocert.DirCache("/tmp"),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(domain),
		}
		server := &http.Server{
			Addr: ":443",
			Handler: mux,
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}
		go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
		if err := server.ListenAndServeTLS("", ""); err != nil {
			log.Println("Failed to serve TLS", err)
		}
	}()

	<-finish
}