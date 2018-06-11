package views

import (
	"net/http"
	"path"
	. "github.com/gmaclinuxer/photoweb/common"
	"io/ioutil"
	"fmt"
	"log"
	"io"
	"os"
)

const UploadDir = "./uploads"

func viewHandler(w http.ResponseWriter, r *http.Request) {

	imageId := r.FormValue("id")
	imagePath := path.Join(UploadDir, imageId)

	if exists := IsExists(imagePath); !exists {
		log.Printf("[%s]: %s, imagePath=%s 404", r.Method, r.URL, imagePath)
		http.NotFound(w, r)
		return
	}

	cType := "text/plain"

	if ext := path.Ext(imagePath); ext != "" {
		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
			cType = "image"
		} else {
			cType = "application/octet-stream"
		}
	}
	w.Header().Set("Content-Type", cType)

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
		src, h, err := r.FormFile("image")
		CheckError(err)

		filename := h.Filename
		defer src.Close()

		//dst, err := ioutil.TempFile(UploadDir, filename)
		//CheckError(err)
		//defer dst.Close()

		dst, err := os.OpenFile(path.Join(UploadDir, filename), os.O_WRONLY|os.O_CREATE, 0666)
		CheckError(err)
		defer dst.Close()

		cnt, err := io.Copy(dst, src)
		CheckError(err)

		log.Printf("save uploaded file: %s(%d)", filename, cnt)

		http.Redirect(w, r, "/views?id="+filename, http.StatusFound)
	}

	http.Error(w, "", http.StatusBadRequest)
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	urlQuery := r.URL.Query()
	fmt.Fprintf(w, "%s\n", urlQuery)

	// 获取GET请求参数
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	// Need parse first
	fmt.Printf("name: %s, age=%s\n", r.Form["name"], r.Form.Get("age"))
	fmt.Printf("name=%s, age=%s\n", r.PostFormValue("name"), r.PostFormValue("age"))
	fmt.Fprintf(w, "name=%s, age=%s\n", r.FormValue("name"), r.FormValue("age"))

	// No need parse
	fmt.Printf("name=%s, age=%s\n", r.FormValue("name"), r.FormValue("age"))

}

var ViewHandler = SafeHandler(viewHandler)
var ListHandler = SafeHandler(listHandler)
var UploadHandler = SafeHandler(uploadHandler)
var TestHandler = SafeHandler(testHandler)
