package transaction

import "go-startup/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
