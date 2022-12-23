package adminDto

type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginJWT struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type RegisterAdminDto struct {
	RoleId   uint   `json:"role_id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
