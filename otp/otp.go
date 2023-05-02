// otp.Println("test") OUTPUTS [main.go:4]: test

package otp

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

var colors = [...]string{
	"\033[31m", // red
	"\033[32m", // green
	"\033[33m", // yellow
	"\033[34m", // blue
	"\033[35m", // magenta
	"\033[36m", // cyan
}

var fileColors = make(map[string]string)

func Println(msg string, skip ...int) {
	if len(skip) == 0 {
		skip = []int{1}
	}

	_, rtpath, line, _ := runtime.Caller(skip[0])
	_, file := filepath.Split(rtpath)
	color, ok := fileColors[file] // get assigned color for file
	if !ok {                      // if color is not assigned for file, assign a new one
		color = colors[rand.Intn(len(colors))]
		fileColors[file] = color // save color for file
	}
	fmt.Printf(fmt.Sprintf("[%s%s\033[0m:%s]: %s\n", color, file, strconv.Itoa(line), msg))
}

func init() {
	rand.Seed(time.Now().UnixNano()) // initialize random seed
}
