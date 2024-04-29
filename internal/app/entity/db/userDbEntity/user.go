package userDbEntity

type UserInfo struct {
	Id        uint32 `json:"id"`        // User ID
	Passport  string `json:"passport"`  // User Passport
	Password  string `json:"password"`  // User Password
	Nickname  string `json:"nickname"`  // User Nickname
	CreateAt  string `json:"create_at"` // Created Time
	UpdateAt  string `json:"update_at"` // Updated Time
	DeletedAt string `json:"deleted_at"`
}

type OnlineList struct {
	List []*UserInfo `json:"list"`
}

func (UserInfo) TableName() string {
	return "user"
}

type Condition struct {
	Id     uint32
	Page   int
	Limit  int
	Offset int
}

type ConditionUser struct {
	Id       uint32 `json:"id"`       // User ID
	Passport string `json:"passport"` // User Passport
	Password string `json:"password"` // User Password
	Nickname string `json:"nickname"` // User Nickname
}
