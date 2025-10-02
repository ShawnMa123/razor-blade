<template>
  <div class="usage-records">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>使用记录管理</span>
          <el-button type="primary" @click="showCreateDialog = true">
            <el-icon><Plus /></el-icon>
            新增记录
          </el-button>
        </div>
      </template>

      <!-- 搜索过滤器 -->
      <div class="filter-section">
        <el-row :gutter="20">
          <el-col :span="6">
            <el-select v-model="filters.razor_id" placeholder="选择剃须刀" clearable @change="handleFilterChange">
              <el-option
                v-for="razor in razors"
                :key="razor.id"
                :label="`${razor.brand} ${razor.model}`"
                :value="razor.id"
              />
            </el-select>
          </el-col>
          <el-col :span="6">
            <el-select v-model="filters.blade_id" placeholder="选择刀片" clearable @change="handleFilterChange">
              <el-option
                v-for="blade in blades"
                :key="blade.id"
                :label="`${blade.brand} ${blade.model}`"
                :value="blade.id"
              />
            </el-select>
          </el-col>
          <el-col :span="6">
            <el-date-picker
              v-model="filters.date_range"
              type="daterange"
              range-separator="至"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
              format="YYYY-MM-DD"
              value-format="YYYY-MM-DD"
              @change="handleFilterChange"
            />
          </el-col>
          <el-col :span="6">
            <el-button @click="resetFilters">重置</el-button>
          </el-col>
        </el-row>
      </div>

      <!-- 数据表格 -->
      <el-table
        v-loading="loading"
        :data="usageRecords"
        style="width: 100%"
        height="400"
      >
        <el-table-column prop="usage_time" label="使用时间" width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.usage_time) }}
          </template>
        </el-table-column>
        <el-table-column label="剃须刀" width="200">
          <template #default="{ row }">
            {{ row.razor.brand }} {{ row.razor.model }}
          </template>
        </el-table-column>
        <el-table-column label="刀片" width="200">
          <template #default="{ row }">
            {{ row.blade.brand }} {{ row.blade.model }}
          </template>
        </el-table-column>
        <el-table-column prop="blade_usage_count" label="使用次数" width="100" />
        <el-table-column label="评分" width="120">
          <template #default="{ row }">
            <el-rate
              v-if="row.rating"
              :model-value="row.rating"
              disabled
              size="small"
            />
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column prop="experience_text" label="使用感受" min-width="200">
          <template #default="{ row }">
            <span v-if="row.experience_text" class="experience-text">
              {{ row.experience_text }}
            </span>
            <span v-else class="no-experience">-</span>
          </template>
        </el-table-column>
        <el-table-column label="需要换片" width="100">
          <template #default="{ row }">
            <el-tag v-if="row.need_blade_change" type="warning" size="small">
              是
            </el-tag>
            <span v-else>-</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="150" fixed="right">
          <template #default="{ row }">
            <el-button
              type="primary"
              size="small"
              @click="handleEdit(row)"
            >
              编辑
            </el-button>
            <el-button
              type="danger"
              size="small"
              @click="handleDelete(row)"
            >
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination-container">
        <el-pagination
          v-model:current-page="pagination.page"
          v-model:page-size="pagination.page_size"
          :page-sizes="[10, 20, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <!-- 创建/编辑对话框 -->
    <UsageRecordDialog
      v-model="showCreateDialog"
      :record="currentRecord"
      @success="handleDialogSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUsageRecordStore, useRazorStore, useBladeStore } from '@/stores'
import { storeToRefs } from 'pinia'
import UsageRecordDialog from '@/components/UsageRecordDialog.vue'
import dayjs from 'dayjs'
import type { UsageRecord } from '@/types'

const usageRecordStore = useUsageRecordStore()
const razorStore = useRazorStore()
const bladeStore = useBladeStore()

const { usageRecords, loading, total } = storeToRefs(usageRecordStore)
const { razors } = storeToRefs(razorStore)
const { blades } = storeToRefs(bladeStore)

const showCreateDialog = ref(false)
const currentRecord = ref<UsageRecord | null>(null)

const pagination = reactive({
  page: 1,
  page_size: 20
})

const filters = reactive({
  razor_id: undefined as number | undefined,
  blade_id: undefined as number | undefined,
  date_range: null as string[] | null
})

const formatDateTime = (dateString: string) => {
  return dayjs(dateString).format('YYYY-MM-DD HH:mm')
}

const handleFilterChange = () => {
  pagination.page = 1
  fetchData()
}

const resetFilters = () => {
  filters.razor_id = undefined
  filters.blade_id = undefined
  filters.date_range = null
  pagination.page = 1
  fetchData()
}

const handleSizeChange = (newSize: number) => {
  pagination.page_size = newSize
  fetchData()
}

const handleCurrentChange = (newPage: number) => {
  pagination.page = newPage
  fetchData()
}

const handleEdit = (record: UsageRecord) => {
  currentRecord.value = record
  showCreateDialog.value = true
}

const handleDelete = async (record: UsageRecord) => {
  try {
    await ElMessageBox.confirm(
      '确定要删除这条使用记录吗？',
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )

    await usageRecordStore.deleteUsageRecord(record.id)
    ElMessage.success('删除成功')
    fetchData()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

const handleDialogSuccess = () => {
  showCreateDialog.value = false
  currentRecord.value = null
  fetchData()
}

const fetchData = () => {
  const params = {
    page: pagination.page,
    page_size: pagination.page_size,
    ...filters
  }

  // 移除空值
  Object.keys(params).forEach(key => {
    if (params[key] === undefined || params[key] === null) {
      delete params[key]
    }
  })

  usageRecordStore.fetchUsageRecords(params)
}

onMounted(() => {
  fetchData()
  razorStore.fetchRazors({ page: 1, page_size: 100 })
  bladeStore.fetchBlades({ page: 1, page_size: 100 })
})
</script>

<style scoped>
.usage-records {
  padding: 0;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.filter-section {
  margin-bottom: 20px;
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 4px;
}

.experience-text {
  max-width: 200px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  display: inline-block;
}

.no-experience {
  color: #c0c4cc;
}

.pagination-container {
  margin-top: 20px;
  text-align: right;
}
</style>