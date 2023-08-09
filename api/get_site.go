package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/9bany/task/db/sqlc"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetSite(ctx *gin.Context) {
	url, ok := ctx.GetQuery("url")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("can not get url from param request")))
		return
	}

	if len(url) == 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("url not empty")))
		return
	}

	if !urlValid(url) {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("url must be URL format")))
		return
	}

	site, err := s.store.GetSiteByURL(ctx, url)
	if err != nil {
		if err == sql.ErrNoRows {
			site, err := s.renewSite(ctx, url)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			ctx.JSON(http.StatusOK, site)
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, site)

}

func (s *Server) renewSite(ctx context.Context, url string) (*db.Sites, error) {
	key, err := s.store.GetRandomKey(ctx)
	if err != nil {
		return nil, err
	}

	data, err := s.iframelyClient.FetchURL(ctx, key.Key, url)
	if err != nil || data == nil {
		return nil, fmt.Errorf("can not fetch url: %s", err.Error())
	}

	err = s.store.IncreaseKeyUsageCount(ctx, key.ID)
	if err != nil {
		// unnecessary lock user when have error here
		fmt.Println("warning: can not increase usage count of key with id: ", key.ID)
	}

	arg := db.CreateSiteParams{
		Url:      url,
		MetaData: data,
	}

	site, err := s.store.CreateSite(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("can not create new site: %s", err.Error())
	}
	return &site, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
