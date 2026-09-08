package generator

import (
	"sync"
)

// Generator 核心密码字典生成器
type Generator struct {
	seeds      []string
	rules      []Rule
	filters    []Filter
	transforms []Transform
	mu         sync.RWMutex
	results    map[string]bool // 用于去重
}

// Rule 组合规则定义
type Rule struct {
	Name     string
	Pattern  string
	Weight   int
}

// Filter 过滤条件
type Filter struct {
	Name  string
	Check func(string) bool
}

// Transform 变形规则
type Transform struct {
	Name      string
	Transform func(string) []string
}

// New 创建新的生成器实例
func New() *Generator {
	return &Generator{
		seeds:   make([]string, 0),
		rules:   make([]Rule, 0),
		filters: make([]Filter, 0),
		results: make(map[string]bool),
	}
}

// AddSeed 添加种子数据
func (g *Generator) AddSeed(seed string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.seeds = append(g.seeds, seed)
}

// AddRule 添加组合规则
func (g *Generator) AddRule(rule Rule) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.rules = append(g.rules, rule)
}

// Generate 生成密码字典
func (g *Generator) Generate() []string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	var results []string
	for password := range g.results {
		results = append(results, password)
	}
	return results
}

// GetStats 获取统计信息
func (g *Generator) GetStats() map[string]int {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return map[string]int{
		"total":    len(g.results),
		"seeds":    len(g.seeds),
		"rules":    len(g.rules),
		"filters":  len(g.filters),
	}
}
