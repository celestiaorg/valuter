package api

import (
	"log"
	"net/http"

	"github.com/celestiaorg/valuter/tasks"
	"github.com/celestiaorg/valuter/tools"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*
* This function implements GET /challenges/staking
 */
func GetStakingWinners(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	winnersList, err := tasks.GetStakingWinners()

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, winnersList.GetItems())
}

/*-------------*/
