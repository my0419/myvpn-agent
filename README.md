# MyVPN Agent

Utility to install and configure VPN protocols on the server

### Features

* Setup VPN protocols
* HTTP server gives the setup status and vpn configuration with AES encryption

### Layer Interaction

![Screenshot](diagram.png)

### Environment

* `VPN_TYPE` name of protocol to be installed `l2tp`, `pptp`, `openvpn` or `wireguard`
* `ENCRYPT_KEY` random AES key (32 characters)
* `VPN_CLIENT_CONFIG_FILE` a place to save the client connection config. `default: /tmp/myvpn-client-config`
* `DEBUG_AGENT` debug mode `(default: disabled)`