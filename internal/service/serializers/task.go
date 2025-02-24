package serializers

type TaskCreateRequest struct {
	Description string `json:"description" binding:"required"`
	Title       string `json:"title" binding:"required"`
}

type TaskCreateResponse struct {
	ID string `json:"id"`
}

type TaskGetAllResponse struct {
	Tasks []TaskGetAllResponseItem `json:"tasks"`
}

type TaskGetAllResponseItem struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
}

type TaskGetByIdResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
}
