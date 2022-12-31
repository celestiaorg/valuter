package api

import (
	"log"
	"net/http"

	"github.com/celestiaorg/valuter/tools"
	"github.com/celestiaorg/valuter/tx"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*-------------*/
/*
* This function implements GET /tx/:hash
 */
func GetTx(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	hash := params.ByName("hash")

	record, err := tx.GetTx(hash)
	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if record == nil {
		http.Error(resp, "not found", http.StatusNotFound)
		return
	}

	tools.SendJSON(resp, record)
}
