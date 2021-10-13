// +build ignore

package main

import (
	"flag"
	"log"
	"os"
	"text/template"

	"github.com/drummer3333/pipeerror/gen/genhelper"
)

type data struct {
	Output    string
	Version   string
	BuildTime string
}

func main() {
	var d data
	flag.StringVar(&d.Output, "output", "", "generated file name")
	flag.Parse()

	d.Version = getFromEnv("BUILD_VERSION", "SNAPSHOT")
	d.BuildTime = getFromEnv("CI_PIPELINE_CREATED_AT", "0")
	t := template.Must(template.New("version").Parse(tpl))
	rc, err := genhelper.TemplatePipe(t, d)
	if err != nil {
		log.Fatalf("generating from template: %v", err)
	}

	writtenBytes, err := genhelper.WriteToFile(rc, d.Output)
	if err != nil {
		log.Fatalf("error writing file: %v", err)
	}

	log.Printf("generated version infos: %s, %s. Written %d bytes", d.Version, d.BuildTime, writtenBytes)
}

func getFromEnv(variable string, defaultValue string) string {
	value := os.Getenv(variable)
	if value != "" {
		return value
	}

	return defaultValue
}

var tpl = `
/*
* CODE GENERATED AUTOMATICALLY WITH gitlab.zdf.de/gb-ist/vio/vio-uploadportal/gen/version
* THIS FILE SHOULD NOT BE EDITED BY HAND
*/

package version

var Version = "{{.Version}}"
var BuildTime = "{{.BuildTime}}"

`
