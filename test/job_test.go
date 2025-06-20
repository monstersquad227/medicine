package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
	"time"
)

type JobData struct {
	UserID       int    `json:"user_id"`
	PlanID       int    `json:"plan_id"`
	MedicineName string `json:"medicine_name"`
}

func TestJob(t *testing.T) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", "root", "1qaz@WSX", "47.103.98.61", "3306", "medicine", "utf8")
	MysqlClient, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("打开 Mysql 连接失败: %v", err)
	}
	if err = MysqlClient.Ping(); err != nil {
		log.Fatalf("连接 Mysql 失败: %v", err)
	}
	log.Println("MySQL 连接成功")

	now := time.Now().Format("2006-01-02")
	nowStartTime := now + " 00:00:00"
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	yesterdayStartTime := yesterday + " 00:00:00"
	yesterdayEndTime := yesterday + " 23:59:59"

	query := "SELECT " +
		"    user_id, " +
		"    plan_id, " +
		"    medicine_name " +
		"FROM " +
		"    medicine_plan_record " +
		"WHERE " +
		"    actual_time BETWEEN ? AND ?"

	yesterdayRecord := make([]*JobData, 0)
	rows, err := MysqlClient.Query(query, yesterdayStartTime, yesterdayEndTime)
	if err != nil {

	}
	for rows.Next() {
		obj := JobData{}
		if err := rows.Scan(&obj.UserID, &obj.PlanID, &obj.MedicineName); err != nil {
			log.Fatalf("Query Scan 失败: %v", err)
		}
		yesterdayRecord = append(yesterdayRecord, &obj)
	}

	insert := "INSERT " +
		"INTo medicine_plan_record(user_id, plan_id, medicine_name, actual_time) " +
		"VALUES (?, ?, ?, ?)"
	for _, record := range yesterdayRecord {
		result, err := MysqlClient.Exec(insert, record.UserID, record.PlanID, record.MedicineName, nowStartTime)
		if err != nil {
			log.Fatalf("Insert SQL 失败: %v", err)
		}
		lastId, err := result.LastInsertId()
		if err != nil {
			log.Fatalf("没有出入数据: %v", err)
		}
		log.Println(lastId)
	}
}
