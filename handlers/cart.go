package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	cartdto "waysbook/dto/cart"
	dto "waysbook/dto/result"
	"waysbook/models"
	"waysbook/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerCart struct {
	CartRepository repositories.CartRepository
}

func HandlerCart(CartRepository repositories.CartRepository) *handlerCart {
	return &handlerCart{CartRepository}
}

func (h *handlerCart) CreateCart(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int(userInfo["id"].(float64))

	request := new(cartdto.CreateCart)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "cek dto"}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	validation := validator.New()
	err := validation.Struct(request)


	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "error validation"}
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction, err := h.CartRepository.GetTransactionID(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "ID Transaction Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	book, err := h.CartRepository.GetBookCart(request.BookID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Book Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if transaction.ID == 0 {

		transID := int(time.Now().Unix())

		transaction := models.Transaction{
			ID: 			transID,
			UserID: 	userID,
			Status: 	"waiting",
			CreateAt: time.Now(),
		}

		createTrans, err := h.CartRepository.CreateTransaction(transaction)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Cart Failed!"}
			json.NewEncoder(w).Encode(response)
			return
		}

		dataCart := models.Cart {
		BookID: 				request.BookID,
		TransactionID: 	int(createTrans.ID),
		Total: 					book.Price,
		CreateAt: 			time.Now(),
		}

		cart, err := h.CartRepository.CreateCart(dataCart)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Cart Failed!"}
			json.NewEncoder(w).Encode(response)
			return
		}

		res, _ := h.CartRepository.GetCart(cart.ID)

		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: "success", Data: res}
		json.NewEncoder(w).Encode(response)
	} else {
		dataCart := models.Cart {
		BookID: 				request.BookID,
		TransactionID: 	int(transaction.ID),
		Total: 					book.Price,
		CreateAt: 			time.Now(),
		}

		cart, err := h.CartRepository.CreateCart(dataCart)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Cart Failed!"}
			json.NewEncoder(w).Encode(response)
			return
		}

		res, _ := h.CartRepository.GetCart(cart.ID)

		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Code: "success", Data: res}
		json.NewEncoder(w).Encode(response)
	}
}

func (h *handlerCart) DeleteCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	cart, err := h.CartRepository.GetCart(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.CartRepository.DeleteCart(cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	dataDetele := data.ID
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: dataDetele}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCart) GetCartByTransID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userID := int64(userInfo["id"].(float64))

	transaction, _ := h.CartRepository.GetTransactionID(int(userID))

	cart, err := h.CartRepository.GetCartByTransID(int(transaction.ID))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: cart}
	json.NewEncoder(w).Encode(response)
}