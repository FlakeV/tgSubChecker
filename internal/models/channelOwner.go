package models

type ChannelOwner struct {
	OwnerID       int  `json:"owner_id"`
	Notifications bool `json:"notifications"`
}
