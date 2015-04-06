package Lib

import (
	"flag"
	"github.com/HostsSitter/Enum"
)

type flags struct {
	Command  *string
	SourceId *int
}

var Flags flags

func init() {
	Flags.Command = flag.String("c", Enum.FLAGS_COMMAND_UPDATE, "The Command: update; backup; recovery;")
	Flags.SourceId = flag.Int("sid", Enum.FLAGS_SOURCE_ENUM_360kb, "The SourceId: 1 360kb")

	flag.Parse()
}
