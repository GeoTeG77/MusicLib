package repository

import (
	"fmt"
	storage "musiclib/internal/infrastructure/storage"
	models "musiclib/internal/models"
)

type Repository struct {
	storage *storage.Storage
}

type SongsRepository interface {
	GetAllSongs(limit int, offset int, filter string) (*[]models.FullSong, error)
	GetSongText(song *models.Song) (string, error)
	PostSong(song *models.FullSong) error
	UpdateSong(song *models.FullSong) error
	DeleteSong(song *models.Song) error
}

func NewRepository(storage *storage.Storage) *Repository {
	return &Repository{storage: storage}
}

func (r *Repository) GetAllSongs(limit int, offset int, filter string) (*[]models.FullSong, error) {
	query := fmt.Sprintf(`
	SELECT
		s.group_name,
		s.song_name,
		sd.release_date,
		sd.text,
		sd.link
	FROM
		songs s
	JOIN
		song_details sd ON s.detail_id = sd.id
	ORDER BY
		%s  -- Подставляем столбец для сортировки
	LIMIT $1 OFFSET $2
`, filter)

	rows, err := r.storage.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []models.FullSong
	for rows.Next() {
		var song models.FullSong
		err := rows.Scan(&song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &songs, nil
}

func (r *Repository) GetSongText(song *models.Song) (string, error) {
	query := `
	SELECT
		sd.text
	FROM
		songs s
	JOIN
		song_details sd ON s.detail_id = sd.id
	WHERE
		s.group_name = $1 AND s.song_name = $2;
`

	var text string
	err := r.storage.DB.QueryRow(query, song.Group, song.Song).Scan(&text)
	if err != nil {
		return "", err
	}
	return text, nil
}

func (r *Repository) PostSong(song *models.FullSong) error {
	tx, err := r.storage.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	insertSongDetailsQuery := `
		INSERT INTO song_details (release_date, text, link)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var songDetailID int
	err = tx.QueryRow(insertSongDetailsQuery, song.ReleaseDate, song.Text, song.Link).Scan(&songDetailID)
	if err != nil {
		return err
	}

	insertSongQuery := `
		INSERT INTO songs (group_name, song_name, detail_id)
		VALUES ($1, $2, $3)
	`
	_, err = tx.Exec(insertSongQuery, song.Group, song.Song, songDetailID)
	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateSong(song *models.FullSong) error {
	tx, err := r.storage.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var detailID int
	getDetailIDQuery := `
		SELECT detail_id
		FROM songs
		WHERE group_name = $1 AND song_name = $2
	`
	err = tx.QueryRow(getDetailIDQuery, song.Group, song.Song).Scan(&detailID)
	if err != nil {
		return err
	}

	updateSongDetailsQuery := `
		UPDATE song_details
		SET
			release_date = COALESCE($1, release_date), 
			text = COALESCE($2, text), 
			link = COALESCE($3, link)
		WHERE id = $4
	`
	_, err = tx.Exec(updateSongDetailsQuery, song.ReleaseDate, song.Text, song.Link, detailID)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteSong(song *models.Song) error {
	tx, err := r.storage.DB.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	var detailID int
	getDetailIDQuery := `
		SELECT detail_id
		FROM songs
		WHERE group_name = $1 AND song_name = $2
	`
	err = tx.QueryRow(getDetailIDQuery, song.Group, song.Song).Scan(&detailID)
	if err != nil {
		return err
	}

	deleteSongDetailsQuery := `
		DELETE FROM song_details
		WHERE id = $1
	`
	_, err = tx.Exec(deleteSongDetailsQuery, detailID)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
