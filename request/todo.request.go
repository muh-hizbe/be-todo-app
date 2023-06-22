package request

type TodoCreateRequest struct {
	Name       string `json:"name" form:"name" validate:"required"`
	Note       string `json:"note" form:"note"`
	IsComplete bool   `json:"is_complete" form:"is_complete"`
}

type TodoUpdateRequest struct {
	Name       string `json:"name" form:"name" validate:"required"`
	Note       string `json:"note" form:"note"`
	IsComplete bool   `json:"is_complete" form:"is_complete"`
}
