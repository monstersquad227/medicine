package service

import (
	"medicine/model"
	"medicine/repository"
	"sort"
	"strings"
)

type RecordService struct {
	RecordRepository *repository.RecordRepository
}

func (svc *RecordService) Fetch(userID int) ([]*model.RecordModel, error) {
	return svc.RecordRepository.List(userID)
}

func (svc *RecordService) FetchV2(userID int) (interface{}, error) {
	type ContentDetail struct {
		ID           int     `json:"id"`
		MedicineName string  `json:"medicine_name"`
		Memo         *string `json:"memo"`
		Status       int     `json:"status"`
	}

	type GroupedContentItem struct {
		ActualDate string          `json:"actual_time"` // 只包含日期
		Contents   []ContentDetail `json:"contents"`
	}
	result, err := svc.RecordRepository.List(userID)
	if err != nil {
		return nil, err
	}

	groupMap := make(map[string][]ContentDetail)

	for _, value := range result {
		dateOnly := strings.Split(value.ActualTime, " ")[0]

		item := ContentDetail{
			ID:           value.ID,
			MedicineName: value.MedicineName,
			Memo:         value.Memo,
			Status:       value.Status,
		}

		groupMap[dateOnly] = append(groupMap[dateOnly], item)
	}

	var groupedResult []GroupedContentItem
	for date, contents := range groupMap {
		groupedResult = append(groupedResult, GroupedContentItem{
			ActualDate: date,
			Contents:   contents,
		})
	}

	// 可选：按时间排序（降序）
	sort.Slice(groupedResult, func(i, j int) bool {
		return groupedResult[i].ActualDate > groupedResult[j].ActualDate
	})

	return groupedResult, nil
}

func (svc *RecordService) Create(userID int, record *model.RecordModel) (int64, error) {
	record.UserID = userID
	return svc.RecordRepository.Create(record)
}
