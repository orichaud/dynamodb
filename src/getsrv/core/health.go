package core

import (
	"fmt"
	"net/http"
)

func HandleHealthz(context *Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "InvocationCount=%d\n", context.GetInvocationCount())
	fmt.Fprintf(w, "SuccessCount=%d\n", context.GetSuccessCount())
	fmt.Fprintf(w, "FailedCount=%d\n", context.GetFailedCount())
}
