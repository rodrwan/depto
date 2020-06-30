package inventory

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/baggyapp/depto/services/inventory"
	"github.com/baggyapp/depto/services/inventoryitem"
)

func TestGetInventory(t *testing.T) {
	req, err := http.NewRequest("GET", "/inventory/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = req.WithContext(context.WithValue(context.Background(), "id", "1"))
	rw := httptest.NewRecorder()

	ctx := Context{
		InventoryService: inventory.MockService{},
	}

	resp, err := getInventory(ctx, rw, req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != http.StatusOK {
		t.Fatalf("invalid status code, should be : %v receive: %v", http.StatusOK, resp.Status)
	}
}

func TestCreateInventory(t *testing.T) {
	bodyBytes := []byte(`
{
	"name": "Test inventory"
}
	`)
	bodyBuffer := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest("POST", "/inventory", bodyBuffer)
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()

	ctx := Context{
		InventoryService: inventory.MockService{},
	}

	resp, err := createInventory(ctx, rw, req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != http.StatusCreated {
		t.Fatalf("invalid status code, should be : %v receive: %v", http.StatusCreated, resp.Status)
	}
}

func TestAddItemToInventory(t *testing.T) {
	bodyBytes := []byte(`
{
	"item_id": "1",
	"count": 2
}
	`)
	bodyBuffer := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest("POST", "/inventory/1", bodyBuffer)
	if err != nil {
		t.Fatal(err)
	}

	req = req.WithContext(context.WithValue(context.Background(), "id", "1"))
	rw := httptest.NewRecorder()

	ctx := Context{
		InventoryService:      inventory.MockService{},
		InventoryItemsService: inventoryitem.MockService{},
	}

	resp, err := addItemToInventory(ctx, rw, req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != http.StatusCreated {
		t.Fatalf("invalid status code, should be : %v receive: %v", http.StatusCreated, resp.Status)
	}
}

func TestUpdateItemInInventory(t *testing.T) {
	bodyBytes := []byte(`
{
	"item_id": "1",
	"count": 1
}
	`)
	bodyBuffer := bytes.NewBuffer(bodyBytes)
	req, err := http.NewRequest("POST", "/inventory/1", bodyBuffer)
	if err != nil {
		t.Fatal(err)
	}

	req = req.WithContext(context.WithValue(context.Background(), "id", "1"))
	rw := httptest.NewRecorder()

	ctx := Context{
		InventoryService:      inventory.MockService{},
		InventoryItemsService: inventoryitem.MockService{},
	}

	resp, err := updateItemInInventory(ctx, rw, req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != http.StatusOK {
		t.Fatalf("invalid status code, should be : %v receive: %v", http.StatusCreated, resp.Status)
	}
}
