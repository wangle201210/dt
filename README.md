# data-transfer
数据转换
## 安装
`go install github.com/wangle201210/dt@latest`

## 功能列表
- [x] string 转 md5 `dt md5 -s wanna` echo `aa4dec62924881f79122e03e2254131a`

- [x] date 转 时间戳 `dt time date -d "2024-07-23 15:50:39"` echo `1721749839`
- [x] 时间戳 转 date `dt time ts -t 1721749839` echo `2024-07-23 15:50:39`

- [x] 字符串 转 base64 `dt base64 -e wanna` echo `d2FubmE=`
- [x] base64 转 字符串 `dt base64 -d d2FubmE=` echo `wanna`
- [x] base64 转 图片 `dt base64 -i "data:image/png;base64,***"` echo `temp_image_**.png` 可以不加前缀 `data:image/png;base64,`
- [x] 图片 转 base64 `dt base64 -p temp_image_**.png` echo `编码后的图片数据`
> 组合命令 `dt md5 -s wanna | xargs dt base64 -e` echo `YWE0ZGVjNjI5MjQ4ODFmNzkxMjJlMDNlMjI1NDEzMWE=`
