package utils

import (
    "errors"
)

func GetInput(word string) (bool,error){
	switch word{
	case "foo":
		return true,nil
	case "bar":
		err := errors.New("bad word!!!")
		return false,err
	}
	
	return false, nil
}
