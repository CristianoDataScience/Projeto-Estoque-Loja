package services

import (
	"estoque/internal/models"
	"fmt"
	"strconv"
	"time"
)

type Estoque struct {
	items map[string]models.Item
	logs  []models.Log
}

// Informação de estoque
func NewEstoque() *Estoque {
	return &Estoque{
		items: make(map[string]models.Item),
		logs:  []models.Log{},
	}
}

// Adicionar Estoque
func (e *Estoque) AddItem(item models.Item) error {
	if item.Quantily <= 0 {
		return fmt.Errorf("erro ao adicionar item: [ID: %d] possui uma quantidade menor insuficiente (0 ou negativa)",
			item.ID)
	}
	existingItem, exists := e.items[strconv.Itoa(item.ID)]
	if exists {
		item.Quantily += existingItem.Quantily
	}
	e.items[strconv.Itoa(item.ID)] = item

	e.logs = append(e.logs, models.Log{
		TimeStamp: time.Now(),
		Action:    "Entrada de Item no Estoque",
		User:      "Administrador",
		ItemId:    item.ID,
		Quantily:  item.Quantily,
		Reason:    "Adicionando novos items no estoque",
	})
	return nil
}

// Metodo para listar Itens do estoque
func (e *Estoque) ListItems() []models.Item {
	var itemList []models.Item
	for _, item := range e.items {
		itemList = append(itemList, item)
	}
	return itemList
}

// Método para listar Logs
func (e *Estoque) ViewLogs() []models.Log {
	return e.logs
}

// Método para calcular o preço Total
func (e *Estoque) CalculateTotalCost() float64 {
	var totalCost float64
	for _, item := range e.items {
		totalCost += float64(item.Quantily) * item.Price
	}
	return totalCost
}

func FindBy[T any](data []T, comparator func(T) bool) ([]T, error) {
	var result []T
	for _, v := range data {
		if comparator(v) {
			result = append(result, v)
		}

	}
	if len(result) == 0 {
		return nil, fmt.Errorf("Nenhum item foi encontrado")
	}
	return result, nil
}
