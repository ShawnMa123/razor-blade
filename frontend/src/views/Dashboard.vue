<template>
  <div class="dashboard">
    <!-- 统计卡片 -->
    <el-row :gutter="20" class="stats-row">
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-item">
            <div class="stat-icon usage">
              <el-icon><DocumentCopy /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ dashboardData?.statistics?.total_usage || 0 }}</div>
              <div class="stat-label">总使用次数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-item">
            <div class="stat-icon razor">
              <el-icon><Operation /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ dashboardData?.statistics?.razor_count || 0 }}</div>
              <div class="stat-label">剃须刀数量</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-item">
            <div class="stat-icon blade">
              <el-icon><Coin /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ dashboardData?.statistics?.blade_count || 0 }}</div>
              <div class="stat-label">刀片种类</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="stat-card">
          <div class="stat-item">
            <div class="stat-icon rating">
              <el-icon><Star /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ formatRating(dashboardData?.statistics?.average_rating) }}</div>
              <div class="stat-label">平均评分</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20" class="content-row">
      <!-- 快速记录表单 -->
      <el-col :span="12">
        <el-card class="quick-record-card">
          <template #header>
            <div class="card-header">
              <span>快速记录</span>
            </div>
          </template>
          <QuickRecordForm @record-created="handleRecordCreated" />
        </el-card>
      </el-col>

      <!-- 最近记录 -->
      <el-col :span="12">
        <el-card class="recent-records-card">
          <template #header>
            <div class="card-header">
              <span>最近记录</span>
              <el-button type="text" @click="$router.push('/usage-records')">
                查看全部
              </el-button>
            </div>
          </template>
          <div v-if="loading" class="loading-container">
            <el-skeleton :rows="3" animated />
          </div>
          <div v-else>
            <div v-if="dashboardData?.recent_records?.length === 0" class="empty-state">
              <el-empty description="暂无使用记录" />
            </div>
            <div v-else class="recent-records-list">
              <div
                v-for="record in dashboardData?.recent_records"
                :key="record.id"
                class="record-item"
              >
                <div class="record-header">
                  <span class="record-time">{{ formatDate(record.usage_time) }}</span>
                  <el-rate
                    v-if="record.rating"
                    :model-value="record.rating"
                    disabled
                    size="small"
                  />
                </div>
                <div class="record-content">
                  <span class="record-razor">{{ record.razor.brand }} {{ record.razor.model }}</span>
                  <span class="record-blade">{{ record.blade.brand }} {{ record.blade.model }}</span>
                  <span class="record-count">第{{ record.blade_usage_count }}次</span>
                </div>
                <div v-if="record.experience_text" class="record-experience">
                  {{ record.experience_text }}
                </div>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useStatisticsStore } from '@/stores'
import QuickRecordForm from '@/components/QuickRecordForm.vue'
import dayjs from 'dayjs'

const statisticsStore = useStatisticsStore()

const { dashboardData, loading } = storeToRefs(statisticsStore)

const formatDate = (dateString: string) => {
  return dayjs(dateString).format('MM-DD HH:mm')
}

const formatRating = (rating?: number) => {
  if (!rating) return '0.0'
  return rating.toFixed(1)
}

const handleRecordCreated = () => {
  // 记录创建成功后刷新仪表板数据
  statisticsStore.fetchDashboardData()
}

onMounted(() => {
  statisticsStore.fetchDashboardData()
})
</script>

<style scoped>
.dashboard {
  padding: 0;
}

.stats-row {
  margin-bottom: 20px;
}

.stat-card {
  height: 120px;
}

.stat-item {
  display: flex;
  align-items: center;
  height: 100%;
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
  font-size: 28px;
  font-weight: bold;
  color: #2c3e50;
  margin-bottom: 4px;
}

.stat-label {
  font-size: 14px;
  color: #7f8c8d;
}

.content-row {
  margin-bottom: 20px;
}

.quick-record-card,
.recent-records-card {
  height: 500px;
}

.card-header {
  display: flex;
  justify-content: between;
  align-items: center;
}

.loading-container {
  padding: 20px;
}

.empty-state {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 300px;
}

.recent-records-list {
  max-height: 400px;
  overflow-y: auto;
}

.record-item {
  padding: 16px;
  border-bottom: 1px solid #f0f0f0;
  transition: background-color 0.2s;
}

.record-item:hover {
  background-color: #f9f9f9;
}

.record-item:last-child {
  border-bottom: none;
}

.record-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.record-time {
  font-size: 12px;
  color: #909399;
}

.record-content {
  display: flex;
  gap: 12px;
  margin-bottom: 4px;
}

.record-razor {
  font-weight: bold;
  color: #2c3e50;
}

.record-blade {
  color: #606266;
}

.record-count {
  font-size: 12px;
  color: #909399;
}

.record-experience {
  font-size: 12px;
  color: #909399;
  margin-top: 4px;
  line-height: 1.4;
}
</style>