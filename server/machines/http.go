package machines

import (
	"encoding/json"
	"github.com/invictoprojects/architecture-lab-3/server/tools"
	"log"
	"net/http"
)

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			handleMachineUpdate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleMachineUpdate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var m Machine
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		log.Printf("Error decoding machine input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.UpdateMachine(m.Id, m.IsWorking)
	if err == nil {
		tools.WriteJsonOk(rw, &m)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}
