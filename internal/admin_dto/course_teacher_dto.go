package admin_dto

//教师-课程表 add 入参
type CourseTeacherParam struct {
	CourseID  int `json:"course_id" binding:"required"`
	TeacherId int `json:"teacher_id"`
}
