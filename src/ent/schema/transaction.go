package schema

import (
	"math/big"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

type uint256 struct{
	big.Int
	ent.Field.ValueScanner
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Text("from").NotEmpty().MinLen(42),
		field.Text("to").NotEmpty().MinLen(42),
		field.Other("tokenId", &big.Int{}).
			SchemaType(map[string]string{
				dialect.MySQL: "DECIMAL(78,0)",
			}),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return nil
}
