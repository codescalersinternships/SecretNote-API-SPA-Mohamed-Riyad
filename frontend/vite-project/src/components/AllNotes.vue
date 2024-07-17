<template>
  <div class="notes-list">
    <h2>All Notes</h2>
    <ul>
      <li v-for="note in notes" :key="note.ID">
        <router-link :to="`/viewnote/${note.ID}`">{{ note.title }}</router-link>
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { notes } from '../notes'; 

</script>

<style scoped>
.notes-list {
  max-width: 600px;
  margin: auto;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  margin-bottom: 10px;
}

router-link {
  color: #007bff;
  text-decoration: none;
}

router-link:hover {
  text-decoration: underline;
}
</style>
