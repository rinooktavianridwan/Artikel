package artikel

import (
	"gorm.io/gorm"
)

type ArtikelRepository struct {
	db *gorm.DB
}

func NewArtikelRepository(db *gorm.DB) *ArtikelRepository {
	return &ArtikelRepository{db: db}
}

func (r *ArtikelRepository) GetAllArtikel(limit int, offset int, status string) ([]Artikel, error) {
	var artikels []Artikel
	result := r.db.Limit(limit).Offset(offset).Where("status = ?", status).Find(&artikels)
	return artikels, result.Error
}

func (r *ArtikelRepository) CreateArtikel(artikel *Artikel) error {
	result := r.db.Create(artikel)
	return result.Error
}

func (r *ArtikelRepository) GetArtikelByID(id uint) (*Artikel, error) {
	var artikel Artikel
	result := r.db.First(&artikel, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &artikel, nil
}

func (r *ArtikelRepository) UpdateArtikel(artikel *Artikel) error {
	result := r.db.Save(artikel)
	return result.Error
}

func (r *ArtikelRepository) DeleteArtikel(id uint) error {
	var artikel Artikel
	result := r.db.First(&artikel, id)
	if result.Error != nil {
		return result.Error
	}
	artikel.Status = TRASH
	result = r.db.Save(&artikel)
	return result.Error
}
