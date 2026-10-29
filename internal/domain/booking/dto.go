package booking

import "time"

type Request struct {
	resource_id string
	start_time  time.Time
	end_time    time.Time
}

type Flat struct {
	id string
	user_id string
	resource_id string
	start_time  time.Time
	end_time    time.Time
	created_at time.Time

}
