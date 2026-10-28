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
└── migrations/01_init.sql   # 数据库初始化文件，加入products和orders两张表及10条商品信息
```

## 3. 初始化数据库

```powershell
# 做什么：建 products / orders 两张表（脚本自带 USE seckill; 和 IF NOT EXISTS，幂等：建表靠 IF NOT EXISTS；数据靠固定 id + ON DUPLICATE KEY UPDATE，重复执行条数不变）
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
> 说明：所有命令用 `curl.exe`（PowerShell 里 `curl` 是 `Invoke-WebRequest` 的别名，参数不兼容）。
> 下方 JSON 为便于阅读做了缩进，实际响应是单行紧凑 JSON。

### 5.1 商品列表 `GET /api/v1/products`

```powershell
# 做什么：取全部商品
curl.exe -i http://127.0.0.1:8080/api/v1/products
```

预期输出（**共 10 条，此处只截前 2 条示意，非原样输出**）：

```json
{
  "products": [
    { "id": 1, "name": "1", "price": 10000, "stock": 100, "version": 0, 
      "created_at": "2026-10-28T17:05:12+08:00", "updated_at": "2026-10-28T17:05:12+08:00" },
    { "id": 2, "name": "2", "price": 20000, "stock": 100, "version": 0,
      "created_at": "2026-10-28T17:05:12+08:00", "updated_at": "2026-10-28T17:05:12+08:00" }
  ]
}
HTTP 200
```
- `price` 单位是**分**（`price: 10000` = 100 元）
- 顶层键固定为 `products`；库为空时返回 `{"products":null}`

### 5.2 商品详情 `GET /api/v1/products/:id`

```powershell
# 做什么：查 id=1 的商品
curl.exe -s -w "`nHTTP %{http_code}`n" http://127.0.0.1:8080/api/v1/products/1
```

预期输出：

```json
{
  "product": {
    "id": 1, "name": "1", "price": 10000, "stock": 100, "version": 0,
    "created_at": "2026-10-28T17:05:12+08:00", "updated_at": "2026-10-28T17:05:12+08:00"
  }
}
HTTP 200
```

**异常情况**（4 条命令 + 预期，`error` 文案来自代码，可精确断言）：

| 场景 | 命令 | 状态码 | 预期响应体 |
|---|---|---|---|
| id 不存在 | `curl.exe -s -w "`nHTTP %{http_code}`n" http://127.0.0.1:8080/api/v1/products/999` | 404 | `{"error":"商品不存在"}` |
| id 非数字 | `... /api/v1/products/abc` | 400 | `{"error":"id 必须是数字"}` |
| 数据库不可用 | 先停 MySQL80，再 `... /api/v1/products/1` | 500 | `{"error":"<驱动报错原文>"}`（内容随环境变化，断言状态码即可） |

> 说明：500 的 `error` 是数据库驱动返回的原始错误（可能是 `invalid connection` 或 `dial tcp ...: connect: connection refused`），**不做断言**。

## 6. 当前限界（MVP0）

- 无 Redis 预扣/缓存
- 无限流
- 无异步落库
- 无超时关单
- 无单元测试
- 无 Docker
- 无压测报告