package user

type CreateUserDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type GetUsersDto struct {
	Id       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type UpdateUserDto struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}

type RateDto struct {
	Rating  int    `json:"rating" binding:"required"`
	Comment string `json:"comment"`
}

type GetAllRatingsDto struct {
	Rating  float32 `json:"rating"`
	Comment string  `json:"comment"`
	User    string  `json:"user"`
	RatedBy string  `json:"rated_by"`
}

type ApplicationDto struct {
	UserId   string `json:"user_id"`
	Filename string `json:"filename"`
}
