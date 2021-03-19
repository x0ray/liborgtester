// Package main - command liborgtester
package main

/* Expected output:
david@rat:~/go/src/github.com/x0ray/liborgtester$ liborgtester
2021/03/19 16:03:56 Command liborgtester version 0.0.2 of 13Mar2021
2021/03/19 16:03:56 New liborg data: 0xc000072180
2021/03/19 16:03:56 Package Lib 0.0.9 Dump
2021/03/19 16:03:56   Name  : Billy  Tag for lib(0): lib.name
2021/03/19 16:03:56   Name A:   Tag for NestA(0): nesta.name
2021/03/19 16:03:56   Name B: Billy_C  Tag for NestB(0):
2021/03/19 16:03:56   Name C: Billy_C  Tag for NestC(0): nestc.name
2021/03/19 16:03:56 Package Lib 0.0.9 Dump
2021/03/19 16:03:56   Name  : Billy  Tag for lib(0): lib.name
2021/03/19 16:03:56   Name A: Fred_A  Tag for NestA(0): nesta.name
2021/03/19 16:03:56   Name B: Fred_B  Tag for NestB(0):
2021/03/19 16:03:56   Name C: Fred_B  Tag for NestC(0): nestc.name
2021/03/19 16:03:56 Command liborgtester ended
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
