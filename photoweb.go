package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/gmaclinuxer/photoweb/views"
	"path"
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

const KeyDir = "./ssl"

func main() {

	version := flag.Bool("v", false, "print version info")
	flag.Parse()

	// print version info only
	if *version {
		printVersion()
		return
	}

	mux := http.NewServeMux()

	views.StaticDirHandler(mux, "/statics/", "./statics", 0)

	mux.HandleFunc("/", views.ListHandler)
	mux.HandleFunc("/view", views.ViewHandler)
	mux.HandleFunc("/upload", views.UploadHandler)

	//err := http.ListenAndServe(":8080", mux)
	err := http.ListenAndServeTLS(":8080",
		path.Join(KeyDir, "bksaas.crt"),
		path.Join(KeyDir, "bksaas.key"),
		mux)
	if err != nil {
		log.Fatal("listen failed: ", err.Error())
	}
}
