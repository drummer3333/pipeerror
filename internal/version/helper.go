package version

import (
	"flag"
	"os"
)

// generate version
//go:generate go run ../../gen/version/version.go -output ./version.go

// ParseVersionFlag print Version and exit if flag present.
func ParseVersionFlag() {
	versionFlag := flag.Bool("version", false, "print version number and exit")
	flag.Parse()

	if *versionFlag {
		println(Version)
		os.Exit(0)
	} else {
		println("current version", Version)
	}
}
