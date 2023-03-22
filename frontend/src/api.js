import axios from 'axios'

/* PORT number and url for the API */
const API_PORT = 8080
const API_URL = `http://localhost:${API_PORT}/api/v1`

/* A full axios instance with the base URL and it's headers */
const axiosInstance = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json'
  }
})

/* universal endpoint used for both GET & POST */
const endpoint = '/entry'

/*
 * Makes calls to API
 *
 * @param { String } method - This value will either be 'post' or 'get'
 * @param { String } url - API endpoint
 * @param { Object } data - Data object to be passed to API
 * 
 * Return { Object } result - Object either returning an error or data
 */
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
    result = { error: e.response.data.error, data: null }
  }
  return result
}

export { AxiosCall }
