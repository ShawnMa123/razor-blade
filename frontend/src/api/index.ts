import axios from 'axios'
import type {
  APIResponse,
  PaginationResponse,
  PaginationRequest,
  Razor,
  Blade,
  UsageRecord,
  CreateRazorRequest,
  UpdateRazorRequest,
  CreateBladeRequest,
  UpdateBladeRequest,
  CreateUsageRecordRequest,
  UpdateUsageRecordRequest,
  DashboardData,
  Statistics
} from '@/types'

const api = axios.create({
  baseURL: '/api/v1',
  timeout: 10000,
})

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    console.error('API Error:', error)
    return Promise.reject(error)
  }
)

// 剃须刀相关API
export const razorAPI = {
  create: (data: CreateRazorRequest): Promise<APIResponse<Razor>> =>
    api.post('/razors', data),

  getById: (id: number): Promise<APIResponse<Razor>> =>
    api.get(`/razors/${id}`),

  getList: (params?: PaginationRequest): Promise<APIResponse<PaginationResponse<Razor>>> =>
    api.get('/razors', { params }),

  update: (id: number, data: UpdateRazorRequest): Promise<APIResponse<Razor>> =>
    api.put(`/razors/${id}`, data),

  delete: (id: number): Promise<APIResponse<null>> =>
    api.delete(`/razors/${id}`)
}

// 刀片相关API
export const bladeAPI = {
  create: (data: CreateBladeRequest): Promise<APIResponse<Blade>> =>
    api.post('/blades', data),

  getById: (id: number): Promise<APIResponse<Blade>> =>
    api.get(`/blades/${id}`),

  getList: (params?: PaginationRequest): Promise<APIResponse<PaginationResponse<Blade>>> =>
    api.get('/blades', { params }),

  update: (id: number, data: UpdateBladeRequest): Promise<APIResponse<Blade>> =>
    api.put(`/blades/${id}`, data),

  delete: (id: number): Promise<APIResponse<null>> =>
    api.delete(`/blades/${id}`)
}

// 使用记录相关API
export const usageRecordAPI = {
  create: (data: CreateUsageRecordRequest): Promise<APIResponse<UsageRecord>> =>
    api.post('/usage-records', data),

  getById: (id: number): Promise<APIResponse<UsageRecord>> =>
    api.get(`/usage-records/${id}`),

  getList: (params?: PaginationRequest): Promise<APIResponse<PaginationResponse<UsageRecord>>> =>
    api.get('/usage-records', { params }),

  update: (id: number, data: UpdateUsageRecordRequest): Promise<APIResponse<UsageRecord>> =>
    api.put(`/usage-records/${id}`, data),

  delete: (id: number): Promise<APIResponse<null>> =>
    api.delete(`/usage-records/${id}`)
}

// 统计相关API
export const statisticsAPI = {
  getDashboard: (): Promise<APIResponse<DashboardData>> =>
    api.get('/dashboard'),

  getStatistics: (): Promise<APIResponse<Statistics>> =>
    api.get('/statistics')
}

export default api