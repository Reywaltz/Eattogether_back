package models

// TODO validations https://echo.labstack.com/docs/request#validate-data
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTResponse struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type JSONMessage struct {
	Message string `json:"message"`
}
