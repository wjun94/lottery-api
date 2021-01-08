package model

// Msg 用户留言
type Msg struct {
	ID     string `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID string `json:"userId"` // 用户id
	Name   string `json:"name"`   // 标题
	Friend string `json:"friend"` // 对方的ID
	SendAt string `json:"sendAt"` // 发送时间
	ReadAt string `json:"readAt"` // 阅读时间
	Desc   string `json:"desc"`   // 描述
	Phone  string `json:"phone"`  // 联系手机
	Status bool   `json:"status"` // 留言状态
}
