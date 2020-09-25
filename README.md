#### what?
> 使用adb，支持飞书和微信的抢红包工具


##### 支持的平台 
> 输入你想抢红包的平台: 1.微信,2:飞书  
> 你可以自己在config.json中 添加 你想获取的颜色，目前位置做了适配，只取中间的颜色，具体可在images文件中查看，也可以在 rekpkg/image.go 中修改 postion，并在 rekpkg/config.go 中 修改代码
>默认值: 1



##### 如何使用

###### 下载 打包好的release

###### 自己编译
* [安装ADB工具](https://github.com/wangshub/wechat_jump_game/wiki/Android-%E5%92%8C-iOS-%E6%93%8D%E4%BD%9C%E6%AD%A5%E9%AA%A4)
* win平台编译 CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o hot-pkg.exe ./main.go
* mac平台编译 CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o hot-pkg.exe ./main.go
* 点击运行 hot-pkg.exe

# 本脚本只做学习使用，请勿作为商业或其他用途，一切使用法律风险与作者无关
