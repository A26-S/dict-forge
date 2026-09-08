<template>
  <div class="presets-view">
    <el-row :gutter="20">
      <el-col :span="8" v-for="preset in presets" :key="preset.name">
        <el-card class="preset-card" @click="selectPreset(preset)" :class="{ active: selectedPreset?.name === preset.name }">
          <div class="preset-icon">{{ preset.icon }}</div>
          <div class="preset-name">{{ preset.name }}</div>
          <div class="preset-desc">{{ preset.description }}</div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const presets = ref([
  {
    name: '社会工程学',
    icon: '🎯',
    description: '基于个人信息的组合预设',
  },
  {
    name: '键盘轨迹',
    icon: '⌨️',
    description: '基于键盘相邻字符的预设',
  },
  {
    name: '吉利数字',
    icon: '🎰',
    description: '基于吉祥数字的组合预设',
  },
]);

const selectedPreset = ref(null);

const selectPreset = (preset) => {
  selectedPreset.value = preset;
};
</script>

<style scoped>
.presets-view {
  animation: fadeIn 0.3s ease-in;
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.preset-card {
  text-align: center;
  cursor: pointer;
  transition: all 0.3s ease;
  border: 2px solid #e5e7eb;
}

.preset-card:hover {
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transform: translateY(-4px);
}

.preset-card.active {
  border-color: #3b82f6;
  background: #eff6ff;
}

.preset-icon {
  font-size: 48px;
  margin-bottom: 10px;
}

.preset-name {
  font-size: 18px;
  font-weight: bold;
  color: #1f2937;
  margin-bottom: 5px;
}

.preset-desc {
  font-size: 12px;
  color: #6b7280;
}
</style>
