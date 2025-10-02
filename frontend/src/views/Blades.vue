<template>
  <div class="blades">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>刀片管理</span>
          <el-button type="primary" @click="showCreateDialog = true">
            <el-icon><Plus /></el-icon>
            新增刀片
          </el-button>
        </div>
      </template>

      <!-- 数据表格 -->
      <el-table
        v-loading="loading"
        :data="blades"
        style="width: 100%"
        height="400"
      >
        <el-table-column prop="brand" label="品牌" width="120" />
        <el-table-column prop="model" label="型号" width="150" />
        <el-table-column label="购买日期" width="120">
          <template #default="{ row }">
            {{ row.purchase_date ? formatDate(row.purchase_date) : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="单价" width="80">
          <template #default="{ row }">
            {{ row.unit_price ? `¥${row.unit_price}` : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="总数量" width="80">
          <template #default="{ row }">
            {{ row.total_quantity }}
          </template>
        </el-table-column>
        <el-table-column label="剩余数量" width="100">
          <template #default="{ row }">
            <span :class="getQuantityClass(row.remaining_quantity, row.total_quantity)">
              {{ row.remaining_quantity }}
            </span>
          </template>
        </el-table-column>
        <el-table-column label="使用进度" width="120">
          <template #default="{ row }">
            <el-progress
              :percentage="getUsagePercentage(row.remaining_quantity, row.total_quantity)"
              :status="getProgressStatus(row.remaining_quantity, row.total_quantity)"
              :stroke-width="6"
            />
          </template>
        </el-table-column>
        <el-table-column prop="notes" label="备注" min-width="150">
          <template #default="{ row }">
            <span v-if="row.notes" class="notes-text">{{ row.notes }}</span>
            <span v-else class="no-notes">-</span>
          </template>
        </el-table-column>
        <el-table-column label="创建时间" width="180">
          <template #default="{ row }">
            {{ formatDateTime(row.created_at) }}
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
    <BladeDialog
      v-model="showCreateDialog"
      :blade="currentBlade"
      @success="handleDialogSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useBladeStore } from '@/stores'
import { storeToRefs } from 'pinia'
import BladeDialog from '@/components/BladeDialog.vue'
import dayjs from 'dayjs'
import type { Blade } from '@/types'

const bladeStore = useBladeStore()
const { blades, loading, total } = storeToRefs(bladeStore)

const showCreateDialog = ref(false)
const currentBlade = ref<Blade | null>(null)

const pagination = reactive({
  page: 1,
  page_size: 20
})

const formatDate = (dateString: string) => {
  return dayjs(dateString).format('YYYY-MM-DD')
}

const formatDateTime = (dateString: string) => {
  return dayjs(dateString).format('YYYY-MM-DD HH:mm')
}

const getQuantityClass = (remaining: number, total: number) => {
  const percentage = (remaining / total) * 100
  if (percentage <= 10) return 'quantity-critical'
  if (percentage <= 30) return 'quantity-warning'
  return 'quantity-normal'
}

const getUsagePercentage = (remaining: number, total: number) => {
  return Math.round(((total - remaining) / total) * 100)
}

const getProgressStatus = (remaining: number, total: number) => {
  const percentage = (remaining / total) * 100
  if (percentage <= 10) return 'exception'
  if (percentage <= 30) return 'warning'
  return 'success'
}

const handleSizeChange = (newSize: number) => {
  pagination.page_size = newSize
  fetchData()
}

const handleCurrentChange = (newPage: number) => {
  pagination.page = newPage
  fetchData()
}

const handleEdit = (blade: Blade) => {
  currentBlade.value = blade
  showCreateDialog.value = true
}

const handleDelete = async (blade: Blade) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除刀片 "${blade.brand} ${blade.model}" 吗？`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )

    await bladeStore.deleteBlade(blade.id)
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
  currentBlade.value = null
  fetchData()
}

const fetchData = () => {
  bladeStore.fetchBlades({
    page: pagination.page,
    page_size: pagination.page_size
  })
}

onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.blades {
  padding: 0;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.notes-text {
  max-width: 150px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  display: inline-block;
}

.no-notes {
  color: #c0c4cc;
}

.quantity-normal {
  color: #67c23a;
  font-weight: bold;
}

.quantity-warning {
  color: #e6a23c;
  font-weight: bold;
}

.quantity-critical {
  color: #f56c6c;
  font-weight: bold;
}

.pagination-container {
  margin-top: 20px;
  text-align: right;
}
</style>