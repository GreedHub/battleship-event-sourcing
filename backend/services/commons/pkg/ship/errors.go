package ship

import (
	"errors"
)

func HandleAlreadyPositionedErr() error {
	return errors.New("Ship has already been positioned")
}
