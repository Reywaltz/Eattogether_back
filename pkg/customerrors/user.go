package customerrors

type UserNotSetError struct{}

func (u *UserNotSetError) Error() string {
	return "user_id not set in echo context"
}
