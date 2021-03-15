// Package main - command liborgtester
package main

import (
	"log"

	o "github.com/x0ray/liborg"
)

const  (
	pgm = "liborgtester"
	ver = "0.0.2"
	date = "13Mar2021"
)

func main() {
	log.Printf("Command %s version %s of  %s",pgm,ver,date)
	d := o.NewLiborg("Billy")
	d.Dump()
	d.SetName("Fred")
	d.Dump() 
	log.Printf("Command %s ended",pgm)
}
