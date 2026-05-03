<template>
  <aside class="sidebar">
    <div class="sidebar-header">
      <h2>Weather App</h2>
      <button class="add-btn" @click="showSearch = true">+</button>
    </div>
    <div v-if="error" class="error-msg">{{ error }}</div>
    <div class="city-list">
      <div
        v-for="city in store.cities"
        :key="city.id"
        class="city-item"
        :class="{
          active: store.selectedCity?.id === city.id,
          pending: typeof city.id === 'string' && city.id.startsWith('temp-')
        }"
        @click="typeof city.id === 'string' && city.id.startsWith('temp-') ? null : store.selectCity(city)"
      >
        <div class="city-info">
          <span class="city-name">{{ city.name }}</span>
          <span class="city-adm">{{ city.adm1 }}</span>
        </div>
        <button v-if="!(typeof city.id === 'string' && city.id.startsWith('temp-'))" class="remove-btn" @click.stop="removeCity(city)">×</button>
      </div>
    </div>
    <CitySearch v-if="showSearch" @close="showSearch = false" />
  </aside>
</template>

<script setup>
import { ref } from 'vue'
import { useWeatherStore } from '../stores/weather'
import CitySearch from './CitySearch.vue'

const store = useWeatherStore()
const showSearch = ref(false)
const error = ref(null)

async function removeCity(city) {
  error.value = null
  try {
    await store.removeCity(city.id)
  } catch (e) {
    error.value = e.response?.data?.error || `Failed to remove ${city.name}`
    setTimeout(() => { error.value = null }, 5000)
  }
}
</script>

<style scoped>
.sidebar {
  width: 292px;
  min-width: 292px;
  background: #f3f3ef;
  border-right: 1px solid #deded7;
  display: flex;
  flex-direction: column;
  height: 100vh;
}

.sidebar-header {
  padding: 18px 16px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid #deded7;
}

.sidebar-header h2 {
  font-size: 16px;
  color: #202123;
  font-weight: 650;
}

.add-btn {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  border: none;
  background: #202123;
  color: #fff;
  font-size: 19px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  line-height: 1;
  transition: background 0.15s, transform 0.15s;
}

.add-btn:hover {
  background: #353740;
  transform: translateY(-1px);
}

.city-list {
  flex: 1;
  overflow-y: auto;
  padding: 10px 8px;
}

.city-item {
  padding: 10px 12px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  border-radius: 8px;
  transition: background 0.15s, color 0.15s;
}

.city-item:hover {
  background: #e8e8e3;
}

.city-item.active {
  background: #ffffff;
  box-shadow: inset 0 0 0 1px #deded7;
}

.city-item.pending {
  opacity: 0.6;
  cursor: default;
}

.city-info {
  display: flex;
  flex-direction: column;
}

.city-name {
  font-size: 15px;
  font-weight: 500;
  color: #202123;
}

.city-adm {
  font-size: 12px;
  color: #6e6e80;
  margin-top: 3px;
}

.remove-btn {
  background: none;
  border: none;
  color: #8e8ea0;
  font-size: 20px;
  cursor: pointer;
  padding: 4px;
  border-radius: 6px;
  opacity: 0;
  transition: opacity 0.15s;
  line-height: 1;
}

.city-item:hover .remove-btn {
  opacity: 1;
}

.remove-btn:hover {
  color: #b42318;
  background: #f8e7e5;
}

.error-msg {
  padding: 10px 20px;
  background: #f8e7e5;
  color: #b42318;
  font-size: 13px;
  text-align: center;
}

@media (max-width: 760px) {
  .sidebar {
    width: 100%;
    min-width: 0;
    height: auto;
    max-height: 40vh;
    border-right: none;
    border-bottom: 1px solid #deded7;
  }
}
</style>
