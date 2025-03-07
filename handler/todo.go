package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		request := &model.CreateTODORequest{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Println(err)
			return
		}
		if request.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		todo, err := h.svc.CreateTODO(r.Context(), request.Subject, request.Description)
		if err != nil {
			log.Println(err)
			return
		}
		response := model.CreateTODOResponse{TODO: *todo}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println(err)
			return
		}
	}

	if r.Method == "PUT" {
		request := &model.UpdateTODORequest{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Println(err)
			return
		}
		if request.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		todo, err := h.svc.UpdateTODO(r.Context(), request.ID, request.Subject, request.Description)
		if err != nil {
			switch err := err.(type) {
			case *model.ErrNotFound:
				w.WriteHeader(http.StatusNotFound)
			default:
				log.Println(err)
			}
			return
		}
		response := model.CreateTODOResponse{TODO: *todo}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println(err)
			return
		}
	}
	if r.Method == "GET" {
		request := &model.ReadTODORequest{}
		query := r.URL.Query()
		request.PrevID, _ = strconv.ParseInt(query.Get("prev_id"), 10, 64)
		request.Size, _ = strconv.ParseInt(query.Get("size"), 10, 64)
		if request.Size == 0 {
			request.Size = 5
		}
		todos, err := h.svc.ReadTODO(r.Context(), request.PrevID, request.Size)
		if err != nil {
			log.Println(err)
			return
		}
		response := model.ReadTODOResponse{TODOs: todos}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println(err)
			return
		}
	}
	if r.Method == "DELETE" {
		request := &model.DeleteTODORequest{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			log.Println(err)
			return
		}
		if len(request.IDs) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := h.svc.DeleteTODO(r.Context(), request.IDs)
		if err != nil {
			switch err := err.(type) {
			case *model.ErrNotFound:
				w.WriteHeader(http.StatusNotFound)
			default:
				log.Println(err)
			}
			return
		}
		response := model.DeleteTODOResponse{}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println(err)
			return
		}
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}
