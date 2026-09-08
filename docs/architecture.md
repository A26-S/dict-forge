# 架构设计

## 系统架构

```
┌─────────────────────────────────────────────────────┐
│          Electron Desktop Application               │
│  (Vue3 + Element Plus + Echarts)                    │
└─────────────────┬───────────────────────────────────┘
                  │ HTTP/IPC
┌─────────────────▼───────────────────────────────────┐
│          Go Backend Service                         │
│  (REST API + WebSocket)                             │
├─────────────────────────────────────────────────────┤
│ ┌──────────┐ ┌────────────┐ ┌───────────────────┐ │
│ │Generator │ │RuleEngine  │ │ Transform/Filter  │ │
│ │          │ │            │ │                   │ │
│ │- Combine │ │- Presets   │ │- Case Transform   │ │
│ │- Cartesian│ │- Custom    │ │- Leet/Number Rplc│ │
│ │- Sample  │ │- Weight    │ │- Pinyin Convert  │ │
│ └──────────┘ └────────────┘ └───────────────────┘ │
├─────────────────────────────────────────────────────┤
│ ┌────────────────────────────────────────────────┐ │
│ │         Storage Layer                          │ │
│ │  LevelDB (Bloom Filter + Cache)                │ │
│ └────────────────────────────────────────────────┘ │
└─────────────────┬───────────────────────────────────┘
                  │
┌─────────────────▼───────────────────────────────────┐
│       Data Layer                                    │
│  ┌─────────────┐ ┌─────────────┐ ┌────────────┐   │
│  │ names.json  │ │ places.json  │ │ presets.json
│  └─────────────┘ └─────────────┘ └────────────┘   │
└─────────────────────────────────────────────────────┘
```

## 模块说明

### 1. 生成引擎（Generator）
- **Cartesian Product Generator** - 笛卡尔积组合
- **Permutation Generator** - 排列组合
- **Sampling Generator** - 智能采样（大数据量时）
- **Stream Generator** - 流式生成（节省内存）

### 2. 规则引擎（Rule Engine）
- **Preset Manager** - 预设管理
- **Custom Rule Builder** - 自定义规则
- **Weight System** - 权重系统
- **Conditional Rules** - 条件规则

### 3. 变形模块（Transform）
- **Case Transform** - 大小写变换
- **Character Replace** - 字符替换（谐音、Leet等）
- **Pinyin Convert** - 拼音转换
- **Traditional-Simplified** - 繁简转换

### 4. 过滤模块（Filter）
- **Length Filter** - 长度过滤
- **Character Type Filter** - 字符类型过滤
- **Entropy Filter** - 熵值过滤
- **Weakness Filter** - 弱点过滤
- **Regex Filter** - 正则表达式过滤

### 5. 存储层（Storage）
- **Bloom Filter** - 快速去重
- **LevelDB** - 本地数据库
- **Memory Cache** - 内存缓存
- **Export Manager** - 导出管理

## 性能优化策略

1. **多线程并行生成**
   - 使用 goroutines 进行并行生成
   - 工作队列模式
   - 动态线程池

2. **高效去重**
   - 布隆过滤器（O(1)查询）
   - 哈希集合
   - 增量式去重

3. **内存管理**
   - 流式处理
   - 分块处理
   - 定期垃圾回收

4. **缓存策略**
   - LRU缓存
   - 热数据预加载
   - 本地持久化

## 数据流

```
用户输入 → 种子数据选择 → 规则组合 → 生成密码 → 变形处理 → 过滤 → 去重 → 统计 → 导出
   ↓           ↓            ↓         ↓         ↓      ↓    ↓     ↓     ↓
 UI输入   从数据库读取   规则引擎   生成器   变形器   过滤器 Bloom  汇总   文件
```

## 扩展性设计

1. **插件系统** - 支持自定义变形规则
2. **规则库** - 可导入导出规则配置
3. **字典库** - 支持自定义词库
4. **云端同步** - 规则和历史记录同步
5. **集成接口** - 与安全工具集成

