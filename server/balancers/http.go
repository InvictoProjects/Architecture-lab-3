package balancers

import (
	"github.com/invictoprojects/architecture-lab-3/server/tools"
	"log"
	"net/http"
)

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListBalancers(store, rw)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListBalancers(store *Store, rw http.ResponseWriter) {
	res, err := store.ListBalancers()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
