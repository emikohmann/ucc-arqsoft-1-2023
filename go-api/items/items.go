package items

import "errors"

func Division (a int,b int ) (int, error){
	if b == 0 {
		return -1, errors.New("b no puede ser cero")
	}
	return a / b, nil
}

