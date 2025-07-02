package service

import (
	"math"
	"medicine/model"
	"medicine/repository"
	"sort"
	"strings"
	"time"
)

type RecordService struct {
	RecordRepository *repository.RecordRepository
	PlanRepository   *repository.PlanRepository
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
		IsChecked    int     `json:"is_checked"`
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
		//if value.ActualTime == nil || *value.ActualTime == "" {
		//	continue // 跳过该条记录，而不是整体中止
		//}
		dateOnly := strings.Split(value.ActualTime, " ")[0]

		item := ContentDetail{
			ID:           value.ID,
			MedicineName: value.MedicineName,
			Memo:         value.Memo,
			Status:       value.Status,
			IsChecked:    value.IsChecked,
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

	if len(groupedResult) == 0 {
		groupedResult = []GroupedContentItem{}
	}

	return groupedResult, nil
}

func (svc *RecordService) Create(userID int, record *model.RecordModel) (int64, error) {
	record.UserID = userID
	return svc.RecordRepository.Create(record)
}

func (svc *RecordService) Update(record *model.RecordModel) (int64, error) {
	/*
		数据库中第一次插入的actual_time 为添加 course 的时间
		计算数据库记录的记录的actual_time 与 实际传入的 actual_time 时间相差多少分钟，15内 status=0 正常打卡 反之异常
	*/
	today := time.Now().Format("2006-01-02")
	planTime, err := svc.PlanRepository.GetPlanTimeByIdAndUserID(record.PlanID)
	if err != nil {
		return 0, err
	}
	planTimeParse, err := time.Parse("2006-01-02 15:04", today+" "+planTime)
	if err != nil {
		return 0, err
	}
	actualTime := time.Now().Format("2006-01-02 15:04")
	actualTimeParse, err := time.Parse("2006-01-02 15:04", actualTime)
	if err != nil {
		return 0, err
	}

	//actualTime, err := svc.RecordRepository.GetActualTimeByPlanIDANDUserID(record.PlanID, record.UserID)
	//if err != nil {
	//	return 0, err
	//}
	//actualTimeParse, err := time.Parse("2006-01-02 15:04:05", actualTime)
	//if err != nil {
	//	return 0, err
	//}
	//recordActualTimeParse, err := time.Parse("2006-01-02 15:04:05", record.ActualTime)
	//if err != nil {
	//	return 0, err
	//}
	//diff := recordActualTimeParse.Sub(actualTimeParse)
	diff := actualTimeParse.Sub(planTimeParse)
	minutes := math.Abs(diff.Minutes())
	if minutes <= 15 {
		record.Status = 0
	} else {
		record.Status = 1
	}

	return svc.RecordRepository.Update(record)
}
