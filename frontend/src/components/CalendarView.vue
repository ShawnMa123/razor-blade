<template>
  <div class="calendar-view">
    <el-card>
      <template #header>
        <div class="calendar-header">
          <span>剃须记录日历</span>
          <div class="calendar-controls">
            <el-button @click="prevMonth" :icon="ArrowLeft" circle />
            <span class="current-month">{{ currentMonthLabel }}</span>
            <el-button @click="nextMonth" :icon="ArrowRight" circle />
          </div>
        </div>
      </template>

      <div class="calendar-body">
        <div class="calendar-weekdays">
          <div v-for="weekday in weekdays" :key="weekday" class="weekday">
            {{ weekday }}
          </div>
        </div>

        <div class="calendar-grid">
          <div
            v-for="date in calendarDates"
            :key="date.key"
            :class="[
              'calendar-date',
              {
                'other-month': !date.isCurrentMonth,
                'today': date.isToday,
                'has-record': date.hasRecord
              }
            ]"
          >
            <div class="date-number">{{ date.day }}</div>
            <div v-if="date.records.length > 0" class="records-indicator">
              <div class="record-count">{{ date.records.length }}</div>
              <div class="record-preview">
                <div
                  v-for="record in date.records.slice(0, 2)"
                  :key="record.id"
                  class="record-item"
                  :title="`${record.razor.brand} ${record.razor.model} - ${record.blade.brand} ${record.blade.model}`"
                >
                  <el-rate
                    v-if="record.rating"
                    :model-value="record.rating"
                    disabled
                    size="small"
                    show-score
                    text-color="#909399"
                    score-template="{value}"
                  />
                  <span v-else class="no-rating">-</span>
                </div>
                <div v-if="date.records.length > 2" class="more-records">
                  +{{ date.records.length - 2 }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { ArrowLeft, ArrowRight } from '@element-plus/icons-vue'
import { useUsageRecordStore } from '@/stores'
import { storeToRefs } from 'pinia'
import dayjs from 'dayjs'
import type { UsageRecord } from '@/types'

const usageRecordStore = useUsageRecordStore()
const { usageRecords } = storeToRefs(usageRecordStore)

const currentDate = ref(dayjs())
const weekdays = ['日', '一', '二', '三', '四', '五', '六']

const currentMonthLabel = computed(() => {
  return currentDate.value.format('YYYY年MM月')
})

const prevMonth = () => {
  currentDate.value = currentDate.value.subtract(1, 'month')
}

const nextMonth = () => {
  currentDate.value = currentDate.value.add(1, 'month')
}

// 生成日历日期数据
const calendarDates = computed(() => {
  const startOfMonth = currentDate.value.startOf('month')
  const endOfMonth = currentDate.value.endOf('month')
  const startOfCalendar = startOfMonth.startOf('week')
  const endOfCalendar = endOfMonth.endOf('week')

  const dates = []
  let current = startOfCalendar

  while (current.isBefore(endOfCalendar) || current.isSame(endOfCalendar)) {
    const dateStr = current.format('YYYY-MM-DD')
    const dateRecords = usageRecords.value.filter((record: UsageRecord) =>
      dayjs(record.usage_time).format('YYYY-MM-DD') === dateStr
    )

    dates.push({
      key: dateStr,
      day: current.date(),
      date: current,
      isCurrentMonth: current.month() === currentDate.value.month(),
      isToday: current.isSame(dayjs(), 'day'),
      hasRecord: dateRecords.length > 0,
      records: dateRecords
    })

    current = current.add(1, 'day')
  }

  return dates
})

onMounted(async () => {
  // 加载更多使用记录以便在日历上显示
  await usageRecordStore.fetchUsageRecords({ page: 1, page_size: 100 })
})
</script>

<style scoped>
.calendar-view {
  width: 100%;
}

.calendar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.calendar-controls {
  display: flex;
  align-items: center;
  gap: 16px;
}

.current-month {
  font-weight: 600;
  font-size: 16px;
  color: #2c3e50;
  min-width: 120px;
  text-align: center;
}

.calendar-body {
  margin-top: 16px;
}

.calendar-weekdays {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 1px;
  margin-bottom: 8px;
}

.weekday {
  padding: 12px;
  text-align: center;
  font-weight: 600;
  color: #606266;
  background-color: #f5f7fa;
  border-radius: 4px;
}

.calendar-grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 1px;
  background-color: #ddd;
  border-radius: 8px;
  overflow: hidden;
}

.calendar-date {
  min-height: 100px;
  padding: 8px;
  background-color: white;
  position: relative;
  transition: background-color 0.2s;
}

.calendar-date:hover {
  background-color: #f8f9fa;
}

.calendar-date.other-month {
  background-color: #fafafa;
  color: #c0c4cc;
}

.calendar-date.today {
  background-color: #e3f2fd;
  border: 2px solid #2196f3;
}

.calendar-date.has-record {
  background-color: #fff3e0;
}

.calendar-date.has-record.today {
  background-color: #e8f5e8;
  border-color: #4caf50;
}

.date-number {
  font-weight: 600;
  font-size: 14px;
  margin-bottom: 4px;
}

.records-indicator {
  font-size: 12px;
}

.record-count {
  background-color: #ff9800;
  color: white;
  border-radius: 50%;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 10px;
  font-weight: bold;
  margin-bottom: 4px;
}

.record-preview {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.record-item {
  display: flex;
  align-items: center;
  font-size: 10px;
}

.record-item :deep(.el-rate) {
  display: flex;
  align-items: center;
}

.record-item :deep(.el-rate__text) {
  font-size: 10px;
  margin-left: 2px;
}

.no-rating {
  color: #909399;
  font-size: 10px;
}

.more-records {
  color: #909399;
  font-size: 10px;
  text-align: center;
  margin-top: 2px;
}

@media (max-width: 768px) {
  .calendar-date {
    min-height: 80px;
    padding: 4px;
  }

  .date-number {
    font-size: 12px;
  }

  .record-count {
    width: 16px;
    height: 16px;
    font-size: 8px;
  }
}
</style>