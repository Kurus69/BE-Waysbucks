package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	orderdto "waysbucks/dto/order"
	dto "waysbucks/dto/result"
	"waysbucks/models"
	"waysbucks/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerOrder struct {
	OrderRepository repositories.OrderRepository
}

func HandlerOrder(OrderRepository repositories.OrderRepository) *handlerOrder {
	return &handlerOrder{OrderRepository}
}

func (h *handlerOrder) AddOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	buyyerID := int(userInfo["id"].(float64))

	request := new(orderdto.AddOrder)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Check dto"}
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

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.OrderRepository.GetProductOrder(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Product Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	topings, err := h.OrderRepository.GetTopingOrder(request.TopingID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Toping Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	var priceTopings = 0
	for _, i := range topings {
		priceTopings += i.Price
	}
	var subTotal = request.Qty * (product.Price + priceTopings)

	dataOrder := models.Order{
		Qty:       request.Qty,
		BuyyerID:  buyyerID,
		ProductID: product.ID,
		Toping:    topings,
		SubTotal:  subTotal,
	}

	order, err := h.OrderRepository.OrderItem(dataOrder)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Order Failed!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	orderItem, _ := h.OrderRepository.GetOrder(order.ID)

	orderItem.Product.Image = os.Getenv("PATH_FILE") + orderItem.Product.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: orderItem}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerOrder) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	cekRole := userInfo["role"]

	if cekRole != "admin" {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "You Can't Access"}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.OrderRepository.DelOrderItem(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data.Product.Image = os.Getenv("PATH_FILE") + data.Product.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerOrder) GetOrderUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	buyyerID := int(userInfo["id"].(float64))

	cekRole := userInfo["role"]

	if cekRole != "admin" {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "You can't Access!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.OrderRepository.FindAllOrderUser(buyyerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for _, p := range data {
		p.Product.Image = os.Getenv("PATH_FILE") + p.Product.Image
		fmt.Println(p.Product.Image)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
