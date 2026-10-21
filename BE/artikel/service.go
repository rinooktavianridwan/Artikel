package artikel

import (
	"fmt"
)

var (
	PUBLISH = "publish"
	DRAFT   = "draft"
	TRASH   = "trash"
)

type ArtikelService struct {
	repo *ArtikelRepository
}

func NewArtikelService(repo *ArtikelRepository) *ArtikelService {
	return &ArtikelService{repo: repo}
}

func (s *ArtikelService) GetAllArtikel(limit int, offset int, status string) ([]Artikel, error) {
	return s.repo.GetAllArtikel(limit, offset, status)
}

func (s *ArtikelService) CreateArtikel(artikel *Artikel) error {
	//title minimal 20 karakter
	if len(artikel.Title) < 20 {
		return fmt.Errorf("title must be at least 20 characters long")
	}

	//content minimal 200 karakter
	if len(artikel.Content) < 200 {
		return fmt.Errorf("content must be at least 200 characters long")
	}

	//category minimal 3 karakter
	if len(artikel.Category) < 3 {
		return fmt.Errorf("category cannot be empty")
	}

	//status harus salah satu dari publish, draft, trash
	if artikel.Status != PUBLISH && artikel.Status != DRAFT && artikel.Status != TRASH {
		return fmt.Errorf("status must be one of: publish, draft, trash")
	}

	return s.repo.CreateArtikel(artikel)
}

func (s *ArtikelService) GetArtikelByID(id uint) (*Artikel, error) {
	return s.repo.GetArtikelByID(id)
}

func (s *ArtikelService) UpdateArtikel(artikel *Artikel) error {
	//artikel minimal 20 karakter
	if len(artikel.Title) < 20 {
		return fmt.Errorf("title must be at least 20 characters long")
	}

	//content minimal 200 karakter
	if len(artikel.Content) < 200 {
		return fmt.Errorf("content must be at least 200 characters long")
	}

	//category minimal 3 karakter
	if len(artikel.Category) < 3 {
		return fmt.Errorf("category must be at least 3 characters long")
	}

	//status harus salah satu dari publish, draft, trash
	if artikel.Status != PUBLISH && artikel.Status != DRAFT && artikel.Status != TRASH {
		return fmt.Errorf("status must be one of: publish, draft, trash")
	}
	return s.repo.UpdateArtikel(artikel)
}

func (s *ArtikelService) DeleteArtikel(id uint) error {
	return s.repo.DeleteArtikel(id)
}
