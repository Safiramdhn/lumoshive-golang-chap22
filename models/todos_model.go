package models

type Todos struct {
	ID                  int    `json:"id"`
	UserID              int    `json:"user_id"`
	Description         string `json:"description"`
	TodoStatus          string `json:"todo_status"`
	TotalTodoNotStarted int    `json:"total_todo_not_started"`
	TotalTodoDone       int    `json:"total_todo_done"`
}
