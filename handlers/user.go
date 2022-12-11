package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	dto "waysbook/dto/result"
	usersdto "waysbook/dto/user"
	"waysbook/models"
	"waysbook/pkg/bcrypt"
	"waysbook/repositories"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)


type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser{
	return &handlerUser{UserRepository}
}

func (h *handlerUser) GetUserByID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(userInfo["id"].(float64))

	user, err := h.UserRepository.GetUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: user}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(userInfo["id"].(float64))

	dataContext := r.Context().Value("dataFile") // add this code
	filepath := dataContext.(string)

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "Waysbook"});

		if err != nil {
		fmt.Println(err.Error())
	}
  request := usersdto.UpdateUser{
		Fullname:  r.FormValue("fullname"),
		Password: r.FormValue("password"),
		Address: r.FormValue("address"),
		Phone: r.FormValue("phone"),
		Gender: r.FormValue("gender"),
	}

	user, err := h.UserRepository.GetUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Fullname != "" {
		user.Fullname = request.Fullname
	}

	if request.Password != "" {
		password, _ := bcrypt.HashingPassword(request.Password)
		user.Password = password
	}

	if filepath != "false" {
		user.Image = resp.SecureURL
	}
	
	if request.Address != "" {
		user.Address = request.Address
	}

	if request.Phone != "" {
		user.Phone = request.Phone
	}

	if request.Gender != "" {
		user.Gender = request.Gender
	}

	user.UpdateAt= time.Now()

	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]
	userID := int(userInfo["id"].(float64))


	if userID != id && userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "You're not admin"}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.UserRepository.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}


func convertResponse(u models.User) usersdto.DeleteResponse {
	return usersdto.DeleteResponse{
		ID:       u.ID,
	}
}
