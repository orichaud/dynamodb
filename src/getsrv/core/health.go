package core

import (
	"net/http"

	"github.com/golang/glog"
)

func HandleHealthz(context *Context, w http.ResponseWriter, r *http.Request) {
	stats := context.GetStats()
	if err := Send(stats, w); err != nil {
		glog.Errorf("Cannot trasnfer stats as JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
