package messages

import (
	"github.com/forbole/juno/v6/database"
	"github.com/forbole/juno/v6/types"
)

// HandleMsg represents a message handler that stores the given message inside the proper database table
func HandleMsg(
	index int, msg types.Message, tx *types.Transaction,
	parseAddresses MessageAddressesParser, db database.Database,
) error {

	// Get the involved addresses
	addresses, err := parseAddresses(tx)
	if err != nil {
		return err
	}

	return db.SaveMessage(int64(tx.Height), tx.TxHash, msg, addresses)
}
