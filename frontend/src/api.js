import axios from 'axios'

const API_PORT = 8080
const API_URL = `http://localhost:${API_PORT}/api/v1`

const axiosInstance = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json'
  }
})

const endpoint = '/entry'

async function AxiosCall(method = 'get', url = endpoint, data = null) {
  let result
  try {
    const response = await axiosInstance({
      method: method,
      url: url,
      data: data
    })
    result = { error: null, ...response.data }
  } catch (e) {
    result = { error: e.status, data: null }
  }
  return result
}

export { AxiosCall }
