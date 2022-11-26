package authdto

type RegisterRequest struct {
	Username string `gorm:"type: text" json:"username" validate:"required"`
	Email    string `gorm:"type: text" json:"email" validate:"required"`
	Password string `gorm:"type: text" json:"password" validate:"required"`
	Role     string `gorm:"type: text" json:"role" validate:"required"`
}

type LoginRequest struct {
	Email    string `gorm:"type: text" json:"email" validate:"required"`
	Password string `gorm:"type: text" json:"password" validate:"required"`
}
