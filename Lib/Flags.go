package Lib

import (
	"flag"
)

type flags struct {
	Ok   *bool
	Id   *int
	Port *string
}

var Flags flags

func init() {
	Flags.Ok = flag.Bool("ok", false, "is ok")
	Flags.Id = flag.Int("id", 0, "id")
	Flags.Port = flag.String("port", ":8080", "http listen port")

	flag.Parse()

}
