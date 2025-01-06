package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pujaprabha/internal/domain/models"
	"pujaprabha/internal/presenter"
)

func (repo *repository) VerifyPayment(req presenter.VerifyPaymentPresenter) (*models.Payment, error) {

	var requestBody presenter.VerifyPaymentList

	client := &http.Client{}
	url := fmt.Sprintf("https://uat.esewa.com.np/api/epay/transaction/status/?product_code=EPAYTEST&total_amount=%v&transaction_uuid=%v", req.TotalAmount, req.TransactionUUID)
	// url := "https://uat.esewa.com.np/api/epay/transaction/status/?product_code=EPAYTEST&total_amount=req.&transaction_uuid=123"
	request, _ := http.NewRequest("GET", url, nil)
	response, err := client.Do(request)
	if response.StatusCode != 200 {
		return nil, err
	}
	defer response.Body.Close()
	// fmt.Printf("response.Body: %v\n", response.Body)

	b, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(b, &requestBody); err != nil {
		return nil, err
	}
	fmt.Printf("requestBody.Status: %v\n", requestBody.Status)
	if requestBody.Status == "COMPLETE" {
		err = repo.db.Create(&models.Payment{
			UserID:          uint(req.UserID),
			ProductID:       req.ProductID,
			TransactionID:   req.TransactionID,
			TransactionUUID: req.TransactionUUID,
			PaymentMode:     req.PaymentMode,
			TotalAmount:     req.TotalAmount,
		}).Error
		return nil, err
	}

	return nil, err
}
