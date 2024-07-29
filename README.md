# data-transfer
数据转换
## 安装
1. 如果有golang环境：`go install github.com/wangle201210/dt@latest`  
2. 如果有home brew：`brew tap wangle201210/dt` & `brew install dt`  
3. 在[releases](https://github.com/wangle201210/dt/releases/tag/v0.0.1-alpha)下载可执行文件

## 功能列表
1. md5
- [x] string 转 md5 `dt md5 wanna` echo `aa4dec62924881f79122e03e2254131a`
2. time
- [x] date 转 时间戳 `dt time date -d "2024-07-23 15:50:39"` echo `1721749839`
- [x] 时间戳 转 date `dt time ts -t 1721749839` echo `2024-07-23 15:50:39`
3. base64
- [x] 字符串 转 base64 `dt base64 -e wanna` echo `d2FubmE=`
- [x] base64 转 字符串 `dt base64 -d d2FubmE=` echo `wanna`
- [x] base64 转 图片 `dt base64 -i "data:image/png;base64,***"` echo `temp_image_**.png` 可以不加前缀 `data:image/png;base64,`
- [x] 图片 转 base64 `dt base64 -p temp_image_**.png` echo `编码后的图片数据`
4. url
- [x] 字符串编码为query `dt url -e "https://www.baidu.com?name=wanna&email=a@b.c"` echo `https%3A%2F%2Fwww.baidu.com%3Fname%3Dwanna%26email%3Da%40b.c`
- [x] query解码为字符串 `dt url -d https%3A%2F%2Fwww.baidu.com%3Fname%3Dwanna%26email%3Da%40b.c` echo `https://www.baidu.com?name=wanna&email=a@b.c`
5. rand
- [x] 随机生成字符串 `dt rand -l 10` echo `f9UyjsHd1Q`
- [x] 随机生成长度为N的数字 `dt rand -n -l 10` echo `9071824719`
6. json
- [x] toml转json `dt json toml demo.toml` echo `... ...`
- [x] json转struct `dt json struct '{"example":{"from":{"json":true}}}'` echo `... ...` 可以 -f 传文件名

> 组合命令 `dt rand -l 10 | xargs dt md5 | xargs dt base64 -e` echo `YWU1ZTRjOGUzZDAxYjIzM2E2OTlkYzk3OTFmYTA3ZjE=`
