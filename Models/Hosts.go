package Models

import (
	"github.com/HostsSitter/Enum"
	"runtime"
)

type hosts struct {
	OS          string
	HostPath    string
	HostPreStr  string
	HostPostStr string
}

var Hosts hosts

func init() {
	Hosts.OS = runtime.GOOS
	if Hosts.OS == "windows" {
		Hosts.HostPath = Enum.HOST_PATH_LINUX
	} else {
		Hosts.HostPath = Enum.HOST_PATH_WINDOWS
	}

	Hosts.HostPreStr = Enum.HOST_PRE_STR
	Hosts.HostPostStr = Enum.HOST_POST_STR
}
