package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mafuyuk/ddd-go-api-template/domain"

	"github.com/go-chi/chi/middleware"
)

type postRequestUser struct {
	Name        string `json:"name" validate:"required,min=1,max=30"`
	Description string `json:"description" validate:"min=1,max=100"`
}

type responseUser struct {
	ID          domain.UserID `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	RequestID   string        `json:"request_id"`
}

func (s *Server) registerUser(w http.ResponseWriter, r *http.Request) {
	log.Println("called registerUser")
	ctx := r.Context()
	requestID := middleware.GetReqID(ctx)

	// リクエストボディの取得
	var requestData postRequestUser
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		log.Println(err)
		rendering.JSON(w, http.StatusBadRequest, &responseError{
			RequestID: requestID,
			Reason:    "Failed request body decode.",
			Error:     err.Error(),
		})
		return
	}

	// バリデーション
	if err := validate.Struct(requestData); err != nil {
		rendering.JSON(w, http.StatusBadRequest, &responseError{
			RequestID: requestID,
			Reason:    "Failed request body validate.",
			Error:     err.Error(),
		})
		return
	}

	// ビジネスロジック処理
	user := &domain.User{
		Name:        requestData.Name,
		Description: requestData.Description,
	}
	if err := s.UserService.Register(ctx, user); err != nil {
		rendering.JSON(w, http.StatusInternalServerError, &responseError{
			RequestID: requestID,
			Reason:    "Failed UserService.Register.",
			Error:     err.Error(),
		})
		return
	}

	// 正常時のレスポンス
	rendering.JSON(w, http.StatusCreated, &responseUser{
		RequestID:   requestID,
		ID:          user.ID,
		Name:        user.Name,
		Description: user.Description,
	})
}
