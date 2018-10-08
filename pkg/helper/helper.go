package helper

import (
	"math/rand"
	"path"
	"runtime"
	"time"
)

const (
	OffStatus = 1
	OnStatus  = 0
	Timeout   = 300
)

func init() {
	rand.Seed(time.Now().Unix())
}

func RandInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandBool() bool {
	return RandInt(0, 2) == 0
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetRootPath() string {
	_, filename, _, _ := runtime.Caller(0)
	path := path.Join(path.Dir(filename), "../../cmd")
	return path
}
