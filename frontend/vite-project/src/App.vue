<script setup lang="ts">
import ViewNote from './components/ViewNote.vue';


const dummyNote =  { id: 1, title: 'First Note' }
</script>

<template>
  <ViewNote :note="dummyNote" />
</template>
