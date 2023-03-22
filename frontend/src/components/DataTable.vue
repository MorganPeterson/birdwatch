<script setup>
import { onMounted } from 'vue'

import { AxiosCall } from '../api'
import { Store } from '../store'

const fields = [
  'entry',
  'time_begin',
  'time_end',
  'time_total',
  'time_break',
  'sex',
  'activity',
  'location_begin',
  'location_end'
]

onMounted(() => {
  AxiosCall('get').then((response) => {
    if (response.error === null) {
      Store.load(response.data)
    }
  })
})
</script>

<template>
  <table id="tableComponent" class="table table-bordered table-striped">
    <thead>
      <tr>
        <th rowspan="2">Entry</th>
        <th colspan="4">Time</th>
        <th rowspan="2">Sex</th>
        <th rowspan="2">Activity</th>
        <th colspan="2">Location</th>
      </tr>
      <tr>
        <th>Beginning</th>
        <th>End</th>
        <th>Total</th>
        <th>Break</th>
        <th>Beginning</th>
        <th>End</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in Store.data" :key="item">
        <td v-for="field in fields" :key="field">
          {{ item[field] }}
        </td>
      </tr>
    </tbody>
  </table>
</template>

<style></style>
