package Helpers

import (
	"Backend/Core/Constants/Keys/ErrorKeys"
	"Backend/Core/Constants/Keys/QueryKeys"
	"errors"
	"net/http"
	"strconv"
)

func StripIDFromUrl(w http.ResponseWriter,
	r *http.Request,
) (string, error) {
	id := r.URL.Query().Get(QueryKeys.ID)
	if id == "" {
		return "", errors.New(ErrorKeys.InvalidID)
	}
	return id, nil
}

func StripPaginationFromUrl(r *http.Request) (int64, int64, error) {
	page := r.URL.Query().Get(QueryKeys.Page)
	limit := r.URL.Query().Get(QueryKeys.Limit)
	pageNumber, limitNumber, err := convertPaginationToNumber(page, limit)
	return int64(pageNumber), int64(limitNumber), err
}

func convertPaginationToNumber(page string, limit string) (int64, int64, error) {
	pageInt, pageErr := strconv.ParseInt(page, 10, 64)
	limitInt, limitErr := strconv.ParseInt(limit, 10, 64)
	if pageErr != nil || limitErr != nil {
		return -1, -1, errors.New(ErrorKeys.InvalidPagination)
	}
	return pageInt, limitInt, nil
}
