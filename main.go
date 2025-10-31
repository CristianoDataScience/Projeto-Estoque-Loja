package main

import (
	"estoque/internal/models"
	"estoque/internal/services"
	"fmt"
)

func main() {
	fmt.Println("Gerenciamento de Estoque")

	items := []models.Item{
		{ID: 1, Name: "Camisa", Quantily: 3, Price: 50},
		{ID: 2, Name: "Calça", Quantily: 5, Price: 70},
		{ID: 3, Name: "Fone", Quantily: 10, Price: 100},
		{ID: 4, Name: "Jaqueta Tommy", Quantily: 40, Price: 220},
		{ID: 5, Name: "Blusa Couro", Quantily: 7, Price: 90},
	}
	estoque := services.NewEstoque()

	for _, item := range items {
		err := estoque.AddItem(item)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	fmt.Println(estoque.ListItems())

	for _, item := range estoque.ListItems() {
		fmt.Printf("\nID: %d | Item: %s | Quantidade: %d | Preço: %.2f",
			item.ID, item.Name, item.Quantily, item.Price)
	}

	fmt.Println()
	// fmt.Println(estoque.ViewLogs())
	// logs := estoque.ViewLogs()
	// for _, log := range logs {
	// 	fmt.Printf("\n[%s] Ação: %s - Usuário: %s - Item ID: %d - Quantidade: %d - Motivo: %s",
	// 		log.TimeStamp.Format("02/01 15:04:05"), log.Action, log.User, log.ItemId, log.Quantily, log.Reason)
	// }
	fmt.Println("\nValor Total R$: ", estoque.CalculateTotalCost())

	searchItem, err := services.FindBy(items, func(item models.Item) bool {
		return item.Price <= 30
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(searchItem)

}
