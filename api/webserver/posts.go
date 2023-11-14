package webserver

import (
	"fmt"
	"net/http"

	"github.com/emrekasg/personal-website-api/models"
)

func GetPosts(res http.ResponseWriter, req *http.Request) {
	language := req.URL.Query().Get("language")
	if ok := CheckLanguage(language); !ok {
		WriteResp(res, http.StatusBadRequest, "Language must be one of the following: "+getAllowedLanguages(), nil)
		return
	}

	limit, offset, err := GetLimitAndOffset(req)
	if err != nil {
		WriteResp(res, http.StatusBadRequest, "Limit and offset must be a number", nil)
		return
	}

	posts, err := models.GetPosts(limit, offset, language)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			WriteResp(res, http.StatusNotFound, "No posts found", nil)
			return
		}
		fmt.Println(err)
		WriteResp(res, http.StatusInternalServerError, "Error while getting posts", nil)
		return
	}

	WriteResp(res, http.StatusOK, "Posts fetched successfully", posts)
}
