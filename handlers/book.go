package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	bookdto "waysbook/dto/book"
	dto "waysbook/dto/result"
	"waysbook/models"
	"waysbook/repositories"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerBook struct {
	BookRepository repositories.BookRepository
}

func HandlerBook(BookRepository repositories.BookRepository) *handlerBook {
	return &handlerBook{BookRepository}
}

func (h *handlerBook) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]

	if userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "You're not admin"}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	dataPDF := r.Context().Value("dataPDF")
	filePDF := dataPDF.(string)

	fmt.Println(dataPDF)

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	fmt.Println(dataContex)

	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "Waysbook"})

	if err != nil {
		fmt.Println(err.Error())
	}

	// file, err := cld.Upload.Upload(ctx, filePDF, uploader.UploadParams{Folder: "Waysbook"});

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	pages, _ := strconv.Atoi(r.FormValue("pages"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	isbn, _ := strconv.Atoi(r.FormValue("isbn"))

	request := bookdto.CreateBook{
		Title:           r.FormValue("title"),
		PublicationDate: r.FormValue("publication_date"),
		Pages:           pages,
		ISBN:            isbn,
		Author:          r.FormValue("author"),
		Price:           price,
		Description:     r.FormValue("description"),
	}

	validation := validator.New()

	err = validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	publicDate, _ := time.Parse("2006-01-02", request.PublicationDate)

	book := models.Book{
		Title:           request.Title,
		PublicationDate: publicDate,
		Pages:           request.Pages,
		ISBN:            request.ISBN,
		Author:          request.Author,
		Price:           request.Price,
		Description:     request.Description,
		FilePDF:         filePDF,
		Status:          "regular",
		Image:           resp.SecureURL,
		CreateAt:        time.Now(),
	}

	createBook, err := h.BookRepository.CreateBook(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: createBook}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) FindBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := h.BookRepository.FindBook()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: books}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	book, err := h.BookRepository.GetBook(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: book}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	dataPDF := r.Context().Value("dataPDF")
	filePDF := dataPDF.(string)

	pages, _ := strconv.Atoi(r.FormValue("pages"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	isbn, _ := strconv.Atoi(r.FormValue("isbn"))

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "Waysbook"})
	if err != nil {
		fmt.Println(err.Error())
	}

	// file, err := cld.Upload.Upload(ctx, filePDF, uploader.UploadParams{Folder: "Waysbook", ResourceType: "pdf"});
	// if err != nil {
	//   fmt.Println(err.Error())
	// }

	request := bookdto.UpdateBook{
		Title:           r.FormValue("title"),
		PublicationDate: r.FormValue("publication_date"),
		Pages:           pages,
		ISBN:            isbn,
		Author:          r.FormValue("author"),
		Price:           price,
		Description:     r.FormValue("description"),
	}

	validation := validator.New()

	err = validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	book, _ := h.BookRepository.GetBook(id)

	publicDate, _ := time.Parse("2006-01-02", request.PublicationDate)
	if request.Title != "" {
		book.Title = request.Title
	}

	if request.PublicationDate != "" {
		book.PublicationDate = publicDate
	}

	if request.Pages != 0 {
		book.Pages = request.Pages
	}

	if isbn != 0 {
		book.ISBN = isbn
	}

	if request.Author != "" {
		book.Author = request.Author
	}

	if request.Price != 0 {
		book.Price = request.Price
	}

	if request.Description != "" {
		book.Description = request.Description
	}

	if filepath != "false" {
		book.Image = resp.SecureURL
	}

	if filePDF != "false" {
		book.FilePDF = filePDF
	}

	if request.Status != "" {
		book.Status = request.Status
	}

	book.UpdateAt = time.Now()

	book, _ = h.BookRepository.UpdateBook(book)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: book}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	book, err := h.BookRepository.GetBook(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.BookRepository.DeleteBook(book)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) FindBookPromo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := h.BookRepository.FindBookPromo()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: books}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) FindBookRegular(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := h.BookRepository.FindBookRegular()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: books}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) UpdateBookPromo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	request := new(bookdto.UpdateBookPromo)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	book, _ := h.BookRepository.GetBook(id)

	book.Status = request.Status

	books, _ := h.BookRepository.UpdateBook(book)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "Success", Data: books}
	json.NewEncoder(w).Encode(response)
}
