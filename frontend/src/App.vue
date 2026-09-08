<template>
  <div id="app" class="app-container">
    <el-container>
      <el-aside width="250px" class="sidebar">
        <div class="logo">🔐 Dictionary Forge</div>
        <el-menu>
          <el-menu-item index="1" @click="currentTab = 'generator'">
            <span>📝 字典生成</span>
          </el-menu-item>
          <el-menu-item index="2" @click="currentTab = 'rules'">
            <span>⚙️ 规则配置</span>
          </el-menu-item>
          <el-menu-item index="3" @click="currentTab = 'presets'">
            <span>🎯 预设方案</span>
          </el-menu-item>
          <el-menu-item index="4" @click="currentTab = 'analysis'">
            <span>📊 分析统计</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-container>
        <el-header class="header">Dictionary Forge - 高性能密码字典生成工具</el-header>
        <el-main class="main-content">
          <GeneratorView v-if="currentTab === 'generator'" />
          <RulesView v-else-if="currentTab === 'rules'" />
          <PresetsView v-else-if="currentTab === 'presets'" />
          <AnalysisView v-else-if="currentTab === 'analysis'" />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import GeneratorView from './views/GeneratorView.vue';
import RulesView from './views/RulesView.vue';
import PresetsView from './views/PresetsView.vue';
import AnalysisView from './views/AnalysisView.vue';

const currentTab = ref('generator');
</script>

<style>
* {
  margin: 0;
  padding: 0;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
}

.app-container {
  height: 100vh;
  background: #f5f5f5;
}

.sidebar {
  background: #1f2937;
  color: white;
  padding: 20px;
}

.logo {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 20px;
  padding: 10px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 4px;
  text-align: center;
}

.header {
  background: #ffffff;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  align-items: center;
  padding: 0 20px;
}

.main-content {
  padding: 20px;
  background: #f5f5f5;
  overflow-y: auto;
}

.el-menu {
  border: none;
  background: transparent;
}

.el-menu-item {
  background: transparent !important;
  color: #d1d5db;
}

.el-menu-item:hover {
  color: white;
  background: rgba(255, 255, 255, 0.1) !important;
}

.el-menu-item.is-active {
  background: #3b82f6 !important;
  color: white;
}
</style>
