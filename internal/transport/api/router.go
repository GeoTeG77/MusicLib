package api

import (
	"musiclib/internal/service"
	"net/http"

	"github.com/sirupsen/logrus"
	_ "github.com/swaggo/http-swagger/v2"
	_ "github.com/swaggo/swag"
)

type Router struct {
	Service *service.Service
	Log     *logrus.Logger
	Mux     *http.ServeMux
}

func NewRouter(log *logrus.Logger, service *service.Service) *Router {
	r := &Router{
		Service: service,
		Log:     log,
		Mux:     http.NewServeMux(),
	}

	//r.Mux.HandleFunc("/swagger/*", httpSwagger.WrapHandler)
	// GetAllSongs godoc
	// @Summary Get all songs with optional filters
	// @Description Retrieve a list of songs with optional limit, page, and filter parameters
	// @Accept query
	// @Produce json
	// @Param limit query int false "Limit the number of results"
	// @Param page query int false "Specify the page number"
	// @Param filter query string false "Filter songs by field (song, group, release)"
	// @Success 200 {array} models.FullSong "List of songs"
	// @Failure 400 {object} models.ErrorResponse "Invalid request parameters"
	// @Failure 500 {object} models.ErrorResponse "Internal server error"
	// @Router /GET/songs [get]
	r.Mux.HandleFunc("GET /api/v1/songs/", r.GetAllSongs)

	// @Router /song/ [get]
	r.Mux.HandleFunc("GET /api/v1/song/", r.GetSongText)

	// @Router /song/ [delete]
	r.Mux.HandleFunc("DELETE /api/v1/song/", r.DeleteSong)

	// @Router /song/ [patch]
	r.Mux.HandleFunc("PATCH /api/v1/song/", r.UpdateSong)

	// @Router /song [post]
	r.Mux.HandleFunc("POST /api/v1/song", r.PostSong)
	return r
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.Mux.ServeHTTP(w, req)
}
