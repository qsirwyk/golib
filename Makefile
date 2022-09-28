PROJECTNAME = go

## 在windows下此Makefile需要加Set 像下面那样 在linux和mac上可以使用现在的命令
## SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build
## https://blog.csdn.net/peng2hui1314/article/details/119936821

## linux: 编译打包linux
.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(RACE)
##-o main ./main.go

## win: 编译打包win
.PHONY: win
win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(RACE)
##-o main.exe ./main.go

## mac: 编译打包mac
.PHONY: mac
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(RACE)
##-o main ./main.go
