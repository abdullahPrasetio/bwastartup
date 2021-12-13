package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}
func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindById(input.ID)
	if err != nil {
		return []Transaction{}, err
	}
	if input.User.ID != campaign.UserID {
		return []Transaction{}, errors.New("Not an owner of the campaign")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
