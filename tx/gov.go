package tx

import (
	"github.com/celestiaorg/cosmologger/database"
	clgtx "github.com/celestiaorg/cosmologger/tx"
	"github.com/celestiaorg/valuter/types"
)

func GetGovVotePerProposal(address string, proposalId uint64) ([]types.TxRecord, error) {

	return getTxs(database.RowType{
		database.FIELD_TX_EVENTS_ACTION:      clgtx.ACTION_VOTE,
		database.FIELD_TX_EVENTS_PROPOSAL_ID: proposalId,
		database.FIELD_TX_EVENTS_SENDER:      address,
	})

}
