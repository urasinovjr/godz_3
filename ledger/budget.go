package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type Budget struct {
	Category string  `json:"category"`
	Limit    float64 `json:"limit"`
	Period   string  `json:"period"`
}

var budgets = map[string]Budget{}

func SetBudget(b Budget) {
	budgets[b.Category] = b
}

func LoadBudgets(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("не удалось прочитать бюджеты: %w", err)
	}

	var list []Budget
	err = json.Unmarshal(data, &list)
	if err != nil {
		return fmt.Errorf("не удалось разобрать JSON с бюджетами: %w", err)
	}

	for _, b := range list {
		SetBudget(b)
	}
	return nil
}

func PrintBudgets() {
	fmt.Println("Бюджеты:")
	for _, b := range budgets {
		fmt.Printf("  %s: %.2f руб. на %s\n", b.Category, b.Limit, b.Period)
	}
}
