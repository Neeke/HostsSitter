package main

import (
	"flag"
	"fmt"
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

	var name string
	flag.StringVar(&name, "name", "123", "name")

	flag.Parse()

	fmt.Println("ok:", *Flags.Ok)
	fmt.Println("id:", *Flags.Id)
	fmt.Println("port:", *Flags.Port)
	fmt.Println("name:", name)
}
