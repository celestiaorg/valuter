package api

import (
	"log"
	"net/http"

	"github.com/celestiaorg/valuter/configs"
	"github.com/celestiaorg/valuter/tasks"
	"github.com/celestiaorg/valuter/tools"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*
* This function implements GET /winners
 */
func GetWinners(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	winnersList, err := tasks.GetAllWinners()

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, winnersList.GetItems())
}

/*-------------*/
/*
* This function implements GET /winners/:address
 */
func GetWinner(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	address := params.ByName("address")
	winnerResults, err := tasks.GetWinnerByAddress(address)

	if err != nil {
		log.Printf("Error in db query: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, winnerResults)
}

/*-------------*/
/*
* This function implements GET /challenges
 */
func GetListOfChallenges(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	tools.SendJSON(resp, configs.Configs.Tasks)
}
