<script setup>
import { reset } from '@formkit/core'

import { AxiosCall } from '../api'
import { Store } from '../store'

defineProps({
  title: {
    type: String,
    required: true
  }
})

/*
 * Convert a time string of HH:MM to number of seconds.
 *
 * @param t {string} - a string of format "00:00"
 * @return number
 */
const timeToMinutes = (t) => {
  const [h, m] = t.split(':').map((x) => parseInt(x))
  return h * 60 + m
}

const validate = (data) => {
  const b = timeToMinutes(data.time.begin)
  const e = timeToMinutes(data.time.end)
  return {
    entry: parseInt(data.entry),
    sex: data.sex,
    activity: data.activity,
    time_begin: b,
    time_end: e,
    time_break: data.time.break,
    time_total: e - b,
    location_begin: data.location.begin,
    location_end: data.location.end
  }
}

function submitObservation(data) {
  const validatedData = validate(data)
  AxiosCall('post', '/entry', validatedData).then(() => {
    Store.add(validate(data))
    reset('entryForm')
  })
}
</script>

<template>
  <FormKit
    type="form"
    id="entryForm"
    submit-label="Create"
    form-class="p-2 border rounded"
    :submit-attrs="{ inputClass: 'btn btn-primary' }"
    @submit="submitObservation"
  >
    <h3 class="mb-3">{{ title }}</h3>
    <div class="row">
      <FormKit
        type="number"
        name="entry"
        id="entry"
        validation="required"
        label="Entry"
        outer-class="col-4 mb-3"
        input-class="form-control"
      />
      <FormKit
        type="text"
        name="activity"
        id="activity"
        validation="required"
        label="Activity"
        outer-class="col-4 mb-3"
        input-class="form-control"
      />
      <FormKit
        type="text"
        name="sex"
        id="sex"
        validation="required"
        label="Sex"
        outer-class="col-4 mb-3"
        input-class="form-control"
      />
    </div>
    <section>
      <h4 class="mb-4">Time</h4>
      <div class="row">
        <FormKit type="group" name="time" id="time">
          <FormKit
            type="time"
            name="begin"
            id="timeBegin"
            validation="required"
            label="Begin"
            outer-class="col-5 mb-3"
            input-class="form-control"
          />
          <FormKit
            type="time"
            name="end"
            id="timeEnd"
            validation="required"
            label="End"
            outer-class="col-5 mb-3"
            input-class="form-control"
          />
          <FormKit
            type="text"
            name="break"
            id="timeBreak"
            validation="required"
            label="Break"
            outer-class="col-2 mb-3"
            input-class="form-control"
          />
        </FormKit>
      </div>
    </section>
    <section>
      <h4 class="mb-4">Location</h4>
      <div class="row">
        <FormKit type="group" name="location" id="location">
          <FormKit
            type="text"
            name="begin"
            id="locationBegin"
            validation="required"
            label="Begin"
            outer-class="col-6 mb-3"
            input-class="form-control"
          />
          <FormKit
            type="text"
            name="end"
            id="locationEnd"
            validation="required"
            label="End"
            outer-class="col-6 mb-3"
            input-class="form-control"
          />
        </FormKit>
      </div>
    </section>
  </FormKit>
</template>

<style scoped></style>
