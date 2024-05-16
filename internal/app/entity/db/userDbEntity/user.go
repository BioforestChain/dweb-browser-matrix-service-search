package userDbEntity

type UserProfile struct {
	UserId        string `json:"user_id"`
	DisplayName   string `json:"display_name"`
	AvatarUrl     string `json:"avatar_url"`
	WalletAddress string `json:"wallet_address"`
}

type OnlineList struct {
	List []*UserProfile `json:"list"`
}

func (UserProfile) TableName() string {
	return "user"
}

type Condition struct {
	Id     uint32
	Page   int
	Limit  int
	Offset int
}
