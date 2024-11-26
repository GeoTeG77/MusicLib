package api

import (
	"encoding/json"
	"fmt"
	"musiclib/internal/models"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	_ "github.com/sirupsen/logrus"
	_ "github.com/swaggo/swag"
)

// GetAllSongs godoc
// @Summary Get all songs with optional filters
// @Description Retrieve a list of songs with optional limit, page, and filter parameters
// @tags song
// @Produce json
// @Param limit query string false "Limit the number of results"
// @Param page query string false "Specify the page number"
// @Param filter query string false "Filter songs by field (song, group, release)"
// @Success 200 {array} models.FullSong "List of songs"
// @Failure 400 "Invalid request parameters"
// @Failure 500 "Internal server error"
// @Router /songs [get]
func (r *Router) GetAllSongs(w http.ResponseWriter, req *http.Request) {
	strLimit := req.URL.Query().Get("limit")
	strPage := req.URL.Query().Get("page")
	filter := req.URL.Query().Get("filter")

	var limit int
	var page int
	var offset int
	var err error

	if strLimit != "" {
		limit, err = strconv.Atoi(strLimit)
		if err != nil {
			r.Log.Info("Invalid request parameters")
			http.Error(w, "Invalid request parameters", http.StatusBadRequest)
			return
		}
	} else {
		limit = -1
		offset = 0
	}

	if strPage != "" {
		page, err = strconv.Atoi(strPage)
		if err != nil {
			r.Log.Info("Invalid request parameters")
			http.Error(w, "Invalid request parameters", http.StatusBadRequest)
			return
		}
	} else {
		page = 0
	}

	if page == 0 && limit != -1 {
		offset = 0
	}

	if page != 0 && limit == -1 {
		r.Log.Info("Invalid request parameters")
		http.Error(w, "Invalid request parameters", http.StatusBadRequest)
		return
	} else if page != 0 && limit != -1 {
		offset = limit + (limit * (page - 1))
	}

	switch filter {
	case "song":
		filter = "song_name"
	case "group":
		filter = "group_name"
	case "release":
		filter = "release_date"
	case "":
		filter = "NULL"
	default:
		r.Log.Info("Invalid filter parameter")
		http.Error(w, "Invalid filter parameter", http.StatusBadRequest)
		return
	}

	var Response *[]models.FullSong
	Response, err = r.Service.GetAllSongs(limit, offset, filter)
	if err != nil {
		r.Log.Info("Internal error")
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response)
	w.WriteHeader(http.StatusOK)
}

// GetSongText godoc
// @Summary Get song text
// @Description Retrieve the text for a specific song by its group and song name
// @tags song
// @Accept json
// @Produce json
// @Param group query string false "Group name"
// @Param song query string false "Song name"
// @Param verse query int false "Verse number (default 1)"
// @Success 200 {string} string "Text of the song"
// @Failure 400 "Invalid request parameters"
// @Failure 500 "Internal error"
// @Router /song/ [get]
func (r *Router) GetSongText(w http.ResponseWriter, req *http.Request) {
	song := req.URL.Query().Get("song")
	group := req.URL.Query().Get("group")

	strVerse := req.URL.Query().Get("verse")

	if song == "" || group == "" {
		r.Log.Info("Invalid request parameters")
		http.Error(w, "Invalid request parameters", http.StatusBadRequest)
		return
	}

	var err error
	var verse int

	if strVerse != "" {
		verse, err = strconv.Atoi(strVerse)
		if err != nil {
			r.Log.Info("Invalid request parameters")
			http.Error(w, "Invalid request parameters", http.StatusBadRequest)
			return
		}
	} else {
		verse = 1
	}

	request := models.Song{
		Song:  song,
		Group: group,
	}
	var Response string

	Response, err = r.Service.GetSongText(&request, verse)

	if err != nil {
		r.Log.Info("Internal error")
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response)
	w.WriteHeader(http.StatusOK)
}

