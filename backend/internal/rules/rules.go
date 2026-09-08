package rules

import (
	"fmt"
)

// RuleEngine 规则引擎
type RuleEngine struct {
	presets map[string]Preset
	custom  []CustomRule
}

// Preset 预设配置
type Preset struct {
	Name        string
	Description string
	Rules       []CustomRule
}

// CustomRule 自定义规则
type CustomRule struct {
	Name      string
	Type      string // combine, transform, filter
	Config    map[string]interface{}
	Weight    int
}

// New 创建规则引擎
func New() *RuleEngine {
	re := &RuleEngine{
		presets: make(map[string]Preset),
		custom:  make([]CustomRule, 0),
	}
	re.initPresets()
	return re
}

// initPresets 初始化预设
func (re *RuleEngine) initPresets() {
	// 社会工程学预设
	re.presets["social-engineering"] = Preset{
		Name:        "社会工程学",
		Description: "基于个人信息的组合预设",
		Rules: []CustomRule{
			{
				Name:   "name-birth-combine",
				Type:   "combine",
				Weight: 10,
			},
			{
				Name:   "name-city-combine",
				Type:   "combine",
				Weight: 8,
			},
		},
	}

	// 键盘轨迹预设
	re.presets["keyboard-trace"] = Preset{
		Name:        "键盘轨迹",
		Description: "基于键盘相邻字符的预设",
		Rules: []CustomRule{
			{
				Name:   "qwerty-pattern",
				Type:   "combine",
				Weight: 10,
			},
			{
				Name:   "number-row-pattern",
				Type:   "combine",
				Weight: 8,
			},
		},
	}

	// 吉利数字预设
	re.presets["lucky-numbers"] = Preset{
		Name:        "吉利数字",
		Description: "基于吉祥数字的组合预设",
		Rules: []CustomRule{
			{
				Name:   "lucky-number-combine",
				Type:   "combine",
				Weight: 10,
			},
		},
	}
}

// GetPreset 获取预设
func (re *RuleEngine) GetPreset(name string) (Preset, error) {
	if preset, exists := re.presets[name]; exists {
		return preset, nil
	}
	return Preset{}, fmt.Errorf("preset not found: %s", name)
}

// ListPresets 列出所有预设
func (re *RuleEngine) ListPresets() map[string]string {
	result := make(map[string]string)
	for name, preset := range re.presets {
		result[name] = preset.Description
	}
	return result
}

// AddCustomRule 添加自定义规则
func (re *RuleEngine) AddCustomRule(rule CustomRule) {
	re.custom = append(re.custom, rule)
}

// GetCustomRules 获取自定义规则
func (re *RuleEngine) GetCustomRules() []CustomRule {
	return re.custom
}
