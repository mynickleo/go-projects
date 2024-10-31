package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"redis-service/internal/cache"
)

type JsonResponse struct {
	Message string `json:"message"`
}

type CacheHandler struct {
	cache cache.Cache
}

func NewCacheHandler(cache cache.Cache) *CacheHandler {
	return &CacheHandler{cache: cache}
}

func setJsonResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(JsonResponse{Message: message})
}

func (h *CacheHandler) SetKey(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		setJsonResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := h.cache.Set(data.Key, data.Value); err != nil {
		setJsonResponse(w, http.StatusInternalServerError, "Failed to set cache")
		return
	}

	setJsonResponse(w, http.StatusCreated, "Key set successfully")
}

func (h *CacheHandler) GetKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	if key == "" {
		setJsonResponse(w, http.StatusBadRequest, "Missing key parametr")
		return
	}

	value, err := h.cache.Get(key)
	if err != nil {
		setJsonResponse(w, http.StatusNotFound, "Key not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(JsonResponse{Message: fmt.Sprintf("Value: %s", value)})
}

func (h *CacheHandler) DeleteKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	if key == "" {
		setJsonResponse(w, http.StatusBadRequest, "Missing key parametr")
		return
	}

	if err := h.cache.Delete(key); err != nil {
		setJsonResponse(w, http.StatusInternalServerError, "Failter to delete key")
		return
	}

	setJsonResponse(w, http.StatusOK, "Key deleted successfully")
}
