package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"
	dto "waysbucks/dto/result"
	transactiondto "waysbucks/dto/transaction"
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

	request := new(transactiondto.OrderRequest)

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

	product, err := h.OrderRepository.GetProductOrder(request.ProductID)

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

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	buyyerID := int(userInfo["id"].(float64))
	// fmt.Println()

	CekRequestTrans, _ := h.OrderRepository.CheckTransactionUser(buyyerID)

	setTransID := int(time.Now().Unix())

	var transID int
	if CekRequestTrans.ID != 0 {
		transID = CekRequestTrans.ID
	} else {
		requestTrans := models.Transaction{
			ID:      setTransID,
			Name:    "-",
			Address: "-",
			Status:  "Pending",
			Total:   0,
			UserID:  buyyerID,
		}
		transOrder, _ := h.OrderRepository.RequestTransaction(requestTrans)
		transID = transOrder.ID
	}

	// fmt.Println(CekRequestTrans.ID)
	// fmt.Println(transID)
	// fmt.Println(subTotal)

	dataOrder := models.Order{
		Qty:           request.Qty,
		ProductID:     product.ID,
		Toping:        topings,
		SubTotal:      subTotal,
		TransactionID: transID,
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

	myTrans, err := h.OrderRepository.CheckTransactionUser(buyyerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.OrderRepository.FindAllOrderTransaction(myTrans.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// for _, p := range data {
	// 	p.Product.Image = os.Getenv("PATH_FILE") + p.Product.Image
	// }
	// fmt.Println(data)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerOrder) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(transactiondto.UpdateOrder)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	order, err := h.OrderRepository.GetOrder(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	priceItem := order.SubTotal / order.Qty
	order.Qty = request.Qty
	order.SubTotal = priceItem * order.Qty

	data, err := h.OrderRepository.UpdateOrder(order)
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
