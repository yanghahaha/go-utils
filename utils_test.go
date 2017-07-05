package wooUtils

import (
	"fmt"
	"testing"
)

func Test_StrFormatTime(t *testing.T) {
	fmt.Println(StrFormatTime("%Y-%m-%d %H-%i-%s"))
}

func Test_Try(t *testing.T) {
	Try(func() {
		panic("test Exception")
	}, func(exception interface{}) {
		fmt.Println(exception)
	})
	fmt.Println("ok")
}

func Test_GetCurrentPath(t *testing.T) {
	fmt.Println(GetCurrentPath())
}
