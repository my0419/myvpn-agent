package system

import "os"

const DEBUG_FILE = "/tmp/myvpn-debug.log"

func DebugEnabled() bool  {
	return os.Getenv("MYVPN_DEBUG") == "true"
}

func CreateDebugLogFile() *os.File  {

	if false == DebugEnabled() {
		return nil
	}

	f, err := os.OpenFile(DEBUG_FILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		return nil
	}
	return f
}