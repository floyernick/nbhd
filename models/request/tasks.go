package request

type TasksCreate struct {
	Token         string    `json:"token" validate:"required,uuid"`
	Title         string    `json:"title" validate:"required,min=1,max=75"`
	Category      int       `json:"category" validate:"required"`
	Location      []float64 `json:"location" validate:"required,len=2"`
	Description   string    `json:"description" validate:"required,min=1,max=280"`
	Time          int       `json:"time" validate:"omitempty"`
	Encouragement int       `json:"encouragement" validate:"required,min=1,max=3"`
	Pay           float64   `json:"pay" validate:"omitempty"`
}

type TasksGet struct {
	Token string `json:"token" validate:"required,uuid"`
	Id    string `json:"id" validate:"required,uuid"`
}

type TasksList struct {
	Token string `json:"token" validate:"required,uuid"`
}

type TasksDelete struct {
	Token string `json:"token" validate:"required,uuid"`
	Id    string `json:"id" validate:"required,uuid"`
}

type TasksPerformanceRequest struct {
	Token string `json:"token" validate:"required,uuid"`
	Id    string `json:"id" validate:"required,uuid"`
}

type TasksPerformanceCancel struct {
	Token string `json:"token" validate:"required,uuid"`
	Id    string `json:"id" validate:"required,uuid"`
}

type TasksPerformerAccept struct {
	Token  string `json:"token" validate:"required,uuid"`
	TaskId string `json:"task_id" validate:"required,uuid"`
	UserId string `json:"user_id" validate:"required,uuid"`
}

type TasksPerformerDecline struct {
	Token  string `json:"token" validate:"required,uuid"`
	TaskId string `json:"task_id" validate:"required,uuid"`
	UserId string `json:"user_id" validate:"required,uuid"`
}

type TasksRate struct {
	Token  string `json:"token" validate:"required,uuid"`
	Id     string `json:"id" validate:"required,uuid"`
	Rating int    `json:"rating" validate:"required,min=1,max=5"`
}
