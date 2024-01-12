// gcc -o fscan_loader_c.exe main.c

#include <stdio.h>
#include <windows.h>

int main() {
    HMODULE hDLL;
    hDLL = LoadLibrary("fscan.dll");

    if (hDLL != NULL) {
        // 定义函数指针类型
        typedef void (*FunctionType)();

        // 为每个导出函数创建一个函数指针
        FunctionType DllCanUnloadNow = (FunctionType)GetProcAddress(hDLL, "DllCanUnloadNow");
        FunctionType DllGetClassObject = (FunctionType)GetProcAddress(hDLL, "DllGetClassObject");
        FunctionType DllRegisterServer = (FunctionType)GetProcAddress(hDLL, "DllRegisterServer");
        FunctionType DllUnregisterServer = (FunctionType)GetProcAddress(hDLL, "DllUnregisterServer");

        // 调用每个函数（如果函数指针非空）
        if (DllCanUnloadNow) DllCanUnloadNow();
        if (DllGetClassObject) DllGetClassObject();
        if (DllRegisterServer) DllRegisterServer();
        if (DllUnregisterServer) DllUnregisterServer();

        // 卸载 DLL
        FreeLibrary(hDLL);
        printf("All functions executed successfully.\n");
    }
    else {
        printf("Failed to load DLL.\n");
    }

    return 0;
}