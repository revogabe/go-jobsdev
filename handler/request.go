package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateJobs
type CreateJobsRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Role string `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote bool `json:"remote"`
	Link string `json:"link"`
	Experience string `json:"experience"`
	Salary string `json:"salary"`
	Approved bool `json:"approved"`
}

func (r *CreateJobsRequest) Validate() error {
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Experience == "" {
		return errParamIsRequired("experience", "string")
	}
	if r.Salary == "" {
		return errParamIsRequired("salary", "string")
	}
	if r.Approved {
		return errParamIsRequired("approved", "bool")
	}
	return nil
}

type ApprovedJobsRequest struct {
	Approved bool `json:"approved"`
}
