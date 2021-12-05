package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserId           int    `json:"user_id`
	Name             string `json:"name`
	ShortDescription string `json:"short_description`
	ImageURL         string `json:"image_url`
	GoalAmount       int    `json:"goal_amount`
	CurrentAmount    int    `json:"current_amount`
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
