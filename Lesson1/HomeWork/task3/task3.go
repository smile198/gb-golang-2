package task3

import (
	"os"
	"strconv"
)

func Run(path string, fileLimit int) error {
	for i := 0; i < fileLimit; i++ {
		err := func() error {
			file, err := os.Create(path + "/file_" + strconv.Itoa(i) + ".txt")
			defer file.Close()

			if err != nil {
				return err
			}

			return nil
		}()

		if err != nil {
			return err
		}
	}

	return nil
}
