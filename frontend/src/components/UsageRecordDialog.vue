<template>
  <el-dialog
    v-model="visible"
    :title="isEdit ? '编辑使用记录' : '新增使用记录'"
    width="600px"
    @close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="100px"
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
          :rows="4"
          placeholder="记录使用感受..."
        />
      </el-form-item>

      <el-form-item>
        <el-checkbox v-model="form.need_blade_change">
          需要更换刀片
        </el-checkbox>
      </el-form-item>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="handleClose">取消</el-button>
        <el-button type="primary" :loading="loading" @click="handleSubmit">
          {{ isEdit ? '更新' : '创建' }}
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, computed, watch } from 'vue'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import { useUsageRecordStore, useRazorStore, useBladeStore } from '@/stores'
import { storeToRefs } from 'pinia'
import dayjs from 'dayjs'
import type { UsageRecord } from '@/types'

interface Props {
  modelValue: boolean
  record?: UsageRecord | null
}

interface Emits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'success'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const formRef = ref<FormInstance>()
const usageRecordStore = useUsageRecordStore()
const razorStore = useRazorStore()
const bladeStore = useBladeStore()

const { loading } = storeToRefs(usageRecordStore)
const { razors } = storeToRefs(razorStore)
const { blades } = storeToRefs(bladeStore)

const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const isEdit = computed(() => !!props.record)

const form = reactive({
  usage_time: '',
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

const loadFormData = (record: UsageRecord) => {
  form.usage_time = dayjs(record.usage_time).format('YYYY-MM-DDTHH:mm:ss.SSSZ')
  form.razor_id = record.razor_id
  form.blade_id = record.blade_id
  form.blade_usage_count = record.blade_usage_count
  form.rating = record.rating || 0
  form.experience_text = record.experience_text || ''
  form.need_blade_change = record.need_blade_change
}

const handleSubmit = async () => {
  if (!formRef.value) return

  try {
    await formRef.value.validate()

    const submitData = {
      ...form,
      rating: form.rating > 0 ? form.rating : undefined
    }

    if (isEdit.value && props.record) {
      await usageRecordStore.updateUsageRecord(props.record.id, submitData)
      ElMessage.success('使用记录更新成功')
    } else {
      await usageRecordStore.createUsageRecord(submitData)
      ElMessage.success('使用记录创建成功')
    }

    emit('success')
  } catch (error) {
    console.error('操作失败:', error)
    ElMessage.error(isEdit.value ? '更新失败' : '创建失败')
  }
}

const handleClose = () => {
  visible.value = false
  resetForm()
}

// 监听 dialog 打开和记录变化
watch([() => props.modelValue, () => props.record], ([show, record]) => {
  if (show) {
    if (record) {
      loadFormData(record)
    } else {
      resetForm()
    }
  }
})

// 确保数据已加载
watch(visible, (show) => {
  if (show) {
    if (razors.value.length === 0) {
      razorStore.fetchRazors({ page: 1, page_size: 100 })
    }
    if (blades.value.length === 0) {
      bladeStore.fetchBlades({ page: 1, page_size: 100 })
    }
  }
})
</script>

<style scoped>
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
</style>