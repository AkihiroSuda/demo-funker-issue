package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/bfirsh/funker-go"
)

type data struct {
	X    string    `json:"x"`
	Time time.Time `json:"time"`
}

func callee(durationString string) error {
	d, err := time.ParseDuration(durationString)
	if err != nil {
		return err
	}
	log.Printf("[CALLEE] starting..")
	return funker.Handle(func(funkArgs *data) data {
		log.Printf("[CALLEE] handling args %+v", funkArgs)
		log.Printf("[CALLEE] sleeping for %s", d)
		time.Sleep(d)
		ret := data{X: funkArgs.X, Time: time.Now()}
		log.Printf("[CALLEE] returning %+v", ret)
		return ret
	})
}

func caller(funk string) error {
	funkArgs := data{X: "hello", Time: time.Now()}
	log.Printf("[CALLER] Calling funker %s(%+v)", funk, funkArgs)
	ret, err := funker.Call(funk, funkArgs)
	log.Printf("[CALLER] Called funker %s(%+v)=(%+v, %v)", funk, funkArgs, ret, err)
	return err
}

func printUsage(w io.Writer) {
	fmt.Fprintf(w, "Usage: %s [caller|callee] ..\n", os.Args[0])
	fmt.Fprintf(w, "%s caller FUNK\n", os.Args[0])
	fmt.Fprintf(w, "%s callee DURATION\n", os.Args[0])
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func main() {
	if len(os.Args) != 3 {
		printUsage(os.Stderr)
		os.Exit(1)
		return
	}
	var err error
	switch s := os.Args[1]; s {
	case "callee":
		err = callee(os.Args[2])
	case "caller":
		err = caller(os.Args[2])
	default:
		err = fmt.Errorf("unexpected args[1]: %q", s)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		printUsage(os.Stderr)
		os.Exit(1)
	}
}
