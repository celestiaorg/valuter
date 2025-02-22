package tools

import (
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/celestiaorg/valuter/configs"
	"github.com/celestiaorg/valuter/types"
)

/*------------------------------*/

func GetLimitOffsetFromHttpReq(req *http.Request) types.DBLimitOffset {
	qryParams := req.URL.Query()

	page := 1
	if _, ok := qryParams["page"]; ok {

		var err error
		page, err = strconv.Atoi(qryParams["page"][0])
		if err != nil {
			log.Printf("Error in page number: %v", err)
			page = 1
		}
		if page <= 0 {
			page = 1
		}
	}

	offset := (uint64(page) - 1) * configs.Configs.API.RowsPerPage

	return types.DBLimitOffset{
		Limit:  configs.Configs.API.RowsPerPage,
		Offset: offset,
		Page:   uint64(page),
	}
}

/*------------------------------*/

func GetPagination(totalRows, pageNumber uint64) types.Pagination {
	totalPages := uint64(math.Ceil(float64(totalRows) / float64(configs.Configs.API.RowsPerPage)))
	return types.Pagination{
		CurrentPage: pageNumber,
		TotalPages:  totalPages,
		TotalRows:   totalRows,
	}
}

/*------------------------------*/
