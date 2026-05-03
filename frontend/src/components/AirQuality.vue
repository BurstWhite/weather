<template>
  <div class="air-quality">
    <div class="aqi-display">
      <div class="aqi-circle" :style="{ background: aqiColor }">
        <span class="aqi-value">{{ air.aqi }}</span>
        <span class="aqi-label">AQI</span>
      </div>
      <div class="aqi-info">
        <span class="category" :style="{ color: aqiColor }">{{ air.category }}</span>
        <span class="pollutant" v-if="air.primaryPollutant">
          Primary: {{ air.primaryPollutant }}
        </span>
      </div>
    </div>
    <div class="pollutants">
      <div class="pollutant-item" v-if="air.pm2p5 != null">
        <span class="poll-name">PM2.5</span>
        <span class="poll-value">{{ air.pm2p5.toFixed(1) }} μg/m³</span>
      </div>
      <div class="pollutant-item" v-if="air.pm10 != null">
        <span class="poll-name">PM10</span>
        <span class="poll-value">{{ air.pm10.toFixed(1) }} μg/m³</span>
      </div>
      <div class="pollutant-item" v-if="air.no2 != null">
        <span class="poll-name">NO₂</span>
        <span class="poll-value">{{ air.no2.toFixed(1) }} ppb</span>
      </div>
      <div class="pollutant-item" v-if="air.o3 != null">
        <span class="poll-name">O₃</span>
        <span class="poll-value">{{ air.o3.toFixed(1) }} ppb</span>
      </div>
      <div class="pollutant-item" v-if="air.co != null">
        <span class="poll-name">CO</span>
        <span class="poll-value">{{ air.co.toFixed(2) }} ppm</span>
      </div>
      <div class="pollutant-item" v-if="air.so2 != null">
        <span class="poll-name">SO₂</span>
        <span class="poll-value">{{ air.so2.toFixed(1) }} ppb</span>
      </div>
    </div>
    <div class="health" v-if="air.healthAdvice">
      <p>{{ air.healthAdvice }}</p>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  air: { type: Object, required: true },
})

const aqiColor = computed(() => {
  const v = props.air.aqi
  if (v <= 50) return '#00e400'
  if (v <= 100) return '#ffff00'
  if (v <= 150) return '#ff7e00'
  if (v <= 200) return '#ff0000'
  if (v <= 300) return '#99004c'
  return '#7e0023'
})
</script>

<style scoped>
.air-quality {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.aqi-display {
  display: flex;
  align-items: center;
  gap: 20px;
}

.aqi-circle {
  width: 96px;
  height: 96px;
  border-radius: 50%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: #202123;
  box-shadow: inset 0 0 0 1px rgba(32, 33, 35, 0.12);
}

.aqi-value {
  font-size: 32px;
  font-weight: 700;
  line-height: 1;
}

.aqi-label {
  font-size: 12px;
  font-weight: 600;
  margin-top: 4px;
}

.aqi-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.category {
  font-size: 20px;
  font-weight: 600;
}

.pollutant {
  font-size: 14px;
  color: #6e6e80;
}

.pollutants {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
  gap: 10px;
}

.pollutant-item {
  background: #f7f7f4;
  border: 1px solid #ecece6;
  border-radius: 8px;
  padding: 12px;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.poll-name {
  font-size: 12px;
  color: #6e6e80;
  text-transform: uppercase;
}

.poll-value {
  font-size: 16px;
  font-weight: 600;
  color: #202123;
}

.health {
  background: #f7f7f4;
  border: 1px solid #ecece6;
  border-radius: 8px;
  padding: 16px;
  font-size: 14px;
  color: #353740;
  line-height: 1.5;
}

@media (max-width: 560px) {
  .aqi-display {
    align-items: flex-start;
  }

  .aqi-circle {
    width: 84px;
    height: 84px;
  }

  .aqi-value {
    font-size: 28px;
  }
}
</style>
