package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type pc struct {
	OEMInformation OEMInfo `yaml:"OEMInformation"`
}

func getPCinYAML() ([]byte, error) {
	o, err := GetOEMInfo()
	if err != nil {
		return nil, err
	}
	p := pc{}
	p.OEMInformation = o
	y, err := yaml.Marshal(&p)
	if err != nil {
		return nil, err
	}
	return y, nil
}

func setPCFromYAML(r io.Reader) error {
	var p pc
	if err := yaml.NewDecoder(r).Decode(&p); err != nil {
		return err
	}
	if err := SetOEMInfo(p.OEMInformation); err != nil {
		return err
	}
	return nil
}

var (
	version = "dev"
)

func main() {
	usage := fmt.Sprintf("Usage:\n\t%s [command]\n\nCommands:\n\tget\tGets OEM information from registry and prints it to stdout in YAML\n\tset\tReads OEM information from stdin in YAML and writes it to registry (admin rights required)", os.Args[0])
	usage += fmt.Sprintf("\n\nOEMEdit Version %s by ebiiim\n", version)
	usage += "CLI tool for editing Windows OEM information\n"

	if len(os.Args) == 1 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(64)
	}

	switch strings.ToLower(os.Args[1]) {
	case "get":
		y, err := getPCinYAML()
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s", y)
	case "set":
		err := setPCFromYAML(os.Stdin)
		if err != nil {
			log.Print(err)
		}
	}
}
