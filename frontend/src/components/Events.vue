<script lang="ts" setup>
import {computed, onMounted, reactive} from 'vue';
import {CalculateElapsedTime} from '../../wailsjs/go/main/App';

interface Event {
  name: string;
  time: Date;
  elapsedTime?: string;
}

const state = reactive<{
  events: Event[];
  loading: boolean;
}>({
  events: [
    { name: 'Zach', time: new Date('1993-09-04T14:16:00Z') },
    { name: 'Olivia', time: new Date('1995-05-29T03:59:00Z') },
    { name: 'Elijah', time: new Date('2024-07-27T20:12:00Z') },
    { name: 'Wedding', time: new Date('2023-08-26T15:00:00Z') },
    { name: 'Now', time: new Date() },
  ],
  loading: false,
});

async function fetchElapsedTimes() {
  state.loading = true;
  for (const event of state.events) {
    try {
      event.elapsedTime = await CalculateElapsedTime(event.time.toISOString());
    } catch (error) {
      event.elapsedTime = 'Error calculating time';
      console.error(error);
    }
  }
  state.loading = false;
}

onMounted(fetchElapsedTimes);

function refreshTimes() {
  fetchElapsedTimes();
}

// Computed property to get formatted UTC and local times for display
const formattedEvents = computed(() =>
    state.events.map(event => ({
      ...event,
      utcTime: event.time.toISOString(), // ISO format for UTC
      localTime: event.time.toLocaleString(), // Local time based on browser settings
    }))
);
</script>

<template>
  <div>
    <h1>Event Tracker</h1>
    <table class="event-table">
      <thead>
      <tr>
        <th>Event</th>
        <th>UTC Time</th>
        <th>Local Time</th>
        <th>Elapsed Time</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="(event, index) in formattedEvents" :key="index">
        <td>{{ event.name }}</td>
        <td>{{ event.utcTime }}</td>
        <td>{{ event.localTime }}</td>
        <td>
          <span v-if="event.elapsedTime">{{ event.elapsedTime }}</span>
          <span v-else>Loading...</span>
        </td>
      </tr>
      </tbody>
    </table>
    <button class="refresh-button" @click="refreshTimes" :disabled="state.loading">
      {{ state.loading ? 'Refreshing...' : 'Refresh Times' }}
    </button>
  </div>
</template>

<style scoped>
.event-table {
  width: 100%;
  border-collapse: collapse;
  margin: 20px 0;
  text-align: center;
}

.event-table th,
.event-table td {
  border: 1px solid #ddd;
  padding: 12px;
}

.event-table th {
  //background-color: #f4f4f4;
  font-weight: bold;
}

.event-table td {
  vertical-align: middle;
}

.event-table tr:nth-child(even) {
  //background-color: #f9f9f9;
}

.event-table tr:hover {
  background-color: #a81616;
}

.refresh-button {
  display: block;
  margin: 20px auto;
  padding: 10px 20px;
  font-size: 16px;
  color: #fff;
  background-color: #007bff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.refresh-button:disabled {
  background-color: #a1a1a1;
  cursor: not-allowed;
}

.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
