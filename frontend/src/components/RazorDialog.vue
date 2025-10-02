<template>
  <el-dialog
    v-model="visible"
    :title="isEdit ? '编辑剃须刀' : '新增剃须刀'"
    width="500px"
    @close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="80px"
    >
      <el-form-item label="品牌" prop="brand">
        <el-input v-model="form.brand" placeholder="请输入品牌" />
      </el-form-item>

      <el-form-item label="型号" prop="model">
        <el-input v-model="form.model" placeholder="请输入型号" />
      </el-form-item>

      <el-form-item label="购买日期">
        <el-date-picker
          v-model="form.purchase_date"
          type="date"
          placeholder="选择购买日期"
          format="YYYY-MM-DD"
          value-format="YYYY-MM-DD"
          style="width: 100%"
        />
      </el-form-item>

      <el-form-item label="价格">
        <el-input-number
          v-model="form.price"
          :min="0"
          :precision="2"
          placeholder="请输入价格"
          style="width: 100%"
        />
      </el-form-item>

      <el-form-item label="备注">
        <el-input
          v-model="form.notes"
          type="textarea"
          :rows="3"
          placeholder="请输入备注信息..."
        />
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
import { useRazorStore } from '@/stores'
import { storeToRefs } from 'pinia'
import type { Razor } from '@/types'

interface Props {
  modelValue: boolean
  razor?: Razor | null
}

interface Emits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'success'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const formRef = ref<FormInstance>()
const razorStore = useRazorStore()
const { loading } = storeToRefs(razorStore)

const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const isEdit = computed(() => !!props.razor)

const form = reactive({
  brand: '',
  model: '',
  purchase_date: '',
  price: undefined as number | undefined,
  notes: ''
})

const rules: FormRules = {
  brand: [
    { required: true, message: '请输入品牌', trigger: 'blur' }
  ],
  model: [
    { required: true, message: '请输入型号', trigger: 'blur' }
  ]
}

const resetForm = () => {
  form.brand = ''
  form.model = ''
  form.purchase_date = ''
  form.price = undefined
  form.notes = ''

  if (formRef.value) {
    formRef.value.clearValidate()
  }
}

const loadFormData = (razor: Razor) => {
  form.brand = razor.brand
  form.model = razor.model
  form.purchase_date = razor.purchase_date || ''
  form.price = razor.price
  form.notes = razor.notes || ''
}

const handleSubmit = async () => {
  if (!formRef.value) return

  try {
    await formRef.value.validate()

    const submitData = {
      ...form,
      purchase_date: form.purchase_date || undefined,
      price: form.price || undefined
    }

    if (isEdit.value && props.razor) {
      await razorStore.updateRazor(props.razor.id, submitData)
      ElMessage.success('剃须刀更新成功')
    } else {
      await razorStore.createRazor(submitData)
      ElMessage.success('剃须刀创建成功')
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

// 监听 dialog 打开和剃须刀数据变化
watch([() => props.modelValue, () => props.razor], ([show, razor]) => {
  if (show) {
    if (razor) {
      loadFormData(razor)
    } else {
      resetForm()
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