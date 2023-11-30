package webserver

import (
	"net/http"

	"github.com/emrekasg/personal-website-api/models"
)

func RandomSadVideo(res http.ResponseWriter, req *http.Request) {
	sadVideo, err := models.GetRandomSadVideo()
	if sadVideo == "" || err != nil {
		WriteResp(res, http.StatusInternalServerError, "Error while getting sad video", nil)
		return
	}

	WriteResp(res, http.StatusOK, "Sad video fetched successfully", sadVideo)
}
