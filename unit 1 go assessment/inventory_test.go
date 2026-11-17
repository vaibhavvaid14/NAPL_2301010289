package inventory

import (
	"testing"
)

func TestAddItem(t *testing.T) {
	inv := NewInventory()
	item := Item{ID: 1, Name: "Laptop", Quantity: 5, Price: 1200.0}

	// Test successful addition
	err := inv.AddItem(item)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(inv.Items) != 1 {
		t.Errorf("expected 1 item in inventory, got %d", len(inv.Items))
	}

	// Test duplicate ID error handling
	errDup := inv.AddItem(item)
	if errDup == nil {
		t.Error("expected error for duplicate item ID, got nil")
	}

	// Test invalid negative quantity validation
	errInvalid := inv.AddItem(Item{ID: 2, Name: "Mouse", Quantity: -1, Price: 25.0})
	if errInvalid == nil {
		t.Error("expected error for negative quantity, got nil")
	}
}

func TestUpdateItem(t *testing.T) {
	inv := NewInventory()
	inv.AddItem(Item{ID: 1, Name: "Laptop", Quantity: 5, Price: 1200.0})

	// Test successful update
	err := inv.UpdateItem(1, 10, 1150.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if inv.Items[0].Quantity != 10 || inv.Items[0].Price != 1150.0 {
		t.Errorf("item not updated correctly: %+v", inv.Items[0])
	}

	// Test updating non-existent item
	errMissing := inv.UpdateItem(99, 5, 100.0)
	if errMissing == nil {
		t.Error("expected error when updating non-existent item, got nil")
	}
}

func TestRemoveItem(t *testing.T) {
	inv := NewInventory()
	inv.AddItem(Item{ID: 1, Name: "Laptop", Quantity: 5, Price: 1200.0})

	// Test successful removal
	err := inv.RemoveItem(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(inv.Items) != 0 {
		t.Errorf("expected 0 items after removal, got %d", len(inv.Items))
	}

	// Test removing non-existent item
	errMissing := inv.RemoveItem(99)
	if errMissing == nil {
		t.Error("expected error when removing non-existent item, got nil")
	}
}

func TestComputeSummaryAsync(t *testing.T) {
	inv := NewInventory()
	inv.AddItem(Item{ID: 1, Name: "Mouse", Quantity: 10, Price: 25.0})   // Value: 250
	inv.AddItem(Item{ID: 2, Name: "Keyboard", Quantity: 4, Price: 50.0}) // Value: 200

	// Trigger concurrent summary calculation
	summaryCh := inv.ComputeSummaryAsync()
	summary := <-summaryCh

	expectedCount := 14
	expectedValue := 450.0

	if summary.TotalCount != expectedCount {
		t.Errorf("expected total count %d, got %d", expectedCount, summary.TotalCount)
	}
	if summary.TotalValue != expectedValue {
		t.Errorf("expected total value %.2f, got %.2f", expectedValue, summary.TotalValue)
	}
}
