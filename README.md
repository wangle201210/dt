# data-transfer
数据转换
## 安装
`go install github.com/wangle201210/dt@latest`

## 功能列表
- [x] string 转 md5 `dt md5 -s wanna` echo `aa4dec62924881f79122e03e2254131a`
- [x] date 转 时间戳 `dt time date -d "2024-07-23 15:50:39"` echo `1721749839`
- [x] 时间戳 转 date `dt time ts -t 1721749839` echo `2024-07-23 15:50:39`