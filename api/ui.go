package api

import (
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"strings"

	routing "github.com/julienschmidt/httprouter"
)

// IndexPage implements GET / Just a test msg to see if it works
func IndexPage(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	allAPIs := []string{
		"/configs",
		"/participants",
		"/participants/:address",
		"/validators",
		"/validators/validator/:address",
		"/validators/validator/:address/signed-block/:height",
		"/validators/genesis",
		"/validators/joined",
		"/validators/unjailed",
		"/winners",
		"/winners/:address",
		"/tx/:hash",
		"/blocks/missing",
		"/gov/:proposal_id/vote/:address",
		"/staking/delegations/:address",
		"/staking/redelegations/:address",
		"/staking/undelegations/:address",
		"/staking/withdraw-rewards/:address",
		"/challenges",
		"/challenges/gov",
		"/challenges/gov/:proposal_id",
		"/challenges/staking",
		"/challenges/node-upgrade",
		"/challenges/validators-genesis",
		"/challenges/validators-joined",
		"/challenges/jail-unjail",
		"/challenges/uptime",
		"/challenges/uptime/:burst_index",
		"/challenges/contracts/max-net-rewards",
		"/challenges/contracts/subsidize-users-fees",
	}

	modName := "unknown"
	buildInfo := ""
	if bi, ok := debug.ReadBuildInfo(); ok {
		modName = bi.Path

		buildInfo += "<br /><h3>Build Info:</h3><table>"
		for _, s := range bi.Settings {
			buildInfo += fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>", s.Key, s.Value)
		}
		buildInfo += "</table>"
	}

	html := `<!DOCTYPE html><html><head><style>
	table {border-collapse: collapse; width: 100%;}
	td, th {border: 1px solid #222;text-align: left; padding: 8px;}
	tr:nth-child(even) {background-color: #222;}
	a {
		text-decoration:none;border-bottom: 2px solid #10747f;
		color: #f1ff8f;transition: background 0.1s cubic-bezier(.33,.66,.66,1);
	}
	a:hover {background: #10747f;}
	body {
		color: #FFF; font-family: sans-serif;
		justify-content: center;align-items: center;
		line-height:1.8;margin:0;padding:0 40px;
		background-image: linear-gradient(135deg, rgba(0, 0, 0, 0.85) 0%,rgba(0, 0, 0,1) 100%);
	  }
	</style></head><body>`

	html += fmt.Sprintf("Ciao, this is `%v` \n\n<p>", modName)
	html += "<h3>List of endpoints:</h3>"
	for _, a := range allAPIs {

		href := strings.TrimPrefix(a, "/") // it fixes the links if the service is running under a path
		html += fmt.Sprintf(`<a href="%s">%s</a><br />`, href, a)
	}

	html += buildInfo

	resp.Header().Set("Content-Type", "text/html; charset=utf-8")
	resp.Write([]byte(html))
}

/*-------------------------*/

func UI(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	rootPath := os.Getenv("EXEC_PATH")
	if rootPath == "" {
		rootPath = "./"
	}

	http.FileServer(http.Dir(rootPath)).ServeHTTP(resp, req)
}

/*-------------------------*/
