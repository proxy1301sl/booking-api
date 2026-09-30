package validator

import "booking-api/cmd/api"

func init() {
	validate = api.Validate
}
