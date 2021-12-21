package payment

import (
	"go-startup/transaction"
	"go-startup/user"
	"strconv"

	midtrans "github.com/veritrans/go-midtrans"
)

type service struct {
}

type Service interface {
	GetToken(transaction transaction.Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetToken(transaction transaction.Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.Serverkey = "YOUR-VT-SERVER-KEY"
	midclient.Clientkey = "YOUR-VT-CLIENT-KEY"
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway
	snapGateway = midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			Fname: user.Name,
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
