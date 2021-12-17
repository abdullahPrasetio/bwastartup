package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignImages := ""
	if len(campaign.CampaignImages) > 0 {
		campaignImages = campaign.CampaignImages[0].FileName
	}
	campaignFormatter := CampaignFormatter{
		ID:               campaign.ID,
		UserId:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         campaignImages,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.GoalAmount,
		Slug:             campaign.Slug,
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	// Cara 1
	// if len(campaigns) == 0 {
	// 	return []CampaignFormatter{}
	// }
	// cara 2
	campaignsFormatter := []CampaignFormatter{}
	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID               int                       `json:"id"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"short_description"`
	Description      string                    `json:"description"`
	ImageURL         string                    `json:"image_url"`
	GoalAmount       int                       `json:"goal_amount"`
	CurrentAmount    int                       `json:"current_amount"`
	BackerCount      int                       `json:"backer_count"`
	UserID           int                       `json:"user_id"`
	Slug             string                    `json:"slug"`
	Perks            []string                  `json:"perks"`
	User             CampaignUserFormatter     `json:"user"`
	Images           []CampaignImagesFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImagesFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {

	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.BackerCount = campaign.BackerCount
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.ImageURL = ""
	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailFormatter.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{
		Name: user.Name, ImageURL: user.AvatarFileName,
	}

	campaignDetailFormatter.User = campaignUserFormatter

	images := []CampaignImagesFormatter{}
	if len(campaign.CampaignImages) > 0 {
		for _, image := range campaign.CampaignImages {
			isPrimary := false
			if image.IsPrimary == 1 {
				isPrimary = true
			}
			campaignImageFormatter := CampaignImagesFormatter{
				ImageUrl:  image.FileName,
				IsPrimary: isPrimary,
			}
			images = append(images, campaignImageFormatter)
		}
	}
	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}
