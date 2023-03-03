package main

import (
	"fmt"
	"log"
	"github.com/gabrielluizsf/get_domain_ip/cli"
	"os"
)

func main() {
	fmt.Println("Starting point");

	application := cli.Generate();
	err := application.Run(os.Args);
		cli.LogError(err);
}	