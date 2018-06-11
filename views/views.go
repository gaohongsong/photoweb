package views

import (
	"net/http"
	"path"
	. "github.com/gmaclinuxer/photoweb/common"
	"io/ioutil"
	"io"
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

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileArr, err := ioutil.ReadDir(UploadDir)
	CheckError(err)

	ctx := make(Context)
	imgs := []string{}
	for _, f := range fileArr {
		imgs = append(imgs, f.Name())
	}
	ctx["images"] = imgs

	RenderTemplate(w, "list", ctx)

}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		RenderTemplate(w, "upload", nil)
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		CheckError(err)

		filename := h.Filename
		defer f.Close()

		t, err := ioutil.TempFile(UploadDir, filename)
		CheckError(err)
		defer t.Close()

		_, err = io.Copy(t, f)
		CheckError(err)

		http.Redirect(w, r, "/views?id="+filename, http.StatusFound)
	}

	http.Error(w, "", http.StatusBadRequest)
}

var ViewHandler = SafeHandler(viewHandler)
var ListHandler = SafeHandler(listHandler)
var UploadHandler = SafeHandler(uploadHandler)
