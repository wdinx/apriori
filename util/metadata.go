package util

import (
	"apriori-backend/model/web"
	"strconv"
)

func GetMetadata(pageParam string) *web.Metadata {
	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	limit := 5

	return &web.Metadata{
		Page:  page,
		Limit: limit,
	}
}

func GetMetadataWIthLimit(pageParam string, limitParam string) *web.Metadata {
	page, err := strconv.Atoi(pageParam)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		limit = 5
	}

	return &web.Metadata{
		Page:  page,
		Limit: limit,
	}
}
