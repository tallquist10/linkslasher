package main

import (
	_ "embed"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/tallquist10/linkslasher/internal/api"
	dbs "github.com/tallquist10/linkslasher/internal/database"
	"github.com/tallquist10/linkslasher/internal/services"
)

type LinkSlasher struct {
	LinksApi *api.LinksApi
}

//go:embed sql/create.sql
var createQuery string

//go:embed sql/readByHash.sql
var readQuery string

// var updateQuery string

//go:embed sql/delete.sql
var deleteQuery string

func initLinkSlasher(db *sql.DB, router *gin.Engine) *LinkSlasher {
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Fast Links",
		})
	})

	linksService, err := services.NewLinksService(
		db,
		createQuery,
		readQuery,
		// updateQuery,
		deleteQuery,
	)
	if err != nil {
		panic(err)
	}
	router.GET("/:hash", func(c *gin.Context) {
		link, err := linksService.GetLink(c.Param("hash"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.Redirect(http.StatusTemporaryRedirect, link.Original)
	})

	linksApi := api.NewLinksApi(router, linksService)

	return &LinkSlasher{LinksApi: linksApi}
}

func (ls *LinkSlasher) run() {
	ls.LinksApi.Run()
}

func main() {
	db, err := dbs.New()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	linkSlasher := initLinkSlasher(db, router)
	linkSlasher.run()
}
