package repository

import "medicine/model"

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
		err := rows.Scan(&obj.MedicineName,
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
