package emitty

import "fmt"

func (s *Signal) log(name, message string, data interface{}, err error) {
	if !s.debug {
		return
	}
	go func() {
		m := fmt.Sprintf("Name: %s | Message: %s | Data %v", name, message, data)
		if err != nil {
			m = fmt.Sprintf("Emitty [ERROR]: %s | Error: %s", m, err.Error())
		} else {
			m = fmt.Sprintf("Emitty [INFO]: %s", m)
		}
		fmt.Println(m)
	}()
}
