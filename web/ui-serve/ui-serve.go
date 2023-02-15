package ui_serve

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/tsotosa/atmm/config"
)

//go:embed build
var content embed.FS // build directory holds a client react build

func clientHandler() http.Handler {
	fsys := fs.FS(content)
	contentStatic, _ := fs.Sub(fsys, "build")
	return http.FileServer(http.FS(contentStatic))

}

func Up() {
	mux := http.NewServeMux()
	mux.Handle("/", clientHandler())
	http.ListenAndServe(fmt.Sprintf(":%d", config.Conf.UiPort), mux)

}
