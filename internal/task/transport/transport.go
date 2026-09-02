package transport

import (
	"ToDoList/internal/task/service"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TaskDTO struct {
	ID             int        `json:"id"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	CreationTime   time.Time  `json:"creation_time"`
	CompletionTime *time.Time `json:"completion_time,omitempty"`
}

func DTOtoModel(dto TaskDTO) service.TaskModel {
	return service.TaskModel{
		ID:             dto.ID,
		Title:          dto.Title,
		Description:    dto.Description,
		CreationTime:   dto.CreationTime,
		CompletionTime: dto.CompletionTime,
	}
}

func ModelToDTO(model service.TaskModel) TaskDTO {
	return TaskDTO{
		ID:             model.ID,
		Title:          model.Title,
		Description:    model.Description,
		CreationTime:   model.CreationTime,
		CompletionTime: model.CompletionTime,
	}
}

type Handler struct {
	svc service.ServiceInterface
}

func NewHandler(svc service.ServiceInterface) *Handler {
	return &Handler{svc: svc}
}

func writeResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(map[string]any{"data": data})
		if err != nil {
			fmt.Println("Encoding error in writeResponse")
		}
	}
}

func writeError(w http.ResponseWriter, status int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != "" {
		err := json.NewEncoder(w).Encode(map[string]any{"error": err})
		if err != nil {
			fmt.Println("Encoding error in writeResponse")
		}
	}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		writeResponse(w, http.StatusMethodNotAllowed, nil)
		return
	}

	defer r.Body.Close()
	var req TaskDTO
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	serviceMod, err := h.svc.CreateTask(r.Context(), DTOtoModel(req))
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusCreated, ModelToDTO(serviceMod))
}
