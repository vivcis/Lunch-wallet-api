package repository

import (
	"errors"
	"github.com/decadevs/lunch-api/internal/core/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GeneratePaginationFromRequest ..
func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	limit := 10
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}

func StatusEnum(text string) (string, error) {
	var result string
	notifyServe := [...]string{"serving", "SERVING", "Serving", "served", "SERVED", "Served"}
	for _, v := range notifyServe {
		if v == text {
			result = text
			return result, nil
		}
	}
	result = ""
	return result, errors.New("incorrect status field")
}
