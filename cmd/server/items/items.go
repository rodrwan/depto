package items

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/baggyapp/depto"
	deptoErrors "github.com/baggyapp/depto/pkg/errors"
	"github.com/baggyapp/depto/pkg/responses"
)

func getItemInfo(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.HTMLResponse, error) {
	id, ok := r.Context().Value("id").(string)
	if !ok {
		return nil, deptoErrors.NewHTML(errors.New("invalid params"), "URL invalida")
	}

	item, err := ctx.ItemsService.GetByID(r.Context(), id)
	if err != nil {
		return nil, deptoErrors.NewHTML(err, "Error interno")
	}

	info, err := ctx.ItemInfoService.GetByItemID(r.Context(), id)
	if err != nil {
		return nil, deptoErrors.NewHTML(err, "Error interno")
	}

	item.Info = append(item.Info, info)

	return &responses.HTMLResponse{
		Template: "detail",
		Data:     item,
		Status:   http.StatusOK,
	}, nil
}

func getItems(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.HTMLResponse, error) {
	items, err := ctx.ItemsService.Select(r.Context())
	if err != nil {
		return nil, deptoErrors.NewHTML(err, "Error al obtener productos")
	}

	return &responses.HTMLResponse{
		Template: "index",
		Data:     items,
		Status:   http.StatusOK,
	}, nil
}

func formItems(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.HTMLResponse, error) {
	return &responses.HTMLResponse{
		Template: "add",
		Data:     nil,
		Status:   http.StatusOK,
	}, nil
}

func createItems(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.HTMLResponse, error) {
	r.ParseForm()

	name := r.PostFormValue("name")
	description := r.PostFormValue("description")

	item := depto.Item{
		Name:        name,
		Description: description,
	}

	if err := ctx.ItemsService.Create(r.Context(), &item); err != nil {
		return nil, deptoErrors.NewHTML(err, "Error interno")
	}

	return &responses.HTMLResponse{
		Data:     item,
		Template: "item-info",
		Status:   http.StatusMovedPermanently,
	}, nil
}

func addItemInfoIndex(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.HTMLResponse, error) {
	itemID, ok := r.Context().Value("id").(string)
	if !ok {
		return nil, deptoErrors.NewHTML(errors.New("invalid params"), "URL invalida")
	}

	item, err := ctx.ItemsService.GetByID(r.Context(), itemID)
	if err != nil {
		return nil, deptoErrors.NewHTML(err, "Error interno")
	}

	return &responses.HTMLResponse{
		Data:     item,
		Template: "item-info",
		Status:   http.StatusMovedPermanently,
	}, nil
}

func addItemInfo(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.HTMLResponse, error) {
	itemID, ok := r.Context().Value("id").(string)
	if !ok {
		return nil, deptoErrors.NewHTML(errors.New("invalid params"), "URL invalida")
	}
	r.ParseForm()

	brand := r.PostFormValue("brand")
	sPrice := r.PostFormValue("price")
	price, err := strconv.ParseFloat(sPrice, 64)
	if err != nil {
		return nil, deptoErrors.NewHTML(err, "Error interno")
	}

	// unit := r.PostFormValue("unit")
	sPurchaseDate := r.PostFormValue("purchase_date")
	purchaseDate, err := time.Parse("02/01/2006", sPurchaseDate)
	if err != nil {
		return nil, deptoErrors.NewHTML(err, "Error interno")
	}
	purchasePlace := r.PostFormValue("purchase_place")
	sExpirationDate := r.PostFormValue("expiration_date")
	expirationDate, err := time.Parse("02/01/2006", sExpirationDate)
	if err != nil {
		return nil, deptoErrors.NewHTML(err, "Error interno")
	}

	info := depto.ItemInfo{
		ItemID:         itemID,
		Brand:          brand,
		Price:          price,
		Unit:           0, //? dont know what to do with this
		PurchaseDate:   purchaseDate,
		PurchasePlace:  purchasePlace,
		ExpirationDate: expirationDate,
	}

	if err := ctx.ItemInfoService.Create(r.Context(), &info); err != nil {
		return nil, deptoErrors.NewHTML(err, "Error interno")
	}

	return &responses.HTMLResponse{
		Data:     info,
		Template: "item-info",
		Status:   http.StatusOK,
	}, nil
}
