import { reactive } from 'vue'

const padTwo = (n) => {
  return n.toString().padStart(2, '0')
}

const convertMinutes = (m) => {
  const hours = Math.floor(m / 60)
  const mins = m % 60

  return `${padTwo(hours)}:${padTwo(mins)}`
}

const Store = reactive({
  data: [],
  add(value) {
    value.time_begin = convertMinutes(value.time_begin)
    value.time_end = convertMinutes(value.time_end)
    this.data.push(value)
  },
  load(response) {
    this.data = response.map((v) => {
      v.time_begin = convertMinutes(v.time_begin)
      v.time_end = convertMinutes(v.time_end)
      return v
    })
  }
})

export { Store }
