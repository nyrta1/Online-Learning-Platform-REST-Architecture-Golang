package forms

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateForm struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type LessonForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	VideoUrl    string `json:"videoUrl"`
}

type CourseForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
