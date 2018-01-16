package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func main() {
	//set up the flag parsers
	outputPath := flag.String("o", "domain/", "output path")
	suffix := flag.String("f", ".ixo.go", "output suffix")

	flag.Parse()

	if len(flag.Args()) != 1 {
		panic("could not resolve the DSN Name")
	}

	//construct the args to be passed to the actual xo
	args := make([]string, 0, 5)
	args = append(args, fmt.Sprintf("-o=%s", *outputPath))
	args = append(args, fmt.Sprintf("-suffix=%s", *suffix))
	args = append(args, fmt.Sprintf("--template-path=%s", "./templates"))
	args = append(args, flag.Args()[0])

	//shell out to the xo binary
	cmd := exec.Command("xo", args...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//print the output
	fmt.Println(string(out))

}
