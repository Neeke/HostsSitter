package Core

import (
	"fmt"
	"github.com/HostsSitter/Enum"
	"github.com/HostsSitter/Lib"
	"github.com/HostsSitter/Models"
)

func Run() bool {
	if *Lib.Flags.SourceId == Enum.FLAGS_SOURCE_ENUM_360kb {
		Models.Source = Models.Source360
	}

	if *Lib.Flags.Command == Enum.FLAGS_COMMAND_UPDATE {
		return update()
	}

	return true
}

func update() bool {
	Lib.ParseUrl()
	fmt.Println("update ok")
	return true
}

func backup() bool {
	fmt.Println("backup ok")
	return true
}

func recovery() bool {
	fmt.Println("recovery ok")
	return true
}
