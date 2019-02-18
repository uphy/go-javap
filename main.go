package main // import "go-javap"

import (
	"fmt"
	"go-javap/command"
	"go-javap/parser"
	"os"
)

func main() {
	cli := command.New()
	if err := cli.Execute(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "failed to execute command: ", err, os.Args[1:])
	}
	os.Exit(0)
	f, _ := os.Open(os.Args[1])
	c, err := parser.ReadClass(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
}
