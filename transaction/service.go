package transaction

import (
	"errors"
	"go-startup/campaign"
)

type service struct {
	repository         Repository
	campaginRepository campaign.Repository
}

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	campaign, err := s.campaginRepository.FindByID(input.ID)

	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}
	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
