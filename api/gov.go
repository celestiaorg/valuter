package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/celestiaorg/valuter/tasks"
	"github.com/celestiaorg/valuter/tools"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*
* This function implements GET /challenges/gov
 */
func GetGovWinners(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	winnersList, err := tasks.GetGovWinners()

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, winnersList.GetItems())
}

/*-------------*/
/*
* This function implements GET /challenges/gov/:proposal_id
 */
func GetGovWinnersPerProposal(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	proposalId, err := strconv.Atoi(params.ByName("proposal_id"))
	if err != nil {
		proposalId = 0
	}
	winnersList, err := tasks.GetGovWinnersPerProposal(uint64(proposalId))

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, winnersList.GetItems())
}

/*-------------*/
