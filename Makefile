v := v0.0.2
bin:
	go build . && tar -zcvf dt_arm64_$(v).tar.gz ./dt
	GOARCH=amd64 go build . && tar -zcvf dt_amd64_$(v).tar.gz ./dt
	GOOS=windows GOARCH=amd64 go build -o dt.exe . && tar -zcvf dt_win_$(v).tar.gz ./dt.exe
