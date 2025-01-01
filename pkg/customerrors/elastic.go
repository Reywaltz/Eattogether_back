package customerrors

type ESDataNotFound struct{}

func (e *ESDataNotFound) Error() string {
	return "es data not found"
}
