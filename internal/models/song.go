package models

// @Description Song represents a song in the music library
type Song struct {
	Group    string `json:"group"` // The name of the group
	Song     string `json:"song"`  // The name of the song
	DetailID int    `json:"id"`    // The id in SongDetail for shema
}

// @Description SongDetail represents a song detail
type SongDetail struct {
	ID          int    `json:"id"`          // The id related with Song in shema
	ReleaseDate string `json:"releaseDate"` // The release date of song
	Text        string `json:"text"`        // The text of song
	Link        string `json:"link"`        // The link to song
}

// @Description FullSong represents a song in some endpoints
type FullSong struct {
	Group       string `json:"group"`       //  The name of the group
	Song        string `json:"song"`        //  The name of the song
	ReleaseDate string `json:"releaseDate"` //  The release date of song
	Text        string `json:"text"`        // The text of song
	Link        string `json:"link"`        // The link to song
}