// PostSong godoc
// @Summary Post a new song
// @Description Submit a new song's details and store it in the database
// @tags song
// @Accept json
// @Produce json
// @Param group query string false "Group name"
// @Param song query string false "Song name"
// @Success 200 {object} models.FullSong "The posted song details"
// @Failure 400 "Invalid request parameters"
// @Failure 500 "Request failed"
// @Router /song [post]
func (r *Router) PostSong(w http.ResponseWriter, req *http.Request) {

	group := req.URL.Query().Get("group")
	song := req.URL.Query().Get("song")

	if group == "" || song == "" {
		r.Log.Info("Missing request parameters")
		http.Error(w, "GGGGGG", http.StatusAccepted)
		return
	}

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	redirectURL := os.Getenv("API_URL")

	params := url.Values{}
	params.Add("group", group)
	params.Add("song", song)

	urlWithParams := fmt.Sprintf("%s?%s", redirectURL, params.Encode())

	newRequest, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		r.Log.Debug("Error Request creation")
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(newRequest)
	if err != nil {
		r.Log.Info("Request failed")
		http.Error(w, "Request failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var Response models.FullSong

	err = json.NewDecoder(resp.Body).Decode(&Response)
	if err != nil {
		r.Log.Info("Bad JSON")
		http.Error(w, "Failed to decode response", http.StatusBadRequest)
		return
	}
	Response.Song = song
	Response.Group = group

	err = r.Service.PostSong(&Response)

	if err != nil {
		r.Log.Info("Bad JSON")
		http.Error(w, "Failed to decode response", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response)
	w.WriteHeader(http.StatusOK)
}

// UpdateSong godoc
// @Summary Update an existing song
// @Description Update the details of a song, identified by group and song name
// @tags song
// @Accept json
// @Produce json
// @Param group query string false "Group name"
// @Param song query string false "Song name"
// @Param songDetails body models.FullSong true "Updated song details"
// @Success 200 {object} models.FullSong "The updated song details"
// @Failure 400 "Invalid request parameters"
// @Failure 500 "Internal error"
// @Router /song/ [patch]
func (r *Router) UpdateSong(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		r.Log.Info("Bad Request: Wrong Method")
		http.Error(w, "Failed to decode response", http.StatusBadRequest)
		return
	}

	group := req.URL.Query().Get("group")
	song := req.URL.Query().Get("song")

	if group == "" || song == "" {
		r.Log.Info("Missing request parameters")
		http.Error(w, "Invalid request parameters", http.StatusBadRequest)
		return
	}

	var Song models.FullSong

	Song.Song = song
	Song.Group = group

	err := json.NewDecoder(req.Body).Decode(&Song)
	if err != nil {
		r.Log.Info("Bad JSON")
		http.Error(w, "Failed to decode response", http.StatusBadRequest)
		return
	}

	err = r.Service.UpdateSong(&Song)
	if err != nil {
		r.Log.Info("Bad JSON")
		http.Error(w, "Failed to decode response", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Song)
	w.WriteHeader(http.StatusOK)
}

// DeleteSong godoc
// @Summary Delete a song
// @Description Delete the song specified by group and song name
// @tags song
// @Accept json
// @Produce json
// @Param group query string true "Group name"
// @Param song query string true "Song name"
// @Success 200 {string} string "Song successfully deleted"
// @Failure 400 "Invalid request parameters"
// @Failure 500 "Internal error"
// @Router /song/ [delete]
func (r *Router) DeleteSong(w http.ResponseWriter, req *http.Request) {
	group := req.URL.Query().Get("group")
	song := req.URL.Query().Get("song")

	if group == "" || song == "" {
		r.Log.Info("Missing request parameters")
		http.Error(w, "Invalid request parameters", http.StatusBadRequest)
		return
	}

	var Response models.Song
	Response.Song = song
	Response.Group = group

	err := r.Service.DeleteSong(&Response)
	if err != nil {
		r.Log.Info("Missing request parameters")
		http.Error(w, "Invalid request parameters", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
