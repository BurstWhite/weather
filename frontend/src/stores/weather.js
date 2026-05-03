import { defineStore } from 'pinia'
import { fetchCities, addCity as apiAddCity, deleteCity as apiDeleteCity, fetchWeather } from '../api'

let tempIdCounter = 0

export const useWeatherStore = defineStore('weather', {
  state: () => ({
    cities: [],
    selectedCity: null,
    weatherData: null,
    loading: false,
    error: null,
  }),

  actions: {
    async loadCities() {
      try {
        const res = await fetchCities()
        this.cities = res.data
      } catch (e) {
        this.error = 'Failed to load cities'
      }
    },

    addCityTemp(searchResult) {
      const tempId = `temp-${++tempIdCounter}`
      const tempCity = {
        id: tempId,
        locationId: searchResult.locationId,
        name: searchResult.name,
        adm1: searchResult.adm1 || '',
        adm2: searchResult.adm2 || '',
        lat: searchResult.lat,
        lon: searchResult.lon,
      }
      this.cities.push(tempCity)
      return tempCity
    },

    replaceCity(tempId, realCity) {
      const idx = this.cities.findIndex(c => c.id === tempId)
      if (idx !== -1) {
        this.cities.splice(idx, 1, realCity)
      } else if (!this.cities.find(c => c.id === realCity.id)) {
        this.cities.push(realCity)
      }
      return realCity
    },

    async addCity(city) {
      const res = await apiAddCity(city)
      if (!this.cities.find(c => c.locationId === res.data.locationId)) {
        this.cities.push(res.data)
      }
      return res.data
    },

    async removeCity(id) {
      await apiDeleteCity(id)
      this.cities = this.cities.filter(c => c.id !== id)
      if (this.selectedCity?.id === id) {
        this.selectedCity = null
        this.weatherData = null
      }
    },

    async selectCity(city) {
      this.selectedCity = city
      this.loading = true
      this.error = null
      try {
        const res = await fetchWeather(city.id)
        this.weatherData = res.data
      } catch (e) {
        this.error = 'Failed to fetch weather data'
      } finally {
        this.loading = false
      }
    },
  },
})
