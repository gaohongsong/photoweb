package views

import (
	"net/http"
	"path"
	. "github.com/gmaclinuxer/photoweb/common"
)

const UploadDir = "./uploads"

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := path.Join(UploadDir, imageId)
	if exists := IsExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

var ViewHandler = SafeHandler(viewHandler)