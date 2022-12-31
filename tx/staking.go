package tx

import (
	"github.com/celestiaorg/cosmologger/database"
	clgtx "github.com/celestiaorg/cosmologger/tx"
	"github.com/celestiaorg/valuter/types"
)

func GetDelegations(delegatorAddr string) ([]types.TxRecord, error) {

	return getTxs(database.RowType{
		database.FIELD_TX_EVENTS_ACTION: clgtx.ACTION_DELEGATE,
		database.FIELD_TX_EVENTS_SENDER: delegatorAddr,
	})

}

func GetRedelegations(delegatorAddr string) ([]types.TxRecord, error) {

	return getTxs(database.RowType{
		database.FIELD_TX_EVENTS_ACTION: clgtx.ACTION_BEGIN_REDELEGATE,
		database.FIELD_TX_EVENTS_SENDER: delegatorAddr,
	})
}

func GetUndelegations(delegatorAddr string) ([]types.TxRecord, error) {

	return getTxs(database.RowType{
		database.FIELD_TX_EVENTS_ACTION: clgtx.ACTION_BEGIN_UNBONDING,
		database.FIELD_TX_EVENTS_SENDER: delegatorAddr,
	})
}

func GetWithdrawDelegationRewards(delegatorAddr string) ([]types.TxRecord, error) {

	return getTxs(database.RowType{
		database.FIELD_TX_EVENTS_ACTION: clgtx.ACTION_WITHDRAW_DELEGATOR_REWARD,
		database.FIELD_TX_EVENTS_SENDER: delegatorAddr,
	})
}
