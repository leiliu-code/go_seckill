# 秒杀系统（MVP0）

go的秒杀项目，主要目的是为了一步一步学习go，了解go的并发，和go的后端开发，现在目前在MVP0，尚未竣工

## 1. 技术栈

- Go  1.27
- MySQL8.0，库名 `seckill`
- 无需 Docker、无需 Redis（MVP0）

## 2. 目录结构

```text
go_seckill/
├── main.go                  # 程序主入口，注册dsn，启动gin引擎
├── handler.go               # 处理http请求
├── repo.go                  # 调用数据库
├── model.go                 # 结构体模板
└── migrations/01_init.sql   # 数据库初始化文件，加入products和orders两张表及50条商品信息
```

## 3. 初始化数据库

```powershell
# 做什么：建 products / orders 两张表（脚本自带 USE seckill; 和 IF NOT EXISTS，可重复执行）
mysql -u dev -p --default-character-set=utf8mb4 < migrations/01_init.sql
```

## 4. 启动

```powershell
# 做什么：设置数据库密码（唯一需要的环境变量）
$env:DB_PASS = "<你的密码>"

# 做什么：编译并启动，监听 :8080；前台阻塞运行，Ctrl+C 停止
go run .
```

> 数据库用户 `dev`、地址 `127.0.0.1:3306`、库名 `seckill` 目前**写死在** `main.go:14-18`，要换需改代码。

## 5. 接口

### 5.1 商品列表 `GET /api/v1/products`

```powershell
# 做什么：取全部商品
curl.exe -i http://127.0.0.1:8080/api/v1/products
```

### 5.2 商品详情 `GET /api/v1/products/:id`

| 请求 | 期望 |
|---|---|
| `/api/v1/products/2` | 200 |
| `/api/v1/products/999` | 404 |
| `/api/v1/products/abc` | 400 |
| （停掉 MySQL）`/api/v1/products/1` | 500 |
## 6. 当前限界（MVP0）

- 无 Redis 预扣/缓存
- 无限流
- 无异步落库
- 无超时关单
- 无单元测试
- 无 Docker
- 无压测报告