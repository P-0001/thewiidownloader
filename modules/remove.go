package modules

import (
	"os"
)

const toRemove = true

func Remove(file string) error {
	/*
		if !toRemove {
			fmt.Println("[WARN] REMOVE FILE IS OFF\n")
			return nil
		}
	*/

	return os.Remove(file)
}
