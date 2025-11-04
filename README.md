# medicine


## MedicineCourse

|字段|描述|备注
|-|-|-|
|ID|自增ID||
|UserID|用户ID||
|MedicineName|药物名称||
|MedicineImage|药物图片||
|MedicineType|药物方式|0: 内服；1: 外用|
|MedicineTiming|用药时机|0: 不限；1:饭前用药；2: 饭后用药；3: 随餐用药；4: 睡前用药|
|CourseStartTime|用药开始时间||
|Status|方案状态|0: 生效；1: 废弃|

### 接口

#### 请求头
```
token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50IjoiMTUwKioqKjI4MjQiLCJjcmVhdGVfYXQiOiIyMDI1LTA1LTI3VDExOjAyOjUyLjgzMzM2OSswODowMCIsImV4cCI6MTc0ODM1ODE3Mn0.Kh0x3uhzN_s8ENBBWF68leYyehx58RUfQvFjPmdGO44
```

#### 获取方案

##### URL
```
GET /medicine/course
```

##### response
```json
{
	"code": 0,
	"message": "success",
	"result": [
		{
			"id": 1,
			"medicine_name": "厄贝沙坦氢氟噻嗪片",
			"medicine_timing": 1,
			"course_start_time": "11:35",
			"created_at": "2025-05-27 11:26:58",
			"updated_at": "2025-05-27 11:27:47"
		},
		{
			"id": 2,
			"medicine_name": "真菌王抑菌膏",
			"medicine_image": "",
			"medicine_timing": 4,
			"course_start_time": "09:00",
			"status": 1,
			"created_at": "2025-05-27 13:10:42",
			"updated_at": "2025-05-27 14:05:00"
		},
		{
			"id": 3,
			"medicine_name": "苯磺酸氨氯地平片",
			"medicine_image": "",
			"medicine_timing": 4,
			"course_start_time": "09:00",
			"created_at": "2025-05-27 14:49:09",
			"updated_at": "2025-05-27 14:51:56"
		}
	]
}
```

#### 创建方案

##### URL
```
POST /medicine/course
```

##### body
```json
{
  "user_id": 1,
  "medicine_name": "真菌王抑菌膏",
  "medicine_image": "",
  "medicine_type": 1, // 外用
  "medicine_timing": 0, // 不限
  "course_start_time": "2025-05-26", // 开始日期
  "status": 0,
  "amount": 1, // 每日次数，每日数量
  "type": "次", // 每次单位
  "plan_time": "09:00" // 时间
}
```

##### response
```json
{
	"code": 0,
	"message": "success",
	"result": 2
}
```

#### 更新方案

##### URL
```
PUT /medicine/course/:id
```

##### body
```json
{
	"medicine_name": "真菌王抑菌膏",
	"medicine_image": "",
	"medicine_type": 0,
	"medicine_timing": 4,
	"course_start_time": "09:00"
}
```

##### response
```json
{
	"code": 0,
	"message": "success",
	"result": 1
}
```

#### 方案禁用

##### URL
```
PATCH /medicine/course/:id
```

##### body
```json
{
	"status": 1
}
```

##### response
```json
{
	"code": 0,
	"message": "success",
	"result": 1
}
```

## MedicinePlan

|字段|描述|备注|
|-|-|-|
|ID|自增ID||
|MedicineID|用药方案ID||
|Amount|用药数量||
|Type|剂量单位||
|PlanTime|用药时间||

### Interface
+ AddMedicinePlan
+ PutMedicinePlan
+ UpdateMedicinePlan

## MedicinePlanRecord

|字段|描述|备注|
|-|-|-|
|ID|自增ID||
|UserID|用户ID||
|MedicineName|药物名称||
|ActualTime|实际用药时间||
|Memo|打卡备注信息||
|Status|是否在规定时间内，实际用药时间与计划用药时间相差15分钟，记为异常||

## User

|字段|描述|备注|
|-|-|-|
|ID|自增ID||
|NickName|用户昵称||
|Image|用户头像||
|PhoneNum|手机号||
|HUAWEIID|华为用户ID||
|Password|用户密码||

### Interface

#### 用户登录

##### URL
```
POST /medicine/user/login
```
##### params
```json
{
	"phone_num": "15056332824",
	"password": "123456"
}
```

#### 更新用户信息

##### URL
```
PUT /medicine/user/:id
```
##### params
```json
{
	"phone_num": "15056332824",
	"password": "123456",
	"nickname": "admin",
	"image": "",
	"huawei_id": ""
}
```