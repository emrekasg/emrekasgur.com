package webserver

import (
	"database/sql"
	"net/http"

	"github.com/emrekasg/personal-website-api/models"
)

func GetTags(res http.ResponseWriter, req *http.Request) {
	tags, err := models.GetTags()
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			WriteResp(res, http.StatusNotFound, "No tags found", nil)
		default:
			WriteResp(res, http.StatusInternalServerError, "Failed to fetch tags", nil)
		}
		return
	}

	WriteResp(res, http.StatusOK, "Tags fetched successfully", tags)
}
