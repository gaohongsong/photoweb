package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/gmaclinuxer/photoweb/views"
)

var (
	Version     string
	BuildCommit string
	BuildTime   string
	GoVersion   string
)

// printVersion print version info from makefile
func printVersion() {
	fmt.Printf(`Version:      %s
Go version:   %s
Git commit:   %s
Built:        %s
`, Version, GoVersion, BuildCommit, BuildTime)
}

func main() {

	version := flag.Bool("v", false, "print version info")
	flag.Parse()

	// print version info only
	if *version {
		printVersion()
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/view", views.ViewHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("listen failed: ", err.Error())
	}
}
