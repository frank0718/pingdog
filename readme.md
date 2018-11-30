pingdog是多个vps互相ping，监控某vps死机的进程

# 特点
1 无主模式
2 报警发邮件只依赖本地的postfix 25端口

# 配置文件格式

```
{
	"mailto": "aaa@example.com",
	"servers": {
		"8.8.8.8": "google",
		"114.114.114.114": "114",
		"1.1.1.1": "cloudflare"
	}
}
```
server down的时候发送邮件。
servers 是互相监听的机器。
# copy一份自己的配置保存为config.json

# build
make  生成bin

# run 
ping 逻辑来自https://github.com/sparrc/go-ping#note-on-linux-support
需要先 sysctl -w net.ipv4.ping_group_range="0   2147483647"

然后./pingdog
