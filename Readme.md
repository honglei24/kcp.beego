# 开发手册
- 安装mysql
```
# sudo docker pull mysql:5.7
# sudo docker run --name mysql -p 3306:3306 -e MYSQL\_ROOT\_PASSWORD=123456 -d mysql:5.7
```

- 运行代码
```
下载所需依赖包(略)

下载代码到$GOPATH/src/目录
# go run main.go

通过http://127.0.0.1:8080/swagger/可以查看接口文档
```

