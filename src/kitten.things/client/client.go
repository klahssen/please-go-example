package main

import (
	"fmt"
	"log"
	"os"

	pb "proto_out/nested_proto/kitten2"

	"github.com/jessevdk/go-flags"

	"kitten.things/kittenlib"
)

var opts struct {
	Port  int    `short:"p" long:"port" default:"9232" description:"Port to communicate with server on"`
	Breed string `short:"b" long:"breed" default:"HALP" description:"Breed of kitten to request"`
}

func main() {
	if _, err := flags.ParseArgs(&opts, os.Args); err != nil {
		os.Exit(1)
	}
	breed, present := pb.Breed_value[opts.Breed]
	if !present {
		log.Fatalf("Unknown breedd: %s", breed)
	}
	kitten := kittenlib.GetKitten(opts.Port, pb.Breed(breed))
	fmt.Printf("Received a kitten:\n%s\n", kitten)
}
