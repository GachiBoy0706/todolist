package transport

import (
	"ToDoList/internal/task/service"
	"encoding/json"
	"net/http"
	"time"
)

type TaskDTO struct {
	ID             int        `json:"id"`
	Topic          string     `json:"topic"`
	Description    string     `json:"description"`
	CreationTime   time.Time  `json:"creation_time"`
	CompletionTime *time.Time `json:"completion_time,omitempty"`
}

func DTOtoModel(dto TaskDTO) service.TaskModel {
	return service.TaskModel{
		ID:             dto.ID,
		Topic:          dto.Topic,
		Description:    dto.Description,
		CreationTime:   dto.CreationTime,
		CompletionTime: dto.CompletionTime,
	}
}

type Handler struct {
	svc service.ServiceInterface
}

func NewHandler(svc service.ServiceInterface) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()
	var req TaskDTO
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.svc.CreateTask(DTOtoModel(req))

}
