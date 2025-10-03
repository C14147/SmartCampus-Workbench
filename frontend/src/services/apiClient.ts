import axios, { AxiosError, AxiosInstance } from 'axios'

const client: AxiosInstance = axios.create({ baseURL: '/api/v1', timeout: 10000 })

client.interceptors.response.use(
  (r) => r,
  (err: AxiosError) => {
    // Normalize error
    if (err.response && err.response.data) {
      return Promise.reject(err.response.data)
    }
    return Promise.reject({ error: err.message || 'network_error' })
  }
)

type ApiResponse<T> = T

export const ApiClient = {
  setToken: (token: string) => {
    if (token) client.defaults.headers.common['Authorization'] = `Bearer ${token}`
    else delete client.defaults.headers.common['Authorization']
  },
  async post<T = any>(url: string, data?: unknown): Promise<ApiResponse<T>> {
    const r = await client.post<T>(url, data)
    return r.data
  },
  async get<T = any>(url: string): Promise<ApiResponse<T>> {
    const r = await client.get<T>(url)
    return r.data
  }
}

export default ApiClient
