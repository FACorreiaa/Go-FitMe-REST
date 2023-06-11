package domain

type Service struct {
	Activity ActivityService
}

type ActivityService interface {
	GetAll() ([]Activity, error)
}
