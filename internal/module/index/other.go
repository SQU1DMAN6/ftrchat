package index

import "net/http"

func Other(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the other directory in internal/module/index/controller.go/Other"))
}
