package handler

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"payment-options/internal/model"
)

// PaymentHandler handles payment related requests
type PaymentHandler struct{}

// NewPaymentHandler creates a new payment handler
func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{}
}

// GetPaymentOptions returns available payment options
func (h *PaymentHandler) GetPaymentOptions(w http.ResponseWriter, r *http.Request) {
	// Set response header
	w.Header().Set("Content-Type", "application/json")

	// Create data structure to store results
	data := make(map[string]*model.PaymentOption)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	// Define payment methods
	paymentMethods := []string{"ovo", "dana", "gopay", "shopeepay", "oneklik", "bridd", "linkaja"}

	// Process each payment method in parallel
	for _, method := range paymentMethods {
		wg.Add(1)
		go func(method string) {
			defer wg.Done()

			var option *model.PaymentOption

			switch method {
			case "ovo":
				option = getOVOProfile()
			case "dana":
				option = getDANAProfile()
			case "gopay":
				option = getGopayProfile()
			case "shopeepay":
				option = getShopeepayProfile()
			case "oneklik":
				option = getOneKlikProfile()
			case "bridd":
				option = getBRIDDProfile()
			case "linkaja":
				option = getLinkAjaProfile()
			}

			// Store result in map with mutex protection
			if option != nil {
				mutex.Lock()
				data[method] = option
				mutex.Unlock()
			}
		}(method)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Create response
	response := model.Response{
		ReturnCode: "200",
		ReturnDesc: "success",
		Data:       data,
	}

	// Marshal and send response
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}

// Mock API call functions

func getOVOProfile() *model.PaymentOption {
	// Simulate API call with delay
	time.Sleep(200 * time.Millisecond)
	return &model.PaymentOption{
		Account: "628812345678",
		Status:  "Active",
		Balance: "10000",
		Icon:    "https://sampleurl.com/ovo.jpg",
	}
}

func getDANAProfile() *model.PaymentOption {
	time.Sleep(150 * time.Millisecond)
	return &model.PaymentOption{
		Account: "628823456789",
		Status:  "Active",
		Balance: "15000",
		Icon:    "https://sampleurl.com/dana.jpg",
	}
}

func getGopayProfile() *model.PaymentOption {
	time.Sleep(180 * time.Millisecond)
	return &model.PaymentOption{
		Account: "628834567890",
		Status:  "Active",
		Balance: "25000",
		Icon:    "https://sampleurl.com/gopay.jpg",
	}
}

func getShopeepayProfile() *model.PaymentOption {
	time.Sleep(160 * time.Millisecond)
	return &model.PaymentOption{
		Account: "628845678901",
		Status:  "Active",
		Balance: "30000",
		Icon:    "https://sampleurl.com/shopeepay.jpg",
	}
}

func getOneKlikProfile() *model.PaymentOption {
	time.Sleep(170 * time.Millisecond)
	return &model.PaymentOption{
		Account: "628856789012",
		Status:  "Active",
		Balance: "50000",
		Icon:    "https://sampleurl.com/oneklik.jpg",
	}
}

func getBRIDDProfile() *model.PaymentOption {
	time.Sleep(190 * time.Millisecond)
	return &model.PaymentOption{
		Account: "628867890123",
		Status:  "Active",
		Balance: "100000",
		Icon:    "https://sampleurl.com/bridd.jpg",
	}
}

func getLinkAjaProfile() *model.PaymentOption {
	time.Sleep(210 * time.Millisecond)
	return &model.PaymentOption{
		Account: "628878901234",
		Status:  "Active",
		Balance: "75000",
		Icon:    "https://sampleurl.com/linkaja.jpg",
	}
}
