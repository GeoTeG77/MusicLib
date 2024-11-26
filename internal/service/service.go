package service

import (
	models "musiclib/internal/models"
	"musiclib/internal/repository"
	"strings"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllSongs(limit int, offset int, filter string) (*[]models.FullSong, error) {
	return s.repo.GetAllSongs(limit, offset, filter)
}

func (s *Service) GetSongText(song *models.Song, verse int) (string, error) {
	text, err := s.repo.GetSongText(song)
	if err != nil {
		return "", err
	}
	verses := strings.Split(text, "\\n\\n")
	if len(verses) < verse {
		return "", err
	}
	return verses[verse], nil
}

func (s *Service) PostSong(song *models.FullSong) error {
	return s.repo.PostSong(song)
}

func (s *Service) UpdateSong(song *models.FullSong) error {
	return s.repo.UpdateSong(song)
}

func (s *Service) DeleteSong(song *models.Song) error {
	return s.repo.DeleteSong(song)
}
