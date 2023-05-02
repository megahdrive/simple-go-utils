package vld

import "Program/otp"

func Check(err error) {
	if err != nil {
		otp.Println("\033[31m"+err.Error()+"\033[0m", 2)
	}
}
