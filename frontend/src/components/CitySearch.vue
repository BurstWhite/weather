<template>
  <div class="search-overlay" @click.self="$emit('close')">
    <div class="search-modal">
      <div class="search-header">
        <h3>Add City</h3>
        <button class="close-btn" @click="$emit('close')">×</button>
      </div>
      <input
        v-model="query"
        type="text"
        placeholder="Search Chinese cities..."
        class="search-input"
        @input="onSearch"
      />
      <div v-if="loading" class="search-status">Searching...</div>
      <div v-else-if="error" class="search-status error">{{ error }}</div>
      <div v-else class="search-results">
        <div
          v-for="city in results"
          :key="city.id"
          class="result-item"
          @click="addCity(city)"
        >
          <span class="result-name">{{ city.name }}</span>
          <span class="result-adm">{{ city.adm1 }} / {{ city.adm2 }}</span>
        </div>
        <div v-if="query && !results.length && !loading" class="no-results">
          No cities found
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { searchCity } from '../api'
import { useWeatherStore } from '../stores/weather'

const emit = defineEmits(['close'])
const store = useWeatherStore()

const query = ref('')
const results = ref([])
const loading = ref(false)
const error = ref(null)

let debounceTimer = null

function onSearch() {
  clearTimeout(debounceTimer)
  if (query.value.length < 2) {
    results.value = []
    return
  }
  debounceTimer = setTimeout(doSearch, 300)
}

async function doSearch() {
  loading.value = true
  error.value = null
  try {
    const res = await searchCity(query.value)
    results.value = res.data || []
  } catch (e) {
    error.value = 'Search failed'
  } finally {
    loading.value = false
  }
}

async function addCity(city) {
  const existing = store.cities.find(c => c.locationId === city.id)
  if (existing) {
    emit('close')
    await store.selectCity(existing)
    return
  }

  const tempCity = store.addCityTemp({
    locationId: city.id,
    name: city.name,
    adm1: city.adm1,
    adm2: city.adm2,
    lat: parseFloat(city.lat),
    lon: parseFloat(city.lon),
  })
  emit('close')
  store.loading = true
  try {
    const added = await store.addCity({
      name: city.name,
      locationId: city.id,
      adm1: city.adm1,
      adm2: city.adm2,
      lat: parseFloat(city.lat),
      lon: parseFloat(city.lon),
    })
    const realCity = store.replaceCity(tempCity.id, added)
    await store.selectCity(realCity)
  } catch (e) {
    store.cities = store.cities.filter(c => c.id !== tempCity.id)
    if (store.selectedCity?.id === tempCity.id) {
      store.selectedCity = null
      store.weatherData = null
    }
    store.loading = false
    store.error = e.response?.data?.error || `Failed to add ${city.name}`
    setTimeout(() => { store.error = null }, 5000)
  }
}
</script>

<style scoped>
.search-overlay {
  position: fixed;
  inset: 0;
  background: rgba(32, 33, 35, 0.34);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
  backdrop-filter: blur(2px);
}

.search-modal {
  background: #ffffff;
  border: 1px solid #deded7;
  border-radius: 8px;
  width: 420px;
  max-width: calc(100vw - 32px);
  max-height: 80vh;
  display: flex;
  flex-direction: column;
  box-shadow: 0 18px 60px rgba(0, 0, 0, 0.18);
}

.search-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 18px 20px;
  border-bottom: 1px solid #ecece6;
}

.search-header h3 {
  font-size: 16px;
  color: #202123;
  font-weight: 650;
}

.close-btn {
  background: none;
  border: none;
  color: #6e6e80;
  font-size: 24px;
  cursor: pointer;
  border-radius: 6px;
  line-height: 1;
  width: 32px;
  height: 32px;
}

.close-btn:hover {
  background: #f3f3ef;
  color: #202123;
}

.search-input {
  margin: 16px 20px;
  padding: 12px 16px;
  border-radius: 8px;
  border: 1px solid #d9d9d0;
  background: #ffffff;
  color: #202123;
  font-size: 15px;
  outline: none;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.04);
}

.search-input:focus {
  border-color: #10a37f;
  box-shadow: 0 0 0 3px rgba(16, 163, 127, 0.12);
}

.search-status {
  padding: 12px 20px;
  color: #6e6e80;
  font-size: 14px;
}

.search-status.error {
  color: #b42318;
}

.search-results {
  flex: 1;
  overflow-y: auto;
  padding: 0 0 8px;
}

.result-item {
  padding: 12px 20px;
  cursor: pointer;
  display: flex;
  flex-direction: column;
  transition: background 0.15s;
}

.result-item:hover {
  background: #f7f7f4;
}

.result-name {
  font-size: 15px;
  font-weight: 500;
  color: #202123;
}

.result-adm {
  font-size: 12px;
  color: #6e6e80;
  margin-top: 2px;
}

.no-results {
  padding: 12px 20px;
  color: #6e6e80;
  font-size: 14px;
  text-align: center;
}
</style>
