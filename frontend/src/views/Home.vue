<template>
  <div class="app-container">
    <CitySidebar />
    <main class="main-content">
      <div v-if="store.loading" class="loading">Loading...</div>
      <div v-else-if="store.error" class="error">{{ store.error }}</div>
      <div v-else-if="!store.selectedCity" class="placeholder">
        <h2>Select a city to view weather</h2>
      </div>
      <template v-else-if="store.weatherData">
        <header class="city-header">
          <h1>{{ store.weatherData.city.name }}</h1>
          <span class="adm">{{ store.weatherData.city.adm1 }}</span>
        </header>
        <section class="weather-section">
          <h2>7-Day Forecast</h2>
          <Weather7Days :weather="store.weatherData.weather" />
        </section>
        <section class="humidity-section" v-if="store.weatherData.weather">
          <h2>Humidity</h2>
          <HumidityBar :weather="store.weatherData.weather" />
        </section>
        <section class="air-section" v-if="store.weatherData.airQuality">
          <h2>Air Quality</h2>
          <AirQuality :air="store.weatherData.airQuality" />
        </section>
      </template>
    </main>
  </div>
</template>

<script setup>
import { useWeatherStore } from '../stores/weather'
import CitySidebar from '../components/CitySidebar.vue'
import Weather7Days from '../components/Weather7Days.vue'
import HumidityBar from '../components/HumidityBar.vue'
import AirQuality from '../components/AirQuality.vue'

const store = useWeatherStore()
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: ui-sans-serif, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;
  background: #f7f7f4;
  color: #202123;
  min-height: 100vh;
  text-rendering: optimizeLegibility;
  -webkit-font-smoothing: antialiased;
}

.app-container {
  display: flex;
  min-height: 100vh;
}

.main-content {
  flex: 1;
  padding: 40px 56px;
  overflow-y: auto;
  background: linear-gradient(180deg, #ffffff 0%, #f7f7f4 100%);
}

.city-header {
  margin: 0 auto 24px;
  max-width: 1040px;
}

.city-header h1 {
  font-size: 32px;
  font-weight: 650;
  color: #202123;
  letter-spacing: 0;
}

.city-header .adm {
  font-size: 14px;
  color: #6e6e80;
  margin-top: 6px;
  display: inline-block;
}

section {
  margin: 0 auto 18px;
  max-width: 1040px;
  background: #ffffff;
  border: 1px solid #e5e5e0;
  border-radius: 8px;
  padding: 22px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.04);
}

section h2 {
  font-size: 14px;
  font-weight: 650;
  margin-bottom: 16px;
  color: #353740;
  text-transform: none;
  letter-spacing: 0;
}

.placeholder {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 60vh;
  color: #6e6e80;
  text-align: center;
}

.loading, .error {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 40vh;
  font-size: 15px;
}

.error {
  color: #b42318;
}

@media (max-width: 760px) {
  .app-container {
    flex-direction: column;
  }

  .main-content {
    padding: 24px 16px;
  }

  .city-header h1 {
    font-size: 26px;
  }
}
</style>
