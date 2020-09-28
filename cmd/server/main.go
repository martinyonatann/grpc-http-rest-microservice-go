package main

import (
	"fmt"
	"os"

	"github.com/martinyonathann/grpc-http-rest-microservice-go/pkg/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
