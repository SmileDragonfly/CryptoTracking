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
		values = append(values, row.Prevtime, row.Symbol, row.Price, row.Prevprice, row.Percent)
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", 5*i+1, 5*i+2, 5*i+3, 5*i+4, 5*i+5))
	}
	query := fmt.Sprintf("INSERT INTO tblBUSD1MinPercent (prevtime, symbol, price, prevPrice, percent) VALUES %s", strings.Join(placeholders, ", "))
	// Execute the query
	_, err := q.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) Insert5MinBUSDPercentRows(ctx context.Context, rows []Insert5MinBUSDPercentParams) error {
	// Build the query
	values := []interface{}{}
	placeholders := []string{}
	for i, row := range rows {
		values = append(values, row.Prevtime, row.Symbol, row.Price, row.Prevprice, row.Percent)
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", 5*i+1, 5*i+2, 5*i+3, 5*i+4, 5*i+5))
	}
	query := fmt.Sprintf("INSERT INTO tblBUSD5MinPercent (prevtime, symbol, price, prevPrice, percent) VALUES %s", strings.Join(placeholders, ", "))
	// Execute the query
	_, err := q.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) Insert10MinBUSDPercentRows(ctx context.Context, rows []Insert10MinBUSDPercentParams) error {
	// Build the query
	values := []interface{}{}
	placeholders := []string{}
	for i, row := range rows {
		values = append(values, row.Prevtime, row.Symbol, row.Price, row.Prevprice, row.Percent)
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", 5*i+1, 5*i+2, 5*i+3, 5*i+4, 5*i+5))
	}
	query := fmt.Sprintf("INSERT INTO tblBUSD10MinPercent (prevtime, symbol, price, prevPrice, percent) VALUES %s", strings.Join(placeholders, ", "))
	// Execute the query
	_, err := q.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) Insert15MinBUSDPercentRows(ctx context.Context, rows []Insert15MinBUSDPercentParams) error {
	// Build the query
	values := []interface{}{}
	placeholders := []string{}
	for i, row := range rows {
		values = append(values, row.Prevtime, row.Symbol, row.Price, row.Prevprice, row.Percent)
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", 5*i+1, 5*i+2, 5*i+3, 5*i+4, 5*i+5))
	}
	query := fmt.Sprintf("INSERT INTO tblBUSD15MinPercent (prevtime, symbol, price, prevPrice, percent) VALUES %s", strings.Join(placeholders, ", "))
	// Execute the query
	_, err := q.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) Insert30MinBUSDPercentRows(ctx context.Context, rows []Insert30MinBUSDPercentParams) error {
	// Build the query
	values := []interface{}{}
	placeholders := []string{}
	for i, row := range rows {
		values = append(values, row.Prevtime, row.Symbol, row.Price, row.Prevprice, row.Percent)
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", 5*i+1, 5*i+2, 5*i+3, 5*i+4, 5*i+5))
	}
	query := fmt.Sprintf("INSERT INTO tblBUSD30MinPercent (prevtime, symbol, price, prevPrice, percent) VALUES %s", strings.Join(placeholders, ", "))
	// Execute the query
	_, err := q.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}
	return nil
}

func (q Queries) Insert60MinBUSDPercentRows(ctx context.Context, rows []Insert60MinBUSDPercentParams) error {
	// Build the query
	values := []interface{}{}
	placeholders := []string{}
	for i, row := range rows {
		values = append(values, row.Prevtime, row.Symbol, row.Price, row.Prevprice, row.Percent)
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", 5*i+1, 5*i+2, 5*i+3, 5*i+4, 5*i+5))
	}
	query := fmt.Sprintf("INSERT INTO tblBUSD60MinPercent (prevtime, symbol, price, prevPrice, percent) VALUES %s", strings.Join(placeholders, ", "))
	// Execute the query
	_, err := q.db.ExecContext(ctx, query, values...)
	if err != nil {
		return err
	}
	return nil
}
