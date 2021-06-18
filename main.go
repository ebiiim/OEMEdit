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

func main() {

	usage := fmt.Sprintf("Usage:\n\t%s get|set\nSubcommands:\n\tget: get OEM information from registry and print it to stdout in YAML\n\tset: read OEM information from stdin in YAML and write it to registry (admin rights required)", os.Args[0])

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "%s", usage)
		os.Exit(1)
	}

	switch strings.ToLower(os.Args[1]) {
	case "get":
		y, err := getPCinYAML()
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s", y)
	case "set":
		err := setPCFromYAML(os.Stdin)
		if err != nil {
			log.Println(err)
		}
	}
}
