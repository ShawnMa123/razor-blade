export interface Razor {
  id: number
  brand: string
  model: string
  purchase_date?: string
  price?: number
  notes: string
  created_at: string
  updated_at: string
}

export interface Blade {
  id: number
  brand: string
  model: string
  compatible_razors: string
  purchase_date?: string
  unit_price?: number
  total_quantity: number
  remaining_quantity: number
  notes: string
  created_at: string
  updated_at: string
}

export interface UsageRecord {
  id: number
  usage_time: string
  razor_id: number
  blade_id: number
  blade_usage_count: number
  rating?: number
  experience_text: string
  need_blade_change: boolean
  created_at: string
  updated_at: string
  razor: Razor
  blade: Blade
}

export interface CreateRazorRequest {
  brand: string
  model: string
  purchase_date?: string
  price?: number
  notes?: string
}

export interface UpdateRazorRequest {
  brand?: string
  model?: string
  purchase_date?: string
  price?: number
  notes?: string
}

export interface CreateBladeRequest {
  brand: string
  model: string
  compatible_razors?: string
  purchase_date?: string
  unit_price?: number
  total_quantity?: number
  remaining_quantity?: number
  notes?: string
}

export interface UpdateBladeRequest {
  brand?: string
  model?: string
  compatible_razors?: string
  purchase_date?: string
  unit_price?: number
  total_quantity?: number
  remaining_quantity?: number
  notes?: string
}

export interface CreateUsageRecordRequest {
  usage_time: string
  razor_id: number
  blade_id: number
  blade_usage_count?: number
  rating?: number
  experience_text?: string
  need_blade_change?: boolean
}

export interface UpdateUsageRecordRequest {
  usage_time?: string
  razor_id?: number
  blade_id?: number
  blade_usage_count?: number
  rating?: number
  experience_text?: string
  need_blade_change?: boolean
}

export interface APIResponse<T = any> {
  success: boolean
  data?: T
  message: string
  error?: string
}

export interface PaginationRequest {
  page?: number
  page_size?: number
}

export interface PaginationResponse<T = any> {
  items: T[]
  page: number
  page_size: number
  total: number
  total_pages: number
}

export interface Statistics {
  total_usage: number
  razor_count: number
  blade_count: number
  average_rating: number
}

export interface DashboardData {
  statistics: Statistics
  recent_records: UsageRecord[]
}