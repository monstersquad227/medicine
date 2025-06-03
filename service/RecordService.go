package service

import (
	"medicine/model"
	"medicine/repository"
)

type RecordService struct {
	RecordRepository *repository.RecordRepository
}

func (svc *RecordService) Fetch(userID int) ([]*model.RecordModel, error) {
	return svc.RecordRepository.List(userID)
}

func (svc *RecordService) Create(userID int, record *model.RecordModel) (int64, error) {
	record.UserID = userID
	return svc.RecordRepository.Create(record)
}
