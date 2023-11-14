package webserver

import (
	"fmt"
	"net/http"

	"github.com/emrekasg/personal-website-api/models"
	"github.com/gorilla/mux"
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
		WriteResp(res, http.StatusInternalServerError, "Error while getting posts", nil)
		return
	}

	WriteResp(res, http.StatusOK, "Posts fetched successfully", posts)
}

func GetPost(res http.ResponseWriter, req *http.Request) {
	language := req.URL.Query().Get("language")
	if ok := CheckLanguage(language); !ok {
		WriteResp(res, http.StatusBadRequest, "Language must be one of the following: "+getAllowedLanguages(), nil)
		return
	}

	postLink := mux.Vars(req)["postLink"]
	if postLink == "" {
		WriteResp(res, http.StatusBadRequest, "Post link must be provided", nil)
		return
	}

	post, err := models.GetPost(postLink, language)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			WriteResp(res, http.StatusNotFound, "Post not found", nil)
			return
		}
		fmt.Println(err)
		WriteResp(res, http.StatusInternalServerError, "Error while getting post", nil)
		return
	}

	WriteResp(res, http.StatusOK, "Post fetched successfully", post)
}
