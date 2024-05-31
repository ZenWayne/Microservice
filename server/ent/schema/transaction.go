package schema

import (
	"database/sql"
	"database/sql/driver"
	"math/big"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/ethereum/go-ethereum/common"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	addressScannerFunc := field.ValueScannerFunc[common.Address, *sql.NullString]{
		V: func(addr common.Address) (driver.Value, error) {
			return addr.Hex(), nil
		},
		S: func(ns *sql.NullString) (common.Address, error) {
			if !ns.Valid {
				return common.Address{}, nil
			}
			addr := common.HexToAddress(ns.String)
			return addr, nil
		},
	}
	hashScannerFunc := field.ValueScannerFunc[common.Hash, *sql.NullString]{
		V: func(hash common.Hash) (driver.Value, error) {
			return hash.Hex(), nil
		},
		S: func(ns *sql.NullString) (common.Hash, error) {
			if !ns.Valid {
				return common.Hash{}, nil
			}
			hash := common.HexToHash(ns.String)
			return hash, nil
		},
	}
	return []ent.Field{
		field.Uint64("blockNumber"),
		field.String("txHash").NotEmpty().GoType(common.Hash{}).ValueScanner(hashScannerFunc),
		field.String("from").NotEmpty().GoType(common.Address{}).ValueScanner(addressScannerFunc),
		field.String("to").NotEmpty().GoType(common.Address{}).ValueScanner(addressScannerFunc),
		field.String("tokenId").GoType(&big.Int{}).ValueScanner(field.TextValueScanner[*big.Int]{}),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return nil
}
