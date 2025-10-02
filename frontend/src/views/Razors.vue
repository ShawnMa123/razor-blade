<template>
  <div class="razors">
    <el-card>
      <template #header>
        <div class="card-header">
          <span>剃须刀管理</span>
          <el-button type="primary" @click="showCreateDialog = true">
            <el-icon><Plus /></el-icon>
            新增剃须刀
          </el-button>
        </div>
      </template>

      <!-- 数据表格 -->
      <el-table
        v-loading="loading"
        :data="razors"
        style="width: 100%"
        height="400"
      >
        <el-table-column prop="brand" label="品牌" width="150" />
        <el-table-column prop="model" label="型号" width="200" />
        <el-table-column label="购买日期" width="120">
          <template #default="{ row }">
            {{ row.purchase_date ? formatDate(row.purchase_date) : '-' }}
          </template>
        </el-table-column>
        <el-table-column label="价格" width="100">
          <template #default="{ row }">
            {{ row.price ? `¥${row.price}` : '-' }}
          </template>
        </el-table-column>
        <el-table-column prop="notes" label="备注" min-width="200">
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
    <RazorDialog
      v-model="showCreateDialog"
      :razor="currentRazor"
      @success="handleDialogSuccess"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRazorStore } from '@/stores'
import { storeToRefs } from 'pinia'
import RazorDialog from '@/components/RazorDialog.vue'
import dayjs from 'dayjs'
import type { Razor } from '@/types'

const razorStore = useRazorStore()
const { razors, loading, total } = storeToRefs(razorStore)

const showCreateDialog = ref(false)
const currentRazor = ref<Razor | null>(null)

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

const handleSizeChange = (newSize: number) => {
  pagination.page_size = newSize
  fetchData()
}

const handleCurrentChange = (newPage: number) => {
  pagination.page = newPage
  fetchData()
}

const handleEdit = (razor: Razor) => {
  currentRazor.value = razor
  showCreateDialog.value = true
}

const handleDelete = async (razor: Razor) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除剃须刀 "${razor.brand} ${razor.model}" 吗？`,
      '确认删除',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )

    await razorStore.deleteRazor(razor.id)
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
  currentRazor.value = null
  fetchData()
}

const fetchData = () => {
  razorStore.fetchRazors({
    page: pagination.page,
    page_size: pagination.page_size
  })
}

onMounted(() => {
  fetchData()
})
</script>

<style scoped>
.razors {
  padding: 0;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.notes-text {
  max-width: 200px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  display: inline-block;
}

.no-notes {
  color: #c0c4cc;
}

.pagination-container {
  margin-top: 20px;
  text-align: right;
}
</style>