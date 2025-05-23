package request

type User struct {
	Email    string  `json:"email" db:"email"`
	Password string  `json:"password" db:"password"`
	Phone    *string `json:"phone" db:"phone"`
	Address  *string `json:"address" db:"address"`
	Nickname *string `json:"nickname" db:"nickname"`
	Logo     *string `json:"logo" db:"logo"`
}

type UserLogin struct {
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required,min=6"`
}
