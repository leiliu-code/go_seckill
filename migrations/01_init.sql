-- 01_init.sql ｜ 作用：建 products 与 orders 两张表 ｜ 执行：mysql -u dev -p -D seckill < migrations/01_init.sql ｜ 幂等：带 IF NOT EXISTS，可重复执行
USE seckill;                                                                           -- 切库：单独执行本脚本时也不会建到别的数据库

CREATE TABLE IF NOT EXISTS `products` (                                                -- 商品表：一行 = 一个可售商品；库存靠 version 字段做乐观锁扣减
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,                                          -- 主键，自增，无业务含义（Go 里是 uint64）
  name VARCHAR(128) NOT NULL,                                                          -- 商品名称，最长 128 个字符
  price BIGINT NOT NULL DEFAULT 0,                                                     -- 单价，单位：分（1 元 = 100 分），不可为负
  stock INT NOT NULL DEFAULT 0,                                                        -- 剩余库存，单位：件；扣减必须带 stock >= 购买件数
  version INT NOT NULL DEFAULT 0,                                                      -- 乐观锁版本号，每次扣减库存后 +1
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,                              -- 创建时间，由数据库自动填充
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- 行变更时自动刷新时间
  PRIMARY KEY (`id`),                                                                  -- 主键单独声明，将来改成联合主键更方便
  KEY `idx_name` (`name`)                                                              -- 按名称搜索商品时避免全表扫描
) ENGINE=InnoDB 
  DEFAULT CHARSET=utf8mb4 
  COLLATE=utf8mb4_0900_ai_ci 
  COMMENT='商品表；一行=一个可售商品；库存由 version 做乐观锁更新，不做物理删除';

CREATE TABLE IF NOT EXISTS `orders` (                                                  -- 订单表：一行 = 一次下单行为；不物理删除，取消只把 status 改成 2
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,                                          -- 主键，自增，无业务含义
  user_id BIGINT UNSIGNED NOT NULL,                                                    -- 下单用户标识，对应 users.id；MVP0 由调用方传入，不做外键校验
  product_id BIGINT UNSIGNED NOT NULL,                                                 -- 购买商品标识，对应 products.id；不做外键校验
  quantity INT NOT NULL DEFAULT 1,                                                     -- 购买件数，单位：件，必须 >= 1
  amount BIGINT NOT NULL DEFAULT 0,                                                    -- 订单总金额，单位：分（= 商品单价 × 购买件数）
  status TINYINT NOT NULL DEFAULT 0,                                                   -- 订单状态：0待支付 1已支付 2已取消
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,                              -- 下单时间，由数据库自动填充
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- 状态变更时自动刷新
  PRIMARY KEY (`id`),                                                                  -- 主键
  UNIQUE KEY `uk_user_product` (`user_id`,`product_id`),                               -- 同一用户对同一商品唯一，防重复下单（幂等）
  KEY `idx_created_at` (`created_at`)                                                  -- 按时间倒序查订单 / 扫描超时未支付单时用
) ENGINE=InnoDB 
  DEFAULT CHARSET=utf8mb4 
  COLLATE=utf8mb4_0900_ai_ci 
  COMMENT='订单表；一行=一次下单行为；不物理删除，取消只把 status 改为 2';
