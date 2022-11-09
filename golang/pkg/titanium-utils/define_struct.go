package titaniumutils

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChainTransactionCount fisco链返回的TransactionCount对应的结构体
type ChainTransactionCount struct {

	// current block number
	BlockNumber string `json:"blockNumber,omitempty"`

	// total failed tx count
	FailedTxSum string `json:"failedTxSum,omitempty"`

	// total tx count
	TxSum string `json:"txSum,omitempty"`
}

// Validate validates this transaction count
func (m *ChainTransactionCount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this transaction count based on context it is used
func (m *ChainTransactionCount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChainTransactionCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChainTransactionCount) UnmarshalBinary(b []byte) error {
	var res ChainTransactionCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
