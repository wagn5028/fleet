package service

import (
	"context"
	"encoding/json"
	"net/http"
)

func decodeCreatePackRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req createPackRequest
	if err := json.NewDecoder(r.Body).Decode(&req.payload); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeModifyPackRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	id, err := idFromRequest(r, "id")
	if err != nil {
		return nil, err
	}
	var req modifyPackRequest
	if err := json.NewDecoder(r.Body).Decode(&req.payload); err != nil {
		return nil, err
	}
	req.ID = id
	return req, nil
}

func decodeDeletePackRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	name, err := nameFromRequest(r, "name")
	if err != nil {
		return nil, err
	}
	var req deletePackRequest
	req.Name = name
	return req, nil
}

func decodeDeletePackByIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	id, err := idFromRequest(r, "id")
	if err != nil {
		return nil, err
	}
	var req deletePackByIDRequest
	req.ID = id
	return req, nil
}

func decodeGetPackRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	id, err := idFromRequest(r, "id")
	if err != nil {
		return nil, err
	}
	var req getPackRequest
	req.ID = id
	return req, nil
}

func decodeListPacksRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	opt, err := listOptionsFromRequest(r)
	if err != nil {
		return nil, err
	}
	return listPacksRequest{ListOptions: opt}, nil
}

func decodeApplyPackSpecsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req applyPackSpecsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil

}
