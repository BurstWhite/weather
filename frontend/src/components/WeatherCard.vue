<template>
  <div class="weather-card">
    <div class="day-name">{{ dayName }}</div>
    <div class="date">{{ dateStr }}</div>
    <img :src="`https://cdn.jsdelivr.net/npm/qweather-icons@1.8.0/icons/${day.iconDay}.svg`" class="weather-icon" />
    <div class="temps">
      <span class="high">{{ day.tempMax }}°</span>
      <span class="low">{{ day.tempMin }}°</span>
    </div>
    <div class="text">{{ day.textDay }}</div>
  </div>
</template>

<script setup>
const props = defineProps({
  day: { type: Object, required: true },
})

const weekDays = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']

const dateObj = new Date(props.day.fxDate)
const dayName = weekDays[dateObj.getDay()]
const dateStr = `${dateObj.getMonth() + 1}/${dateObj.getDate()}`
</script>

<style scoped>
.weather-card {
  background: #f7f7f4;
  border: 1px solid #ecece6;
  border-radius: 8px;
  padding: 16px;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
  min-height: 172px;
  transition: border-color 0.15s, box-shadow 0.15s, transform 0.15s;
}

.weather-card:hover {
  border-color: #d9d9d0;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.07);
  transform: translateY(-1px);
}

.day-name {
  font-size: 14px;
  font-weight: 600;
  color: #202123;
}

.date {
  font-size: 12px;
  color: #6e6e80;
}

.weather-icon {
  width: 48px;
  height: 48px;
}

.temps {
  display: flex;
  gap: 8px;
  font-size: 16px;
}

.high {
  color: #202123;
  font-weight: 600;
}

.low {
  color: #8e8ea0;
}

.text {
  font-size: 12px;
  color: #6e6e80;
}
</style>
