package api

import (
	"log"
	"net/http"

	"github.com/celestiaorg/valuter/tasks"
	"github.com/celestiaorg/valuter/tools"
	"github.com/celestiaorg/valuter/tx"
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
/*
* This function implements GET /staking/delegations/:address
 */
func GetDelegations(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	address := params.ByName("address")

	records, err := tx.GetDelegations(address)
	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(records) == 0 {
		http.Error(resp, "not found", http.StatusNotFound)
		return
	}

	tools.SendJSON(resp, records)
}

/*-------------*/
/*
* This function implements GET /staking/redelegations/:address
 */
func GetRedelegations(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	address := params.ByName("address")

	records, err := tx.GetRedelegations(address)
	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(records) == 0 {
		http.Error(resp, "not found", http.StatusNotFound)
		return
	}

	tools.SendJSON(resp, records)
}

/*-------------*/
/*
* This function implements GET /staking/undelegations/:address
 */
func GetUndelegations(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	address := params.ByName("address")

	records, err := tx.GetUndelegations(address)
	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(records) == 0 {
		http.Error(resp, "not found", http.StatusNotFound)
		return
	}

	tools.SendJSON(resp, records)
}

/*-------------*/
/*
* This function implements GET /staking/withdraw-rewards/:address
 */
func GetWithdrawDelegationRewards(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	address := params.ByName("address")

	records, err := tx.GetWithdrawDelegationRewards(address)
	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(records) == 0 {
		http.Error(resp, "not found", http.StatusNotFound)
		return
	}

	tools.SendJSON(resp, records)
}
