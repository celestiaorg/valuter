package api

import (
	"log"
	"net/http"

	routing "github.com/julienschmidt/httprouter"
)

/*-------------------------*/

func setupRouter() *routing.Router {

	var router = routing.New()

	router.GET("/", IndexPage)
	router.GET("/configs", GetAllConfigs)
	// router.GET("/ui/*file_path", UI)

	// router.GET("/winners", CheckAPIKey(GetWinners))
	router.GET("/winners", GetWinners)
	router.GET("/winners/:address", GetWinner)

	router.GET("/challenges", GetListOfChallenges)
	router.GET("/challenges/gov", GetGovWinners)
	router.GET("/challenges/gov/:proposal_id", GetGovWinnersPerProposal)
	router.GET("/challenges/staking", GetStakingWinners)
	router.GET("/challenges/node-upgrade", GetNodeUpgradeWinners)
	router.GET("/challenges/validators-genesis", GetGenesisValidatorsWinners)
	router.GET("/challenges/validators-joined", GetJoinedAfterGenesisValidatorsWinners)
	router.GET("/challenges/jail-unjail", GetUnjailedValidatorsWinners)
	router.GET("/challenges/uptime", GetPerformanceTestWinners)
	router.GET("/challenges/uptime/:burst_index", GetPerformanceTestWinnersPerLoadBurst)
	router.GET("/challenges/contracts/max-net-rewards", GetMaxNetworkRewardsWinners)
	router.GET("/challenges/contracts/subsidize-users-fees", GetContractsSubsidizeUsersFeesWinners)

	router.GET("/participants", GetParticipants)
	router.GET("/participants/:address", GetParticipant)

	router.GET("/tx/:hash", GetTx)

	router.GET("/blocks/missing", GetMissingBlocks)

	// router.GET("/pfbs/sender/:address", GetPFBs)

	router.GET("/gov/:proposal_id/vote/:address", GetGovVotePerProposal)

	router.GET("/staking/delegations/:address", GetDelegations)
	router.GET("/staking/redelegations/:address", GetRedelegations)
	router.GET("/staking/undelegations/:address", GetUndelegations)
	router.GET("/staking/withdraw-rewards/:address", GetWithdrawDelegationRewards)

	router.GET("/validators", GetValidators)
	router.GET("/validators/validator/:address/signed-block/:height", GetValidatorSignedBlock)
	router.GET("/validators/validator/:address", GetValidator)
	router.GET("/validators/genesis", GetGenesisValidators)
	router.GET("/validators/joined", GetJoinedAfterGenesisValidators)
	router.GET("/validators/unjailed", GetUnjailedValidators)

	return router
}

/*-------------------------*/

// ListenAndServeHTTP serves the APIs
func ListenAndServeHTTP(addr string) {

	router := setupRouter()
	if addr == "" {
		addr = ":8090"
	}

	log.Printf("[INFO ] Serving on %s", addr)

	log.Fatal(http.ListenAndServe(addr, router))
}

/*-------------------------*/
