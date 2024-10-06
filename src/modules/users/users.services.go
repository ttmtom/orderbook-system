package users

import (
	"fmt"
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			_, err := fmt.Fprintf(w, "%v: %v\n", name, h)
			if err != nil {
				return
			}
		}
	}
}
