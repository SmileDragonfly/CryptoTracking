package sqlc

import (
	"context"
	"fmt"
	"strings"
)

func (q Queries) Insert1MinBUSDPercentRows(ctx context.Context, rows []Insert1MinBUSDPercentParams) error {
	// Build the query
	values := []interface{}{}
	placeholders := []string{}
	for i, row := range rows {
		values = append(values, row.Symbol, row.Price, row.Prevprice, row.Percent)
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d)", 4*i+1, 4*i+2, 4*i+3, 4*i+4))
	}
	query := fmt.Sprintf("INSERT INTO tblBUSD1MinPercent (symbol, price, prevPrice, percent) VALUES %s", "("+strings.Join(placeholders, ", ")+")")
	// Execute the query
	_, err := q.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}
	return nil
}
