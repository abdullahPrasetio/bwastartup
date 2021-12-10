package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, input CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(createImageInput CreateImageInput, fileName string) (CampaignImage, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignById(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindById(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaigns := Campaign{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		GoalAmount:       input.GoalAmount,
		UserID:           input.User.ID,
		Slug:             slug.Make(slugCandidate),
	}

	// Pembuatan slug

	newCampaign, err := s.repository.Save(campaigns)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, input CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(inputID.ID)
	if err != nil {
		return campaign, err
	}
	if input.User.ID != campaign.UserID {
		return campaign, errors.New("Not an owner of the campaign")
	}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, err
}

func (s *service) SaveCampaignImage(createImageInput CreateImageInput, fileName string) (CampaignImage, error) {
	campaign, err := s.repository.FindById(createImageInput.CampaignID)
	if err != nil {
		return CampaignImage{}, err
	}

	if createImageInput.User.ID != campaign.UserID {
		return CampaignImage{}, errors.New("Not an owner of the campaign")
	}
	isPrimary := 0
	if createImageInput.IsPrimary {
		_, err := s.repository.MarkAllImagesAsNonPrimary(createImageInput.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
		isPrimary = 1
	}
	campaignImage := CampaignImage{
		CampaignID: createImageInput.CampaignID,
		FileName:   fileName,
		IsPrimary:  isPrimary,
	}

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}
	return campaignImage, nil

}
