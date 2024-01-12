# fscan


## 编译
> 前言，由于需要配合CGO，故需要安装gcc
> 本项目在fscan基础上，完成了一些免杀的操作，如有需要请自行编译.
> 使用grable混淆
```powershell
# 编译dll
$env:CGO_ENABLED=1
gradle build -buildmode=c-shared -o fscan.dll ./cmd/fscan_dll/main.go
# 编译 fscan_loader (go)
gradle build -o fscan_loader_go.exe ./cmd/fscan_loader/main.go
# 编译 fscan_loader (c)
gcc -o fscan_loader_c.exe ./cmd/fscan_loader/main.c
```
## 运行
将编译好的dll（fscan.dll和fscan.h）和 fscan_loader.exe放在同一目录下，运行exe即可
```shell
# 运行 fscan_loader_go.exe
fscan_loader_go.exe
```
## 二进制
如果无法编译，可以直接使用编译好的二进制文件
- [fscan_loader_go.exe](./bin/fscan_loader_go.exe)
- [fscan_loader_c.exe](./bin/fscan_loader_c.exe)
- [fscan.dll](./bin/fscan_go.dll)
- [fscan.h](./bin/fscan.h)
## 详情
- 