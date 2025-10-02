import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type {
  Razor,
  Blade,
  UsageRecord,
  DashboardData,
  Statistics,
  PaginationRequest
} from '@/types'
import {
  razorAPI,
  bladeAPI,
  usageRecordAPI,
  statisticsAPI
} from '@/api'

// 剃须刀状态管理
export const useRazorStore = defineStore('razor', () => {
  const razors = ref<Razor[]>([])
  const currentRazor = ref<Razor | null>(null)
  const loading = ref(false)
  const total = ref(0)

  const fetchRazors = async (params?: PaginationRequest) => {
    loading.value = true
    try {
      const response = await razorAPI.getList(params)
      if (response.success && response.data) {
        razors.value = response.data.items
        total.value = response.data.total
      }
    } catch (error) {
      console.error('获取剃须刀列表失败:', error)
    } finally {
      loading.value = false
    }
  }

  const fetchRazorById = async (id: number) => {
    loading.value = true
    try {
      const response = await razorAPI.getById(id)
      if (response.success && response.data) {
        currentRazor.value = response.data
      }
    } catch (error) {
      console.error('获取剃须刀详情失败:', error)
    } finally {
      loading.value = false
    }
  }

  const createRazor = async (data: any) => {
    loading.value = true
    try {
      const response = await razorAPI.create(data)
      if (response.success) {
        await fetchRazors()
        return response
      }
    } catch (error) {
      console.error('创建剃须刀失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const updateRazor = async (id: number, data: any) => {
    loading.value = true
    try {
      const response = await razorAPI.update(id, data)
      if (response.success) {
        await fetchRazors()
        return response
      }
    } catch (error) {
      console.error('更新剃须刀失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const deleteRazor = async (id: number) => {
    loading.value = true
    try {
      const response = await razorAPI.delete(id)
      if (response.success) {
        await fetchRazors()
        return response
      }
    } catch (error) {
      console.error('删除剃须刀失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  return {
    razors,
    currentRazor,
    loading,
    total,
    fetchRazors,
    fetchRazorById,
    createRazor,
    updateRazor,
    deleteRazor
  }
})

// 刀片状态管理
export const useBladeStore = defineStore('blade', () => {
  const blades = ref<Blade[]>([])
  const currentBlade = ref<Blade | null>(null)
  const loading = ref(false)
  const total = ref(0)

  const fetchBlades = async (params?: PaginationRequest) => {
    loading.value = true
    try {
      const response = await bladeAPI.getList(params)
      if (response.success && response.data) {
        blades.value = response.data.items
        total.value = response.data.total
      }
    } catch (error) {
      console.error('获取刀片列表失败:', error)
    } finally {
      loading.value = false
    }
  }

  const fetchBladeById = async (id: number) => {
    loading.value = true
    try {
      const response = await bladeAPI.getById(id)
      if (response.success && response.data) {
        currentBlade.value = response.data
      }
    } catch (error) {
      console.error('获取刀片详情失败:', error)
    } finally {
      loading.value = false
    }
  }

  const createBlade = async (data: any) => {
    loading.value = true
    try {
      const response = await bladeAPI.create(data)
      if (response.success) {
        await fetchBlades()
        return response
      }
    } catch (error) {
      console.error('创建刀片失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const updateBlade = async (id: number, data: any) => {
    loading.value = true
    try {
      const response = await bladeAPI.update(id, data)
      if (response.success) {
        await fetchBlades()
        return response
      }
    } catch (error) {
      console.error('更新刀片失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const deleteBlade = async (id: number) => {
    loading.value = true
    try {
      const response = await bladeAPI.delete(id)
      if (response.success) {
        await fetchBlades()
        return response
      }
    } catch (error) {
      console.error('删除刀片失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  return {
    blades,
    currentBlade,
    loading,
    total,
    fetchBlades,
    fetchBladeById,
    createBlade,
    updateBlade,
    deleteBlade
  }
})

// 使用记录状态管理
export const useUsageRecordStore = defineStore('usageRecord', () => {
  const usageRecords = ref<UsageRecord[]>([])
  const currentUsageRecord = ref<UsageRecord | null>(null)
  const loading = ref(false)
  const total = ref(0)

  const fetchUsageRecords = async (params?: PaginationRequest) => {
    loading.value = true
    try {
      const response = await usageRecordAPI.getList(params)
      if (response.success && response.data) {
        usageRecords.value = response.data.items
        total.value = response.data.total
      }
    } catch (error) {
      console.error('获取使用记录列表失败:', error)
    } finally {
      loading.value = false
    }
  }

  const fetchUsageRecordById = async (id: number) => {
    loading.value = true
    try {
      const response = await usageRecordAPI.getById(id)
      if (response.success && response.data) {
        currentUsageRecord.value = response.data
      }
    } catch (error) {
      console.error('获取使用记录详情失败:', error)
    } finally {
      loading.value = false
    }
  }

  const createUsageRecord = async (data: any) => {
    loading.value = true
    try {
      const response = await usageRecordAPI.create(data)
      if (response.success) {
        await fetchUsageRecords()
        return response
      }
    } catch (error) {
      console.error('创建使用记录失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const updateUsageRecord = async (id: number, data: any) => {
    loading.value = true
    try {
      const response = await usageRecordAPI.update(id, data)
      if (response.success) {
        await fetchUsageRecords()
        return response
      }
    } catch (error) {
      console.error('更新使用记录失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  const deleteUsageRecord = async (id: number) => {
    loading.value = true
    try {
      const response = await usageRecordAPI.delete(id)
      if (response.success) {
        await fetchUsageRecords()
        return response
      }
    } catch (error) {
      console.error('删除使用记录失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  return {
    usageRecords,
    currentUsageRecord,
    loading,
    total,
    fetchUsageRecords,
    fetchUsageRecordById,
    createUsageRecord,
    updateUsageRecord,
    deleteUsageRecord
  }
})

// 统计数据状态管理
export const useStatisticsStore = defineStore('statistics', () => {
  const dashboardData = ref<DashboardData | null>(null)
  const statistics = ref<Statistics | null>(null)
  const loading = ref(false)

  const fetchDashboardData = async () => {
    loading.value = true
    try {
      const response = await statisticsAPI.getDashboard()
      if (response.success && response.data) {
        dashboardData.value = response.data
      }
    } catch (error) {
      console.error('获取仪表板数据失败:', error)
    } finally {
      loading.value = false
    }
  }

  const fetchStatistics = async () => {
    loading.value = true
    try {
      const response = await statisticsAPI.getStatistics()
      if (response.success && response.data) {
        statistics.value = response.data
      }
    } catch (error) {
      console.error('获取统计数据失败:', error)
    } finally {
      loading.value = false
    }
  }

  return {
    dashboardData,
    statistics,
    loading,
    fetchDashboardData,
    fetchStatistics
  }
})