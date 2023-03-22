<script setup>
import EntryForm from './components/EntryForm.vue'
import DataTable from './components/DataTable.vue'
</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col-4">
        <EntryForm title="Observation" />
      </div>
      <div class="col-8 border rounded rounded-2 ">
        <DataTable />
      </div>
    </div>
  </div>
</template>

<style scoped></style>
