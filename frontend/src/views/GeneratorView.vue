<template>
  <div class="generator-view">
    <el-card class="seed-panel">
      <template #header>
        <div class="card-header">
          <span>🌱 种子数据</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="姓名">
            <el-input v-model="seeds.names" placeholder="输入姓名，用逗号分隔" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="地名">
            <el-input v-model="seeds.places" placeholder="输入地名，用逗号分隔" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="年份">
            <el-input v-model="seeds.years" placeholder="输入年份，用逗号分隔" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="吉祥数字">
            <el-input v-model="seeds.luckyNumbers" placeholder="输入数字，用逗号分隔" />
          </el-form-item>
        </el-col>
      </el-row>
    </el-card>

    <el-card class="transform-panel" style="margin-top: 20px">
      <template #header>
        <div class="card-header">
          <span>🔧 变形选项</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="6">
          <el-checkbox v-model="transforms.capitalize">首字母大写</el-checkbox>
        </el-col>
        <el-col :span="6">
          <el-checkbox v-model="transforms.uppercase">全大写</el-checkbox>
        </el-col>
        <el-col :span="6">
          <el-checkbox v-model="transforms.lowercase">全小写</el-checkbox>
        </el-col>
        <el-col :span="6">
          <el-checkbox v-model="transforms.leet">Leet变形</el-checkbox>
        </el-col>
        <el-col :span="6" style="margin-top: 10px">
          <el-checkbox v-model="transforms.reverse">字符反转</el-checkbox>
        </el-col>
        <el-col :span="6" style="margin-top: 10px">
          <el-checkbox v-model="transforms.numberReplace">数字替换</el-checkbox>
        </el-col>
        <el-col :span="6" style="margin-top: 10px">
          <el-checkbox v-model="transforms.mixedCase">混合大小写</el-checkbox>
        </el-col>
        <el-col :span="6" style="margin-top: 10px">
          <el-checkbox v-model="transforms.homophone">谐音替换</el-checkbox>
        </el-col>
      </el-row>
    </el-card>

    <el-card class="filter-panel" style="margin-top: 20px">
      <template #header>
        <div class="card-header">
          <span>🔍 过滤条件</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="6">
          <el-form-item label="最小长度">
            <el-input-number v-model="filter.minLength" :min="1" :max="32" />
          </el-form-item>
        </el-col>
        <el-col :span="6">
          <el-form-item label="最大长度">
            <el-input-number v-model="filter.maxLength" :min="1" :max="128" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-checkbox v-model="filter.requireDigit">必含数字</el-checkbox>
          <el-checkbox v-model="filter.requireUpper">必含大写</el-checkbox>
          <el-checkbox v-model="filter.requireSymbol">必含特殊字符</el-checkbox>
        </el-col>
      </el-row>
    </el-card>

    <el-card class="action-panel" style="margin-top: 20px">
      <el-row :gutter="20">
        <el-col :span="6">
          <el-button type="primary" size="large" @click="generate" :loading="generating">
            🚀 开始生成
          </el-button>
        </el-col>
        <el-col :span="6">
          <el-button @click="reset">🔄 重置</el-button>
        </el-col>
        <el-col :span="6">
          <el-button @click="exportDict">💾 导出</el-button>
        </el-col>
        <el-col :span="6">
          <el-button @click="copyResults">📋 复制结果</el-button>
        </el-col>
      </el-row>
    </el-card>

    <el-card class="stats-panel" style="margin-top: 20px">
      <template #header>
        <div class="card-header">
          <span>📊 统计信息</span>
        </div>
      </template>
      <el-row :gutter="20">
        <el-col :span="6">
          <div class="stat-item">
            <div class="stat-label">总条目数</div>
            <div class="stat-value">{{ stats.total }}</div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="stat-item">
            <div class="stat-label">去重后</div>
            <div class="stat-value">{{ stats.deduplicated }}</div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="stat-item">
            <div class="stat-label">通过过滤</div>
            <div class="stat-value">{{ stats.filtered }}</div>
          </div>
        </el-col>
        <el-col :span="6">
          <div class="stat-item">
            <div class="stat-label">生成耗时</div>
            <div class="stat-value">{{ stats.duration }}ms</div>
          </div>
        </el-col>
      </el-row>
      <el-progress :percentage="progress" style="margin-top: 20px" />
    </el-card>

    <el-card class="preview-panel" style="margin-top: 20px">
      <template #header>
        <div class="card-header">
          <span>👁️ 预览结果 (前100条)</span>
        </div>
      </template>
      <el-table :data="previewData" style="width: 100%" max-height="300px">
        <el-table-column prop="password" label="密码" />
        <el-table-column prop="entropy" label="熵值" width="80" />
        <el-table-column prop="strength" label="强度" width="80" />
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { ElMessage, ElMessageBox } from 'element-plus';

const seeds = reactive({
  names: '',
  places: '',
  years: '',
  luckyNumbers: '8,6,88,66,888,6666',
});

const transforms = reactive({
  capitalize: true,
  uppercase: true,
  lowercase: true,
  leet: true,
  reverse: false,
  numberReplace: true,
  mixedCase: false,
  homophone: true,
});

const filter = reactive({
  minLength: 6,
  maxLength: 32,
  requireDigit: false,
  requireUpper: false,
  requireSymbol: false,
});

const stats = reactive({
  total: 0,
  deduplicated: 0,
  filtered: 0,
  duration: 0,
});

const generating = ref(false);
const progress = ref(0);
const previewData = ref([]);

const generate = async () => {
  generating.value = true;
  progress.value = 0;
  
  try {
    // 模拟生成过程
    for (let i = 0; i <= 100; i += 10) {
      progress.value = i;
      await new Promise(resolve => setTimeout(resolve, 100));
    }
    
    stats.total = 1000;
    stats.deduplicated = 950;
    stats.filtered = 850;
    stats.duration = 250;
    
    previewData.value = Array.from({ length: 10 }, (_, i) => ({
      password: `password${i}`,
      entropy: (50 + Math.random() * 50).toFixed(2),
      strength: ['弱', '中', '强'][Math.floor(Math.random() * 3)],
    }));
    
    ElMessage.success('生成完成！');
  } catch (error) {
    ElMessage.error('生成失败：' + error.message);
  } finally {
    generating.value = false;
    progress.value = 100;
  }
};

const reset = () => {
  seeds.names = '';
  seeds.places = '';
  seeds.years = '';
  stats.total = 0;
  progress.value = 0;
  previewData.value = [];
};

const exportDict = () => {
  ElMessage.info('导出功能开发中...');
};

const copyResults = () => {
  ElMessage.info('复制功能开发中...');
};
</script>

<style scoped>
.generator-view {
  animation: fadeIn 0.3s ease-in;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.card-header {
  font-size: 16px;
  font-weight: 600;
  color: #1f2937;
}

.stat-item {
  text-align: center;
  padding: 15px;
  background: #f0f9ff;
  border-radius: 8px;
}

.stat-label {
  color: #6b7280;
  font-size: 12px;
  margin-bottom: 5px;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #3b82f6;
}
</style>
