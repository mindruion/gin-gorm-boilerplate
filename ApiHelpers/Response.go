package ApiHelpers

import (
	"github.com/gin-gonic/gin"
	"gorm-gin/Config"
	"math"
	"strconv"
)

type ResponseDataWithPagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}

type ResponseData struct {
	Data interface{} `json:"data"`
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func GeneratePaginationFromRequest(c *gin.Context) Pagination {
	limit := 2
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
	return Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}
func RespondJSON(w *gin.Context, status int, payload interface{}) {
	var res ResponseData
	res.Data = payload
	w.JSON(200, res)
}

func RespondPaginationJSON(w *gin.Context, payload interface{}, pagination *Pagination) {
	var res ResponseDataWithPagination
	res.Data = payload
	Config.DB.Model(payload).Count(&res.TotalRows)
	res.TotalPages = int(math.Ceil(float64(res.TotalRows) / float64(pagination.Limit)))
	res.Page = pagination.Page
	res.Limit = pagination.Limit
	res.Sort = pagination.Sort
	w.JSON(200, res)
}
