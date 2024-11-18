package additions

import (
	"eattogether/pkg/customerrors"
	"fmt"

	"github.com/labstack/echo/v4"
)

// TODO Подумать как сделать функцию более универсальной для всех кейсов (path/query params, данные с формы)
// Посмотреть в сторону либо кастомного байндера, где все будет лежать и возвращаться, либо
// с использованием дефолтного
func RetriveUserAndPayload(c echo.Context, bindInterface interface{}, skipBind bool) (int, error) {
	userID := c.Get("user_id").(int)

	if userID == 0 {
		return 0, &customerrors.UserNotSetError{}
	}

	if !skipBind {
		err := c.Bind(bindInterface)
		if err != nil {
			fmt.Printf("Can't bind to presented struct: %v", err)
			return 0, &customerrors.DataNotBindable{
				Struct: bindInterface,
			}
		}
	}

	fmt.Println(bindInterface)

	return userID, nil
}
