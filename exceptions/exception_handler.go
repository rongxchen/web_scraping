package exceptions

import "log"

func HandleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
