// Package main - command liborgtester
package main

/* Expected output:
david@rat:~/go/src/github.com/x0ray/liborgtester$ liborgtester
2021/03/14 20:54:40 Command liborgtester version 0.0.2 of 13Mar2021
2021/03/14 20:54:40 New liborg data: 0xc00018a040
2021/03/14 20:54:40 Package Lib 0.0.2 Dump
2021/03/14 20:54:40   Name: Billy
2021/03/14 20:54:40 Package Lib 0.0.2 Dump
2021/03/14 20:54:40   Name: Fred
2021/03/14 20:54:40 Command liborgtester ended
*/
import (
	"log"

	o "github.com/x0ray/liborg"
)

const (
	pgm  = "liborgtester"
	ver  = "0.0.2"
	date = "13Mar2021"
)

func main() {
	log.Printf("Command %s version %s of %s", pgm, ver, date)
	d := o.NewLiborg("Billy")
	log.Printf("New liborg data: %p", d)
	d.Dump()
	d.SetName("Fred")
	d.Dump()
	log.Printf("Command %s ended", pgm)
}
