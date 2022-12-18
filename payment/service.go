package payment

import (
	"github.com/veritrans/go-midtrans"
	"os"
	"project_dwi/users"
	"strconv"
)

type service struct{}

type Service interface {
	GetPaymentURL(transaction Transaction, user users.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user users.User) (string, error) {
	midClient := midtrans.NewClient()
	SERVER_KEY := os.Getenv("MIDTRANS_SERVER_KEY")
	CLIENT_KEY := os.Getenv("MIDTRANS_CLIENT_KEY")
	midClient.ServerKey = SERVER_KEY
	midClient.ClientKey = CLIENT_KEY
	midClient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midClient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil
}
