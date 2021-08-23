package model

type User struct {
	Id        int64  `gorm:"column:id" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	Email     string `gorm:"column:email" json:"email"`
	State     int32  `gorm:"column:state" json:"state"`
	Password  string `gorm:"column:password" json:"password"`
	CreatedAt int64  `gorm:"column:created_at,autoCreatedAt" json:"createdAt"`
	UpdatedAt int64  `gorm:"column:updated_at,autoUpdatedAt" json:"updatedAt"`
}

func (User) TableName() string {
	return "users"
}
