package serializers

type UserSignUp struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type UserSignIn struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type AddMonitorIn struct {
	Name      string `json:"name" validate:"required"`
	URL       string `json:"url" validate:"required"`
	Type      string `json:"type" validate:"required"`
	Frequency int    `json:"frequency" validate:"required"`
	Method    string `json:"method" validate:"required"`
}
