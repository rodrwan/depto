package items

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/baggyapp/depto"
	aErr "github.com/baggyapp/depto/pkg/errors"
	"github.com/baggyapp/depto/pkg/responses"
)

func getItems(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.Response, error) {
	item, err := ctx.ItemsService.Select(r.Context())
	if err != nil {
		if err == context.Canceled || err == context.DeadlineExceeded {
			return nil, aErr.NewInternalServerError(err)
		}

		return nil, aErr.NewNotFound(err)
	}

	return &responses.Response{
		Data:   item,
		Status: http.StatusOK,
	}, nil
}

func getItem(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.Response, error) {
	id, ok := r.Context().Value("id").(string)
	if !ok {
		err := errors.New("Invalid params")
		return nil, aErr.NewBadRequest(err, "Invalid params")
	}

	item, err := ctx.ItemsService.GetByID(r.Context(), id)
	if err != nil {
		return nil, aErr.NewInternalServerError(err)
	}

	info, err := ctx.ItemInfoService.GetByItemID(r.Context(), id)
	if err != nil {
		return nil, aErr.NewInternalServerError(err)
	}

	if info != nil {
		item.ID = ""

		item.Info = append(item.Info, info)
		for _, inf := range item.Info {
			images, err := ctx.ItemImagesService.GetByItemInfoID(r.Context(), inf.ID)
			if err != nil {
				return nil, aErr.NewInternalServerError(err)
			}

			inf.Images = images
		}
	}

	return &responses.Response{
		Data:   item,
		Status: http.StatusOK,
	}, nil
}

func createItems(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.Response, error) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var item depto.Item
	if err := dec.Decode(&item); err != nil {
		return nil, aErr.NewInternalServerError(err)
	}

	if err := ctx.ItemsService.Create(r.Context(), &item); err != nil {
		return nil, aErr.NewBadRequest(err, "Could not create item")
	}

	return &responses.Response{
		Data:   item,
		Status: http.StatusCreated,
	}, nil
}

func updateItems(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.Response, error) {
	id, ok := r.Context().Value("id").(string)
	if !ok {
		err := errors.New("Invalid params")
		return nil, aErr.NewBadRequest(err, "Invalid params")
	}

	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var item depto.Item
	if err := dec.Decode(&item); err != nil {
		return nil, aErr.NewInternalServerError(err)
	}

	i, err := ctx.ItemsService.Update(r.Context(), id, &item)
	if err != nil {
		return nil, aErr.NewBadRequest(err, "Could not update item")
	}

	return &responses.Response{
		Data:   i,
		Status: http.StatusOK,
	}, nil
}

func deleteItems(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.Response, error) {
	id, ok := r.Context().Value("id").(string)
	if !ok {
		err := errors.New("Invalid params")
		return nil, aErr.NewBadRequest(err, "Invalid params")
	}

	item, err := ctx.ItemsService.GetByID(r.Context(), id)
	if err != nil {
		return nil, aErr.NewInternalServerError(err)
	}

	if err := ctx.ItemsService.Delete(r.Context(), item); err != nil {
		return nil, aErr.NewBadRequest(err, "Could not delete item")
	}

	return &responses.Response{
		Data:   item,
		Status: http.StatusOK,
	}, nil
}

func addInfoToItem(ctx Context, w http.ResponseWriter, r *http.Request) (*responses.Response, error) {
	id, ok := r.Context().Value("id").(string)
	if !ok {
		err := errors.New("Invalid params")
		return nil, aErr.NewBadRequest(err, "Invalid params")
	}

	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var item depto.ItemInfo

	if err := dec.Decode(&item); err != nil {
		return nil, aErr.NewInternalServerError(err)
	}

	item.ItemID = id
	if err := ctx.ItemInfoService.Create(r.Context(), &item); err != nil {
		return nil, aErr.NewBadRequest(err, "Could not update item")
	}

	return &responses.Response{
		Data:   item,
		Status: http.StatusCreated,
	}, nil
}
