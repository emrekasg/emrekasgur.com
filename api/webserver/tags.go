package webserver

import (
	"fmt"
	"net/http"

	"github.com/emrekasg/personal-website-api/models"
)

func GetTags(res http.ResponseWriter, req *http.Request) {
	tags, err := models.GetTags()
	if err != nil {
		fmt.Println(err)
		if err.Error() == "sql: no rows in result set" {
			WriteResp(res, http.StatusNotFound, "No tags found", nil)
			return
		}
		WriteResp(res, http.StatusInternalServerError, "Error while getting tags", nil)
		return
	}

	WriteResp(res, http.StatusOK, "Tags fetched successfully", tags)
}
