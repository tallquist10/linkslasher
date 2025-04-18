package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tallquist10/linkslasher/internal/links"
	"github.com/tallquist10/linkslasher/internal/services"
)

type LinksApi struct {
	router        *gin.Engine
	linksService  *services.LinksService
	requestLogger *services.RequestLogger
}

type LinksInput struct {
	Original string `json:"original"`
}

func NewLinksApi(router *gin.Engine, linksService *services.LinksService) *LinksApi {
	api := &LinksApi{
		router:        router,
		linksService:  linksService,
		requestLogger: services.NewRequestLogger(make(chan *links.LinksApiRequest)),
	}
	api.Configure()
	return api
}

func (api *LinksApi) GetLink(ctx *gin.Context) {
	api.requestLogger.LogRequest(&links.LinksApiRequest{
		Method: "GET",
		Params: map[string]string{"hash": ctx.Param("hash")},
	})
	link, err := api.linksService.GetLink(ctx.Param("hash"))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.IndentedJSON(http.StatusOK, link)
}

func (api *LinksApi) CreateLink(ctx *gin.Context) {
	var link *LinksInput
	err := ctx.ShouldBindBodyWithJSON(&link)
	if err != nil {
		api.requestLogger.LogRequest(&links.LinksApiRequest{
			Method: "POST",
			Params: map[string]string{"error": err.Error()},
		})
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	api.requestLogger.LogRequest(&links.LinksApiRequest{
		Method: "POST",
		Params: map[string]string{"original": string(link.Original)},
	})

	result, err := api.linksService.CreateLink(&links.Link{Original: link.Original})
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	api.requestLogger.LogRequest(&links.LinksApiRequest{
		Method: "GET",
		Params: map[string]string{"hash": ctx.Param("hash")},
	})
	ctx.IndentedJSON(http.StatusCreated, result)
}

func (api *LinksApi) Configure() {
	links := api.router.Group("/api/links") //, gin.BasicAuth(gin.Accounts{"Caleb": "password"}))
	// items.GET("/", api.GetAllItems)
	links.GET("/:hash", api.GetLink)
	links.POST("/", api.CreateLink)
	// items.PUT("/:id", api.EditWorkItem)
}

func (api *LinksApi) Run() {
	go api.requestLogger.Listen()
	api.router.Run()
}
