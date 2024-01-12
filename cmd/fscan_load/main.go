// gradle build -o fscan_loader_go.exe main.go
package main

import (
	"fmt"
	"syscall"
)

func main() {
	dll, err := syscall.LoadLibrary("fscan_go.dll")
	if err != nil {
		fmt.Printf("[-] Failed to load DLL: %s\n", err)
		return
	}

	defer func(handle syscall.Handle) {
		err := syscall.FreeLibrary(handle)
		if err != nil {
			fmt.Printf("[-] Failed to free DLL: %s\n", err)
		}
	}(dll)

	// 定义并尝试获取 DLL 中的函数
	dllCanUnloadNow, err := syscall.GetProcAddress(dll, "DllCanUnloadNow")
	// 判断err == nil 并且 dllCanUnloadNow != 0
	if err == nil && dllCanUnloadNow != 0 {
		_, _, err := syscall.SyscallN(dllCanUnloadNow, 0, 0, 0, 0)
		if err != 0 {
			fmt.Printf("[-] Failed to execute DllCanUnloadNow: %s\n", err)
		}
	}

	dllGetClassObject, err := syscall.GetProcAddress(dll, "DllGetClassObject")
	if err == nil && dllGetClassObject != 0 {
		_, _, err := syscall.SyscallN(dllGetClassObject, 0, 0, 0, 0)
		if err != 0 {
			fmt.Printf("[-] Failed to execute DllGetClassObject: %s\n", err)
		}
	}

	dllRegisterServer, err := syscall.GetProcAddress(dll, "DllRegisterServer")
	if err == nil && dllRegisterServer != 0 {
		_, _, err := syscall.SyscallN(dllRegisterServer, 0, 0, 0, 0)
		if err != 0 {
			fmt.Printf("[-] Failed to execute DllRegisterServer: %s\n", err)
		}
	}

	dllUnregisterServer, err := syscall.GetProcAddress(dll, "DllUnregisterServer")
	if err == nil && dllUnregisterServer != 0 {
		_, _, err := syscall.SyscallN(dllUnregisterServer, 0, 0, 0, 0)
		if err != 0 {
			fmt.Printf("[-] Failed to execute DllUnregisterServer: %s\n", err)
		}
	}

	fmt.Println("All functions executed successfully.")
}
