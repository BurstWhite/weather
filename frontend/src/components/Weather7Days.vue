<template>
  <div class="weather-grid">
    <WeatherCard v-for="day in weather" :key="day.fxDate" :day="day" />
  </div>
</template>

<script setup>
import WeatherCard from './WeatherCard.vue'

defineProps({
  weather: { type: Array, required: true },
})
</script>

<style scoped>
.weather-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(132px, 1fr));
  gap: 10px;
}
</style>
