package repository

import (
	"fmt"
	"medicine/model"
	"strings"
)

type CourseRepository struct{}

func (repo *CourseRepository) ListCourse(userID int) ([]*model.CourseAndPlan, error) {
	//query := "SELECT id, medicine_name, medicine_image, medicine_type, medicine_timing, course_start_time, status, created_at, updated_at " +
	//	"FROM medicine_course " +
	//	"WHERE user_id = ? " +
	//	"ORDER BY id ASC"
	query := "SELECT " +
		"    mc.medicine_name, " +
		"    mc.medicine_image, " +
		"    mc.medicine_type, " +
		"    mc.medicine_timing, " +
		"    mc.course_start_time, " +
		"    mc.status, " +
		"    ( " +
		"        SELECT COUNT(*) " +
		"        FROM medicine_plan p " +
		"        WHERE p.medicine_id = mc.id" +
		"    ) AS frequency " +
		"FROM " +
		"    medicine_course mc " +
		"WHERE " +
		"    mc.user_id = ? "
	rows, err := MysqlClient.Query(query, userID)
	if err != nil {
		return nil, err
	}
	data := make([]*model.CourseAndPlan, 0)
	for rows.Next() {
		obj := &model.CourseAndPlan{}
		err := rows.Scan(
			&obj.MedicineName,
			&obj.MedicineImage,
			&obj.MedicineType,
			&obj.MedicineTiming,
			&obj.CourseStartTime,
			&obj.Status,
			&obj.Frequency)
		if err != nil {
			return nil, err
		}
		data = append(data, obj)
	}
	return data, nil
}

func (repo *CourseRepository) ListCourseV2(userID int) ([]*model.CourseAndPlan, error) {
	query := "SELECT " +
		"    mc.id AS course_id, " +
		"    mc.medicine_name, " +
		"    mc.medicine_type, " +
		"    mc.course_start_time, " +
		"    mc.medicine_timing, " +
		"    mc.status, " +
		"    GROUP_CONCAT(mp.plan_time ORDER BY mp.plan_time) AS plan_times, " +
		"    COUNT(mp.plan_time) AS frequency, " +
		"    ANY_VALUE(mp.type) AS type, " +
		"	 ANY_VALUE(mp.amount) AS amount " +
		"FROM " +
		"    medicine_course mc " +
		"LEFT JOIN " +
		"    medicine_plan mp ON mp.medicine_id = mc.id " +
		"WHERE " +
		"    mc.user_id = ? " +
		"GROUP BY " +
		"    mc.id " +
		"ORDER BY " +
		"    mc.id ASC"
	rows, err := MysqlClient.Query(query, userID)
	if err != nil {
		return nil, err
	}
	data := make([]*model.CourseAndPlan, 0)
	for rows.Next() {
		obj := &model.CourseAndPlan{}
		err := rows.Scan(
			&obj.CourseID,
			&obj.MedicineName,
			&obj.MedicineType,
			&obj.CourseStartTime,
			&obj.MedicineTiming,
			&obj.Status,
			&obj.PlanTimes,
			&obj.Frequency,
			&obj.Type,
			&obj.Amount,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, obj)
	}
	return data, nil
}

func (repo *CourseRepository) CreateCourse(course *model.CourseAndPlan) (int64, error) {
	query := "INSERT " +
		"INTO medicine_course (user_id, medicine_name, medicine_image, medicine_type, medicine_timing, course_start_time, status) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := MysqlClient.Exec(
		query, course.UserId,
		course.MedicineName,
		course.MedicineImage,
		course.MedicineType,
		course.MedicineTiming,
		course.CourseStartTime,
		course.Status,
	)
	if err != nil {
		return 0, err
	}
	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func (repo *CourseRepository) UpdateCourse(course *model.Course) (int64, error) {
	query := "UPDATE medicine_course " +
		"SET medicine_name = ?, medicine_image = ?, medicine_type= ?, medicine_timing = ?, course_start_time = ? " +
		"WHERE id = ?"
	result, err := MysqlClient.Exec(query, course.MedicineName, course.MedicineImage, course.MedicineType, course.MedicineTiming, course.CourseStartTime, course.ID)
	if err != nil {
		return 0, err
	}
	updatedID, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return updatedID, nil
}

