<template>
  <div class="quick-record-form">
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="80px"
      @submit.prevent="handleSubmit"
    >
      <el-form-item label="使用时间" prop="usage_time">
        <el-date-picker
          v-model="form.usage_time"
          type="datetime"
          placeholder="选择时间"
          format="YYYY-MM-DD HH:mm"
          value-format="YYYY-MM-DDTHH:mm:ss.SSSZ"
          style="width: 100%"
        />
      </el-form-item>

      <el-form-item label="剃须刀" prop="razor_id">
        <el-select
          v-model="form.razor_id"
          placeholder="选择剃须刀"
          style="width: 100%"
          @focus="loadRazors"
        >
          <el-option
            v-for="razor in razors"
            :key="razor.id"
            :label="`${razor.brand} ${razor.model}`"
            :value="razor.id"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="刀片" prop="blade_id">
        <el-select
          v-model="form.blade_id"
          placeholder="选择刀片"
          style="width: 100%"
          @focus="loadBlades"
        >
          <el-option
            v-for="blade in blades"
            :key="blade.id"
            :label="`${blade.brand} ${blade.model}`"
            :value="blade.id"
          />
        </el-select>
      </el-form-item>

      <el-form-item label="使用次数" prop="blade_usage_count">
        <el-input-number
          v-model="form.blade_usage_count"
          :min="1"
          :max="100"
          style="width: 100%"
        />
      </el-form-item>

      <el-form-item label="评分">
        <el-rate v-model="form.rating" show-score />
      </el-form-item>

      <el-form-item label="使用感受">
        <el-input
          v-model="form.experience_text"
          type="textarea"
          :rows="3"
          placeholder="记录使用感受..."
        />
      </el-form-item>

      <el-form-item>
        <el-checkbox v-model="form.need_blade_change">
          需要更换刀片
        </el-checkbox>
      </el-form-item>

      <el-form-item>
        <el-button
          type="primary"
          :loading="loading"
          @click="handleSubmit"
          style="width: 100%"
        >
          记录使用情况
        </el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import { useRazorStore, useBladeStore, useUsageRecordStore } from '@/stores'
import { storeToRefs } from 'pinia'
import dayjs from 'dayjs'

const emit = defineEmits<{
  recordCreated: []
}>()

const formRef = ref<FormInstance>()
const razorStore = useRazorStore()
const bladeStore = useBladeStore()
const usageRecordStore = useUsageRecordStore()

const { razors } = storeToRefs(razorStore)
const { blades } = storeToRefs(bladeStore)
const { loading } = storeToRefs(usageRecordStore)

const form = reactive({
  usage_time: dayjs().format('YYYY-MM-DDTHH:mm:ss.SSSZ'),
  razor_id: undefined as number | undefined,
  blade_id: undefined as number | undefined,
  blade_usage_count: 1,
  rating: 0,
  experience_text: '',
  need_blade_change: false
})

const rules: FormRules = {
  usage_time: [
    { required: true, message: '请选择使用时间', trigger: 'change' }
  ],
  razor_id: [
    { required: true, message: '请选择剃须刀', trigger: 'change' }
  ],
  blade_id: [
    { required: true, message: '请选择刀片', trigger: 'change' }
  ],
  blade_usage_count: [
    { required: true, message: '请输入使用次数', trigger: 'blur' }
  ]
}

const loadRazors = async () => {
  if (razors.value.length === 0) {
    await razorStore.fetchRazors({ page: 1, page_size: 100 })
  }
}

const loadBlades = async () => {
  if (blades.value.length === 0) {
    await bladeStore.fetchBlades({ page: 1, page_size: 100 })
  }
}

const handleSubmit = async () => {
  if (!formRef.value) return

  try {
    await formRef.value.validate()

    const submitData = {
      ...form,
      rating: form.rating > 0 ? form.rating : undefined
    }

    await usageRecordStore.createUsageRecord(submitData)

    ElMessage.success('使用记录创建成功')

    // 重置表单
    resetForm()

    // 通知父组件
    emit('recordCreated')
  } catch (error) {
    console.error('创建使用记录失败:', error)
    ElMessage.error('创建使用记录失败')
  }
}

const resetForm = () => {
  form.usage_time = dayjs().format('YYYY-MM-DDTHH:mm:ss.SSSZ')
  form.razor_id = undefined
  form.blade_id = undefined
  form.blade_usage_count = 1
  form.rating = 0
  form.experience_text = ''
  form.need_blade_change = false

  if (formRef.value) {
    formRef.value.clearValidate()
  }
}

onMounted(() => {
  loadRazors()
  loadBlades()
})
</script>

<style scoped>
.quick-record-form {
  padding: 0;
}

:deep(.el-form-item) {
  margin-bottom: 18px;
}

:deep(.el-form-item__label) {
  font-weight: 500;
}
</style>