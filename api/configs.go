package api

import (
	"net/http"

	"github.com/celestiaorg/valuter/configs"
	"github.com/celestiaorg/valuter/tools"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*-------------*/
/*
* This function implements GET /configs
 */
func GetAllConfigs(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	tools.SendJSON(resp, configs.Configs)
}
