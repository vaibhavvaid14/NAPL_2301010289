package inventory

import (
	"errors"
	"fmt"
)

// Item represents an inventory product
type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

// Inventory manages a slice of items
type Inventory struct {
	Items []Item
}

// NewInventory initializes an empty inventory
func NewInventory() *Inventory {
	return &Inventory{Items: make([]Item, 0)}
}

// AddItem adds an item with input validation and duplicate checking
func (inv *Inventory) AddItem(item Item) error {
	if item.ID <= 0 {
		return errors.New("invalid ID: must be a positive integer")
	}
	if item.Name == "" {
		return errors.New("item name cannot be empty")
	}
	if item.Quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if item.Price < 0 {
		return errors.New("price cannot be negative")
	}

	// Check for duplicate ID
	for _, existing := range inv.Items {
		if existing.ID == item.ID {
			return fmt.Errorf("item with ID %d already exists", item.ID)
		}
	}

	inv.Items = append(inv.Items, item)
	return nil
}

// UpdateItem updates the quantity and price of an existing item by ID
func (inv *Inventory) UpdateItem(id int, newQuantity int, newPrice float64) error {
	if newQuantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	if newPrice < 0 {
		return errors.New("price cannot be negative")
	}

	for i, item := range inv.Items {
		if item.ID == id {
			inv.Items[i].Quantity = newQuantity
			inv.Items[i].Price = newPrice
			return nil
		}
	}
	return fmt.Errorf("item with ID %d not found", id)
}

// RemoveItem deletes an item from the inventory slice by ID
func (inv *Inventory) RemoveItem(id int) error {
	for i, item := range inv.Items {
		if item.ID == id {
			// Remove item from slice while preserving order
			inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("item with ID %d not found", id)
}

// Summary stores computed inventory metrics
type Summary struct {
	TotalCount int
	TotalValue float64
}

// ComputeSummaryAsync computes summary statistics concurrently using a goroutine and a channel
func (inv *Inventory) ComputeSummaryAsync() <-chan Summary {
	ch := make(chan Summary)
	go func() {
		defer close(ch)
		count := 0
		var value float64 = 0.0

		for _, item := range inv.Items {
			count += item.Quantity
			value += float64(item.Quantity) * item.Price
		}

		ch <- Summary{
			TotalCount: count,
			TotalValue: value,
		}
	}()
	return ch
}
