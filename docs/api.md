# API 文档

## 后端 REST API

### 基础信息
- **基础URL**: `http://localhost:8080/api`
- **内容类型**: `application/json`

### 端点列表

#### 1. 生成密码字典

**请求**
```
POST /api/generate
Content-Type: application/json

{
  "seeds": {
    "names": ["张三", "李四"],
    "places": ["北京", "上海"],
    "years": ["2020", "2021"],
    "lucky_numbers": ["8", "6", "88"]
  },
  "transforms": {
    "capitalize": true,
    "uppercase": true,
    "leet": true,
    "number_replace": true
  },
  "filters": {
    "min_length": 6,
    "max_length": 32,
    "require_digit": false,
    "require_upper": false
  },
  "options": {
    "preset": "social-engineering",
    "sample_rate": 1.0,
    "max_results": 1000000
  }
}
```

**响应** (WebSocket 流式)
```json
{
  "event": "progress",
  "data": {
    "progress": 25,
    "current": 250000,
    "total_estimated": 1000000,
    "elapsed_ms": 5000,
    "eta_ms": 15000
  }
}

{
  "event": "sample",
  "data": [
    "password1",
    "password2",
    "password3"
  ]
}

{
  "event": "complete",
  "data": {
    "total": 950000,
    "deduplicated": 900000,
    "filtered": 850000,
    "duration_ms": 20000
  }
}
```

#### 2. 列出预设

**请求**
```
GET /api/presets
```

**响应**
```json
{
  "presets": {
    "social-engineering": {
      "name": "社会工程学",
      "description": "基于个人信息的组合"
    },
    "keyboard-trace": {
      "name": "键盘轨迹",
      "description": "基于键盘相邻字符"
    },
    "lucky-numbers": {
      "name": "吉利数字",
      "description": "基于吉祥数字组合"
    }
  }
}
```

#### 3. 获取预设详情

**请求**
```
GET /api/presets/{preset_name}
```

**响应**
```json
{
  "name": "社会工程学",
  "description": "基于个人信息的组合预设",
  "rules": [
    {
      "name": "name-birth",
      "type": "combine",
      "weight": 10
    },
    {
      "name": "name-city",
      "type": "combine",
      "weight": 8
    }
  ]
}
```

#### 4. 获取种子数据

**请求**
```
GET /api/seeds/{type}
```

参数:
- `type`: `names`, `places`, `years`, `lucky_numbers`, `keyboard_patterns`

**响应**
```json
{
  "type": "names",
  "data": [
    "张",
    "王",
    "李",
    ...
  ]
}
```

#### 5. 计算密码熵值

**请求**
```
POST /api/analyze/entropy
Content-Type: application/json

{
  "passwords": ["password123", "P@ssw0rd", "test123"]
}
```

**响应**
```json
{
  "results": [
    {
      "password": "password123",
      "entropy": 42.5,
      "strength": "中",
      "charset_size": 62
    },
    {
      "password": "P@ssw0rd",
      "entropy": 58.3,
      "strength": "强",
      "charset_size": 94
    }
  ]
}
```

#### 6. 导出结果

**请求**
```
POST /api/export
Content-Type: application/json

{
  "format": "txt",
  "task_id": "abc-123",
  "sort_by": "entropy"
}
```

**响应** (文件流)
```
password1
password2
password3
...
```

#### 7. 获取统计信息

**请求**
```
GET /api/stats/{task_id}
```

**响应**
```json
{
  "task_id": "abc-123",
  "status": "completed",
  "total": 950000,
  "deduplicated": 900000,
  "filtered": 850000,
  "duration_ms": 20000,
  "memory_used_mb": 256,
  "entropy_distribution": {
    "weak": 100000,
    "medium": 500000,
    "strong": 250000
  }
}
```

## 错误响应

```json
{
  "error": "Invalid parameters",
  "code": 400,
  "details": "Seeds cannot be empty"
}
```

### 常见错误码

- `400` - 请求参数错误
- `404` - 资源不存在
- `500` - 服务器错误
- `503` - 服务不可用

