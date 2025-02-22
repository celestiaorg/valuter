package api

import (
	"log"
	"net/http"

	"github.com/celestiaorg/valuter/participants"
	"github.com/celestiaorg/valuter/tools"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*
* This function implements GET /participants
 */
func GetParticipants(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	limitOffset := tools.GetLimitOffsetFromHttpReq(req)
	records, pagination, err := participants.GetParticipantsWithPagination(limitOffset)

	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp,
		map[string]interface{}{
			"pagination": pagination,
			"rows":       records,
		},
	)
}

/*-------------*/
/*
* This function implements GET /participants/:address
 */
func GetParticipant(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	address := params.ByName("address")

	record, err := participants.GetParticipantByAddress(address)
	if err != nil {
		log.Printf("API Call Error: %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if record.AccountAddress == "" {
		http.Error(resp, "Participant Not Found", http.StatusNotFound)
		return
	}

	tools.SendJSON(resp, record)
}

/*-------------*/
