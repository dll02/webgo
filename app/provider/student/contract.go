package student

const StudentKey = "App:student"

type IService interface {
	GetAllStudent() []Student
}

type Student struct {
	ID   int
	Name string
}
