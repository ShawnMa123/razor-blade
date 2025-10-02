<template>
  <el-dialog
    v-model="visible"
    :title="isEdit ? '编辑刀片' : '新增刀片'"
    width="500px"
    @close="handleClose"
  >
    <el-form
      ref="formRef"
      :model="form"
      :rules="rules"
      label-width="100px"
    >
      <el-form-item label="品牌" prop="brand">
        <el-input v-model="form.brand" placeholder="请输入品牌" />
      </el-form-item>

      <el-form-item label="型号" prop="model">
        <el-input v-model="form.model" placeholder="请输入型号" />
      </el-form-item>

      <el-form-item label="兼容剃须刀">
        <el-input
          v-model="form.compatible_razors"
          placeholder="请输入兼容的剃须刀型号"
        />
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

      <el-form-item label="单价">
        <el-input-number
          v-model="form.unit_price"
          :min="0"
          :precision="2"
          placeholder="请输入单价"
          style="width: 100%"
        />
      </el-form-item>

      <el-form-item label="总数量" prop="total_quantity">
        <el-input-number
          v-model="form.total_quantity"
          :min="0"
          placeholder="请输入总数量"
          style="width: 100%"
        />
      </el-form-item>

      <el-form-item label="剩余数量" prop="remaining_quantity">
        <el-input-number
          v-model="form.remaining_quantity"
          :min="0"
          :max="form.total_quantity"
          placeholder="请输入剩余数量"
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
import { useBladeStore } from '@/stores'
import { storeToRefs } from 'pinia'
import type { Blade } from '@/types'

interface Props {
  modelValue: boolean
  blade?: Blade | null
}

interface Emits {
  (e: 'update:modelValue', value: boolean): void
  (e: 'success'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const formRef = ref<FormInstance>()
const bladeStore = useBladeStore()
const { loading } = storeToRefs(bladeStore)

const visible = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value)
})

const isEdit = computed(() => !!props.blade)

const form = reactive({
  brand: '',
  model: '',
  compatible_razors: '',
  purchase_date: '',
  unit_price: undefined as number | undefined,
  total_quantity: 0,
  remaining_quantity: 0,
  notes: ''
})

const rules: FormRules = {
  brand: [
    { required: true, message: '请输入品牌', trigger: 'blur' }
  ],
  model: [
    { required: true, message: '请输入型号', trigger: 'blur' }
  ],
  total_quantity: [
    { required: true, message: '请输入总数量', trigger: 'blur' }
  ],
  remaining_quantity: [
    { required: true, message: '请输入剩余数量', trigger: 'blur' }
  ]
}

const resetForm = () => {
  form.brand = ''
  form.model = ''
  form.compatible_razors = ''
  form.purchase_date = ''
  form.unit_price = undefined
  form.total_quantity = 0
  form.remaining_quantity = 0
  form.notes = ''

  if (formRef.value) {
    formRef.value.clearValidate()
  }
}

const loadFormData = (blade: Blade) => {
  form.brand = blade.brand
  form.model = blade.model
  form.compatible_razors = blade.compatible_razors || ''
  form.purchase_date = blade.purchase_date || ''
  form.unit_price = blade.unit_price
  form.total_quantity = blade.total_quantity
  form.remaining_quantity = blade.remaining_quantity
  form.notes = blade.notes || ''
}

const handleSubmit = async () => {
  if (!formRef.value) return

  try {
    await formRef.value.validate()

    const submitData = {
      ...form,
      purchase_date: form.purchase_date || undefined,
      unit_price: form.unit_price || undefined
    }

    if (isEdit.value && props.blade) {
      await bladeStore.updateBlade(props.blade.id, submitData)
      ElMessage.success('刀片更新成功')
    } else {
      await bladeStore.createBlade(submitData)
      ElMessage.success('刀片创建成功')
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

// 监听总数量变化，自动更新剩余数量
watch(() => form.total_quantity, (newTotal) => {
  if (!isEdit.value && form.remaining_quantity === 0) {
    form.remaining_quantity = newTotal
  }
})

// 监听 dialog 打开和刀片数据变化
watch([() => props.modelValue, () => props.blade], ([show, blade]) => {
  if (show) {
    if (blade) {
      loadFormData(blade)
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