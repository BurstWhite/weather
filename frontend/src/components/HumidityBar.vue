<template>
  <div class="humidity-container">
    <div class="humidity-grid">
      <div v-for="day in weather" :key="day.fxDate" class="humidity-item">
        <div class="hum-day">{{ dayName(day.fxDate) }}</div>
        <div class="hum-bar-wrapper">
          <div
            class="hum-bar"
            :style="{ height: day.humidity + '%' }"
            :title="day.humidity + '%'"
          ></div>
        </div>
        <div class="hum-value">{{ day.humidity }}%</div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  weather: { type: Array, required: true },
})

const weekDays = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']

function dayName(dateStr) {
  const d = new Date(dateStr)
  return weekDays[d.getDay()]
}
</script>

<style scoped>
.humidity-container {
  padding: 8px 0 2px;
}

.humidity-grid {
  display: flex;
  justify-content: space-between;
  gap: 10px;
  align-items: flex-end;
}

.humidity-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
}

.hum-day {
  font-size: 12px;
  color: #6e6e80;
  font-weight: 500;
}

.hum-bar-wrapper {
  width: 32px;
  height: 120px;
  background: #f3f3ef;
  border: 1px solid #ecece6;
  border-radius: 16px;
  display: flex;
  align-items: flex-end;
  padding: 4px;
}

.hum-bar {
  width: 100%;
  background: linear-gradient(to top, #10a37f, #41c6a4);
  border-radius: 12px;
  transition: height 0.3s;
  min-height: 4px;
}

.hum-value {
  font-size: 12px;
  color: #353740;
  font-weight: 600;
}

@media (max-width: 640px) {
  .humidity-grid {
    overflow-x: auto;
    justify-content: flex-start;
    padding-bottom: 4px;
  }

  .humidity-item {
    min-width: 56px;
  }
}
</style>
