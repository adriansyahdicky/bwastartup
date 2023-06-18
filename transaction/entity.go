package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"time"
)

type Transactions struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	CreatedAt  time.Time
	Campaign   campaign.Campaign
}