func (repo *CourseRepository) UpdateCourseV2(course *model.CourseAndPlan) (int64, error) {
	dateTimes := strings.Split(course.CourseStartTime, " ")
	date := dateTimes[0]

	query1 := "UPDATE " +
		"medicine_course SET medicine_type = ?, medicine_timing = ?, course_start_time = ? " +
		"WHERE id = ?"

	query2 := "DELETE " +
		"FROM medicine_plan WHERE medicine_id = ?"

	query3 := "INSERT " +
		"INTO medicine_plan (medicine_id, amount, type, plan_time) " +
		"VALUES (?, ?, ?, ?)"

	query4 := "INSERT " +
		"INTO medicine_plan_record(user_id, plan_id, actual_time, medicine_name, memo, status) " +
		"VALUES (?, ?, ?, ?, ?, ?)"

	tx, err := MysqlClient.Begin()
	if err != nil {
		return 0, err
	}

	// 更新 course
	if _, err = tx.Exec(query1, course.MedicineType, course.MedicineTiming, course.CourseStartTime, course.CourseID); err != nil {
		tx.Rollback()
		return 0, err
	}

	// 删除旧 plan
	if _, err = tx.Exec(query2, course.CourseID); err != nil {
		tx.Rollback()
		return 0, err
	}

	// 插入新的 plan 和 plan_record
	for _, val := range course.CourseStartTimes {
		result, err := tx.Exec(query3, course.CourseID, course.Amount, course.Type, val)
		if err != nil {
			tx.Rollback()
			return 0, err
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			return 0, err
		}

		// 构造实际时间，如果需要加日期前缀
		actualTime := val
		if !strings.Contains(val, date) {
			actualTime = fmt.Sprintf("%s %s", date, val)
		}

		if _, err = tx.Exec(query4, course.UserId, insertId, actualTime, course.MedicineName, "", 0); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return 1, nil
}

//func (repo *CourseRepository) UpdateCourseV2(course *model.CourseAndPlan) (int64, error) {
//	dateTimes := strings.Split(course.CourseStartTime, " ")
//	date := dateTimes[0]
//	// 更新Course用例
//	query1 := "UPDATE medicine_course " +
//		"SET medicine_type = ?, medicine_timing = ?, course_start_time = ? " +
//		"WHERE id = ?"
//	// 删除plan记录
//	query2 := "DELETE " +
//		"FROM medicine_plan " +
//		"WHERE medicine_id = ? "
//	// 重新插入 plan 记录
//	query3 := "INSERT " +
//		"INTO medicine_plan (medicine_id, amount, type, plan_time) " +
//		"VALUES (?, ?, ?, ?) "
//	// 插入记录
//	query4 := "INSERT " +
//		"INTO medicine_plan_record(user_id, plan_id, actual_time, medicine_name, memo, status) " +
//		"VALUES (?, ?, ?, ?, ?, ?)"
//
//	begin, err := MysqlClient.Begin()
//	if err != nil {
//		return 0, err
//	}
//	// 更新用药方案
//	_, err = begin.Exec(query1, course.MedicineType, course.MedicineTiming, course.CourseStartTime, course.CourseID)
//	if err != nil {
//		err := begin.Rollback()
//		if err != nil {
//			return 0, err
//		}
//	}
//	// 删除用药计划
//	_, err = begin.Exec(query2, course.CourseID)
//	if err != nil {
//		err := begin.Rollback()
//		if err != nil {
//			return 0, err
//		}
//	}
//	// 从新添加用药计划
//	for _, val := range course.CourseStartTimes {
//		result, err := begin.Exec(query3, course.CourseID, course.Amount, course.Type, val)
//		if err != nil {
//				begin.Rollback()
//				return 0, err
//			}
//		}
//		insertId, err := result.LastInsertId()
//		if err != nil {
//			return 0, err
//		}
//		_, err = begin.Exec(query4, course.UserId, insertId, date+" "+val, course.MedicineName, "", 0)
//		if err != nil {
//			err = begin.Rollback()
//			if err != nil {
//				return 0, err
//			}
//		}
//	}
//	err = begin.Commit()
//	if err != nil {
//		return 0, err
//	}
//	return 1, err
//}

func (repo *CourseRepository) UpdateCourseStatusByID(course *model.Course) (int64, error) {
	query := "UPDATE medicine_course " +
		"SET status = ? WHERE id = ?"
	result, err := MysqlClient.Exec(query, course.Status, course.ID)
	if err != nil {
		return 0, err
	}
	updatedID, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return updatedID, err
}

func (repo *CourseRepository) RemoveCourse(course *model.Course) (int64, error) {
	query := "DELETE " +
		"FROM medicine_course " +
		"WHERE medicine_name = ? AND user_id = ?"
	result, err := MysqlClient.Exec(query, course.MedicineName, course.UserId)
	if err != nil {
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
}
