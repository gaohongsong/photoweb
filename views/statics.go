package views

import (
	"net/http"
	"github.com/gmaclinuxer/photoweb/common"
	"log"
)

const ListDir = 0x0001

func StaticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {

	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {

		file := staticDir + r.URL.Path[len(prefix)-1:]

		if (flags & ListDir) == 0 {
			if exists := common.IsExists(file); !exists {
				log.Printf("[%s]: %s 404", r.Method, r.URL)
				http.NotFound(w, r)
				return
			}
		}

		log.Printf("[%s]: %s", r.Method, r.URL)
		http.ServeFile(w, r, file)
	})
}
