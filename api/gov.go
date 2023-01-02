package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/celestiaorg/valuter/tasks"
	"github.com/celestiaorg/valuter/tools"
	"github.com/celestiaorg/valuter/tx"
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
/*
* This function implements GET /gov/:proposal_id/vote/:address
 */
func GetGovVotePerProposal(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	address := params.ByName("address")
	proposalId, err := strconv.ParseUint(params.ByName("proposal_id"), 10, 64)
	if err != nil {
		proposalId = 0
	}

	txs, err := tx.GetGovVotePerProposal(address, proposalId)
	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, txs)
}

/*-------------*/
