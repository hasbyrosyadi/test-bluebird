package repository

import (
	"bluebird/model"

	"github.com/jinzhu/gorm"
)

type HistoryOrder struct {
	db *gorm.DB
}

func NewHistoryOrder(d *gorm.DB) HistoryOrderRepository {
	return &HistoryOrder{db: d}
}

type HistoryOrderRepository interface {
	GetAllHistory(userId int) ([]model.HistoryOrder, error)
}

func (h *HistoryOrder) GetAllHistory(userId int) ([]model.HistoryOrder, error) {
	historyOrder := []model.HistoryOrder{}
	err := h.db.Where("user_id = ?", userId).Order("created_at DESC").Find(&historyOrder).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return historyOrder, nil
}
