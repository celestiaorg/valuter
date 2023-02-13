package api

import (
	"log"
	"net/http"

	"github.com/celestiaorg/valuter/blocks"
	"github.com/celestiaorg/valuter/tools"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*-------------*/
/*
* This function implements GET /blocks/missing
 */
func GetMissingBlocks(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	latestHeight, err := blocks.GetLatestBlockHeight()
	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	missingBlocks, err := blocks.FindMissingBlocks(0, latestHeight)
	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, missingBlocks)
}
