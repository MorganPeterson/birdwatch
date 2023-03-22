import { reactive } from 'vue'

/* pad number string to match format 00 */
const padTwo = (n) => {
  return n.toString().padStart(2, '0')
}

/* convert minutes to HH:MM */
const convertMinutes = (m) => {
  const hours = Math.floor(m / 60)
  const mins = m % 60

  return `${padTwo(hours)}:${padTwo(mins)}`
}

/* 
 * Store is our data store with helpers.
 *
 * add - adds a single entry into the store to be displayed in the table
 * load - loads all entries from GET query
 */
const Store = reactive({
  data: [],
  add(value) {
    value.time_begin = convertMinutes(value.time_begin)
    value.time_end = convertMinutes(value.time_end)

    const found = this.data.find(entry => entry.entry === value.entry)

    if (found === undefined) {
      this.data.push(value)
    } else {
      this.data[value.entry-1] = {...found, ...value}
    }
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
