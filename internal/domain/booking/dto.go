package booking

import "time"

type Request struct {
	resource_id int64      `json:"resource_id" validate:"required"`
	start_time  *time.Time `json:"start_time" validate:"start_time < end_time"`
	end_time    *time.Time `json:"end_time" validate:"end_time > start_time"`
}

type Response struct {
	id          int64
	user_id     int64
	resource_id int64
	start_time  time.Time
	end_time    time.Time
}

type ListResponse struct {
	Data []Response
}
