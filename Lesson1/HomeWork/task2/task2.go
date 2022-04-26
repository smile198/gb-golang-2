package task2

import (
	"errors"
	"fmt"
	"time"
)

func Run() (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = ErrGotPanic{
				time:       time.Now(),
				innerError: getError(e, "unknown panic"),
			}
		}
	}()

	gimmePanic()

	return err
}

func getError(e interface{}, defaultErrorText string) (err error) {
	switch x := e.(type) {
	case string:
		err = errors.New(x)
	case error:
		err = x
	default:
		err = errors.New(defaultErrorText)
	}

	return err
}

func gimmePanic() {
	panic("AAAAA")
}

type ErrGotPanic struct {
	time       time.Time
	innerError error
}

func (e ErrGotPanic) Error() string {
	return fmt.Sprintf("Error occured at: %s, reason: %s", e.time.Format("15:04:05 02.01.2006"), e.innerError.Error())
}
