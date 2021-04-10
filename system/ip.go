package system

import (
	"fmt"
	"os/exec"
	"strings"
)

func DomainName(ip string) string  {
	return fmt.Sprintf("%s.nip.io", ip)
}

func PublicIpAddr() (string, error)  {
	out, err := exec.Command("bash", "-c", "hostname -I | awk '{print $1}'").Output()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(out), "\n"), nil
}