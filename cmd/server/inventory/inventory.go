package inventory

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/baggyapp/depto"
	"github.com/baggyapp/depto/pkg/responses"
)

func getInventory(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.HTMLResponse, error) {
	id, ok := r.Context().Value("id").(string)
	if !ok {
		return nil, errors.New("invalid params")
	}

	item, err := ctx.InventoryService.GetByID(r.Context(), id)
	if err != nil {
		return nil, errors.New("not found")
	}

	return &responses.HTMLResponse{
		Data:   item,
		Status: http.StatusOK,
	}, nil
}

func createInventory(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.HTMLResponse, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	defer r.Body.Close()

	var inv depto.Inventory
	if err := json.Unmarshal(body, &inv); err != nil {
		return nil, errors.New("internal server error")
	}

	if err := ctx.InventoryService.Create(r.Context(), &inv); err != nil {
		return nil, errors.New("Could not create item")
	}

	return &responses.HTMLResponse{
		Data:   inv,
		Status: http.StatusCreated,
	}, nil
}

func deleteInventory(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.HTMLResponse, error) {
	id, ok := r.Context().Value("id").(string)
	if !ok {
		return nil, errors.New("invalid params")
	}

	inv, err := ctx.InventoryService.GetByID(r.Context(), id)
	if err != nil {
		return nil, errors.New("not found")
	}

	if err := ctx.InventoryService.Delete(r.Context(), inv); err != nil {
		return nil, errors.New("Could not create item")
	}

	return &responses.HTMLResponse{
		Data:   inv,
		Status: http.StatusOK,
	}, nil
}

func addItemToInventory(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.HTMLResponse, error) {
	id, ok := r.Context().Value("id").(string)
	if !ok {
		return nil, errors.New("invalid params")
	}

	_, err := ctx.InventoryService.GetByID(r.Context(), id)
	if err != nil {
		return nil, errors.New("not found")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	defer r.Body.Close()

	var invItm depto.InventoryItem
	if err := json.Unmarshal(body, &invItm); err != nil {
		return nil, errors.New("internal server error")
	}

	invItm.InventoryID = id

	if err := ctx.InventoryItemsService.Create(r.Context(), &invItm); err != nil {
		return nil, errors.New("Could not create item")
	}

	return &responses.HTMLResponse{
		Data:   invItm,
		Status: http.StatusCreated,
	}, nil
}

func updateItemInInventory(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.HTMLResponse, error) {
	id, ok := r.Context().Value("id").(string)
	if !ok {
		return nil, errors.New("invalid params")
	}

	_, err := ctx.InventoryService.GetByID(r.Context(), id)
	if err != nil {
		return nil, errors.New("not found")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	defer r.Body.Close()

	var invItm depto.InventoryItem
	if err := json.Unmarshal(body, &invItm); err != nil {
		return nil, errors.New("internal server error")
	}

	invItm.InventoryID = id

	if err := ctx.InventoryItemsService.Update(r.Context(), &invItm); err != nil {
		return nil, errors.New("Could not create item")
	}

	return &responses.HTMLResponse{
		Data:   invItm,
		Status: http.StatusOK,
	}, nil
}
