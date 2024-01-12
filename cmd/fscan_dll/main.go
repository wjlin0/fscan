// go build -buildmode=c-shared -o fscan_go.dll main.go
package main

import "C"

import (
	"fmt"
	"time"

	"github.com/wjlin0/fscan/Plugins"
	"github.com/wjlin0/fscan/common"
)

//export DllCanUnloadNow
func DllCanUnloadNow() {}

//export DllGetClassObject
func DllGetClassObject() {}

//export DllRegisterServer
func DllRegisterServer() {}

//export DllUnregisterServer
func DllUnregisterServer() {}

func init() {
	start := time.Now()
	var Info common.HostInfo
	common.Flag(&Info)
	common.Parse(&Info)
	Plugins.Scan(Info)
	t := time.Now().Sub(start)
	fmt.Printf("[*] 扫描结束,耗时: %s\n", t)
}

func main() {}
