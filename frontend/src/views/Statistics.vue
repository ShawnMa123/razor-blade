<template>
  <div class="statistics">
    <el-row :gutter="20">
      <!-- 基础统计卡片 -->
      <el-col :span="24">
        <el-card class="stats-overview">
          <template #header>
            <span>统计概览</span>
          </template>
          <el-row :gutter="20">
            <el-col :span="6">
              <div class="stat-item">
                <div class="stat-icon usage">
                  <el-icon><DocumentCopy /></el-icon>
                </div>
                <div class="stat-info">
                  <div class="stat-value">{{ statistics?.total_usage || 0 }}</div>
                  <div class="stat-label">总使用次数</div>
                </div>
              </div>
            </el-col>
            <el-col :span="6">
              <div class="stat-item">
                <div class="stat-icon razor">
                  <el-icon><Operation /></el-icon>
                </div>
                <div class="stat-info">
                  <div class="stat-value">{{ statistics?.razor_count || 0 }}</div>
                  <div class="stat-label">剃须刀数量</div>
                </div>
              </div>
            </el-col>
            <el-col :span="6">
              <div class="stat-item">
                <div class="stat-icon blade">
                  <el-icon><Coin /></el-icon>
                </div>
                <div class="stat-info">
                  <div class="stat-value">{{ statistics?.blade_count || 0 }}</div>
                  <div class="stat-label">刀片种类</div>
                </div>
              </div>
            </el-col>
            <el-col :span="6">
              <div class="stat-item">
                <div class="stat-icon rating">
                  <el-icon><Star /></el-icon>
                </div>
                <div class="stat-info">
                  <div class="stat-value">{{ formatRating(statistics?.average_rating) }}</div>
                  <div class="stat-label">平均评分</div>
                </div>
              </div>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="charts-row">
      <!-- 使用频率图表 -->
      <el-col :span="12">
        <el-card class="chart-card">
          <template #header>
            <span>使用频率分析</span>
          </template>
          <div v-loading="loading" class="chart-container">
            <div class="coming-soon">
              <el-icon size="48"><TrendCharts /></el-icon>
              <p>图表功能即将上线</p>
              <p class="tip">将展示每日/每周/每月使用频率趋势</p>
            </div>
          </div>
        </el-card>
      </el-col>

      <!-- 刀片寿命分析 -->
      <el-col :span="12">
        <el-card class="chart-card">
          <template #header>
            <span>刀片寿命分析</span>
          </template>
          <div v-loading="loading" class="chart-container">
            <div class="coming-soon">
              <el-icon size="48"><PieChart /></el-icon>
              <p>图表功能即将上线</p>
              <p class="tip">将展示不同品牌刀片的使用寿命对比</p>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="charts-row">
      <!-- 成本分析 -->
      <el-col :span="12">
        <el-card class="chart-card">
          <template #header>
            <span>成本分析</span>
          </template>
          <div v-loading="loading" class="chart-container">
            <div class="coming-soon">
              <el-icon size="48"><Money /></el-icon>
              <p>图表功能即将上线</p>
              <p class="tip">将展示剃须成本趋势和性价比分析</p>
            </div>
          </div>
        </el-card>
      </el-col>

      <!-- 评分趋势 -->
      <el-col :span="12">
        <el-card class="chart-card">
          <template #header>
            <span>使用感受趋势</span>
          </template>
          <div v-loading="loading" class="chart-container">
            <div class="coming-soon">
              <el-icon size="48"><DataLine /></el-icon>
              <p>图表功能即将上线</p>
              <p class="tip">将展示使用感受评分的变化趋势</p>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 数据详情表格 -->
    <el-row :gutter="20">
      <el-col :span="24">
        <el-card>
          <template #header>
            <span>详细数据</span>
          </template>
          <el-descriptions :column="2" border>
            <el-descriptions-item label="总使用次数">
              {{ statistics?.total_usage || 0 }} 次
            </el-descriptions-item>
            <el-descriptions-item label="平均评分">
              {{ formatRating(statistics?.average_rating) }} 分
            </el-descriptions-item>
            <el-descriptions-item label="剃须刀数量">
              {{ statistics?.razor_count || 0 }} 把
            </el-descriptions-item>
            <el-descriptions-item label="刀片种类">
              {{ statistics?.blade_count || 0 }} 种
            </el-descriptions-item>
          </el-descriptions>

          <div class="tips-section">
            <h4>💡 使用建议</h4>
            <ul>
              <li>建议记录每次剃须的详细感受，有助于找到最适合的剃须刀和刀片组合</li>
              <li>定期查看统计数据，了解使用习惯和偏好</li>
              <li>关注刀片使用次数，及时更换以获得最佳剃须效果</li>
              <li>对比不同品牌的性价比，优化采购决策</li>
            </ul>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'
import { useStatisticsStore } from '@/stores'
import { storeToRefs } from 'pinia'

const statisticsStore = useStatisticsStore()
const { statistics, loading } = storeToRefs(statisticsStore)

const formatRating = (rating?: number) => {
  if (!rating) return '0.0'
  return rating.toFixed(1)
}

onMounted(() => {
  statisticsStore.fetchStatistics()
})
</script>

<style scoped>
.statistics {
  padding: 0;
}

.stats-overview {
  margin-bottom: 20px;
}

.stat-item {
  display: flex;
  align-items: center;
  padding: 20px;
}

.stat-icon {
  width: 60px;
  height: 60px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 16px;
  color: white;
  font-size: 24px;
}

.stat-icon.usage {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.stat-icon.razor {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.stat-icon.blade {
  background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%);
}

.stat-icon.rating {
  background: linear-gradient(135deg, #43e97b 0%, #38f9d7 100%);
}

.stat-info {
  flex: 1;
}

.stat-value {
  font-size: 32px;
  font-weight: bold;
  color: #2c3e50;
  margin-bottom: 4px;
}

.stat-label {
  font-size: 14px;
  color: #7f8c8d;
}

.charts-row {
  margin-bottom: 20px;
}

.chart-card {
  height: 400px;
}

.chart-container {
  height: 320px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.coming-soon {
  text-align: center;
  color: #909399;
}

.coming-soon p {
  margin: 12px 0 4px;
  font-size: 16px;
}

.coming-soon .tip {
  font-size: 12px;
  color: #c0c4cc;
}

.tips-section {
  margin-top: 24px;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 4px;
}

.tips-section h4 {
  margin: 0 0 12px;
  color: #2c3e50;
}

.tips-section ul {
  margin: 0;
  padding-left: 20px;
}

.tips-section li {
  margin-bottom: 8px;
  color: #606266;
  line-height: 1.4;
}
</style>