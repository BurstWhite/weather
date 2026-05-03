import axios from 'axios'

const api = axios.create({
  baseURL: '/api',
  timeout: 15000,
})

export function fetchCities() {
  return api.get('/cities')
}

export function addCity(city) {
  return api.post('/cities', city)
}

export function deleteCity(id) {
  return api.delete(`/cities/${id}`)
}

export function searchCity(q) {
  return api.get('/search', { params: { q } })
}

export function fetchWeather(id) {
  return api.get(`/weather/${id}`)
}

export default api
