## word
```shell
$ go run main.go word -h  
该子命令支持各种单词格式的转换,模式如下: 
1: 全部单词转为大写
2: 全部单词转换为小写
3: 下划线单词转换为大写驼峰单词
4: 下划线单词转换为小写驼峰单词
5: 驼峰单词转换为下划线单词

Usage:
   word [flags]

Flags:
  -h, --help         help for word
  -m, --mode int8    请输入单词转换的格式
  -s, --str string   请输入单词
```
## sql
```shell
$ go run main.go sql struct  -h
sql 转换

Usage:
   sql struct [flags]

Flags:
      --charset string    请输入数据库的编码 (default "utf8mb4")
      --db string         请输入数据库名称
  -h, --help              help for struct
      --host string       请输入数据库的HOST (default "127.0.0.1:3306")
      --password string   请输入数据库的密码
      --table string      请输入表名
      --type string       请输入数据库的类型 (default "mysql")
      --username string   请输入数据库的账号

```
example
```shell
$ go run main.go sql struct --username root --password "" --db go_admin_dev --table casbin_rule
type CasbinRule struct {
    // p_type 
    PType string `json: p_type` 
    // v0 
    V0 string `json: v0` 
    // v1 
    V1 string `json: v1` 
    // v2 
    V2 string `json: v2` 
    // v3 
    V3 string `json: v3` 
    // v4 
    V4 string `json: v4` 
    // v5 
    V5 string `json: v5` 
}

```
