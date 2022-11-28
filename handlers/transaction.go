package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	dto "waysbucks/dto/result"
	transactiondto "waysbucks/dto/transaction"
	"waysbucks/models"
	"waysbucks/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(transactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{transactionRepository}
}

func (h *handlerTransaction) AddTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	request := new(transactiondto.AddTransaction)
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

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	accountID := int(userInfo["id"].(float64))

	orderData, _ := h.TransactionRepository.GetOrderByUser(accountID)

	var totalTransaction = 0

	for _, i := range orderData {
		totalTransaction += i.SubTotal
	}

	dataOrders := models.Transaction{
		Name:      request.Name,
		Address:   request.Address,
		Status:    "Pending",
		Order:     orderData,
		Total:     totalTransaction,
		AccountID: accountID,
	}

	itemTransaction, err := h.TransactionRepository.AddTransaction(dataOrders)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Checkout Failed!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	orderItem, _ := h.TransactionRepository.GetTransaction(itemTransaction.ID)

	data := models.Transaction{
		ID:        orderItem.ID,
		Name:      request.Name,
		Address:   request.Address,
		Status:    "Pending",
		Order:     orderData,
		Total:     totalTransaction,
		AccountID: orderItem.AccountID,
		Account:   orderItem.Account,
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerTransaction) CancelTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TransactionRepository.CancelTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerTransaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	cekRole := userInfo["role"]

	if cekRole != "admin" {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "You can't Access!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	request := new(transactiondto.UpdateTransaction)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	transaction, err := h.TransactionRepository.GetTransaction(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Status != "" {
		transaction.Status = request.Status
	}

	data, err := h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
