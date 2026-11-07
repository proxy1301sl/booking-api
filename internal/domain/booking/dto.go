package booking

import "time"

type Request struct {
	Resource_id int64      `json:"resource_id" validate:"required,gt=0"`
	Start_time  *time.Time `json:"start_time" validate:"required"`
	End_time    *time.Time `json:"end_time" validate:"required"`
}

type Response struct {
	ID          int64     `json:"id"`
	User_id     int64     `json:"user_id"`
	Resource_id int64     `json:"resource_id"`
	Start_time  time.Time `json:"start_time"`
	End_time    time.Time `json:"end_time"`
}

type ListResponse struct {
	Data []Response
}
