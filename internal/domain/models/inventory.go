package models

type Item struct {
	Type     string
	Quantity int
}

func InventoryFromMerchList(merchList []string) []Item {
	merchCounts := make(map[string]int)

	for _, merch := range merchList {
		if _, ok := merchCounts[merch]; !ok {
			merchCounts[merch] = 0
		}

		merchCounts[merch]++
	}

	var inventory []Item

	for merch, count := range merchCounts {
		inventory = append(inventory, Item{
			Type:     merch,
			Quantity: count,
		})
	}

	return inventory
}
