<template>
  <div class="view-note">
    <h2>Note Details</h2>
    <div v-if="note">
      <h3>Note Title :{{ note.title }}</h3>
      <p>{{ note.content }}</p>
    </div>
    <div v-else>
      <p>Loading...</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRoute } from 'vue-router';
import { notes } from '../notes'; 


const route = useRoute();
const noteId = ref(parseInt(route.params.id as string, 10));

const note = computed(() => {
  return notes.value.find(n => n.ID === noteId.value);
});

</script>

<style scoped>
.view-note {
  max-width: 600px;
  margin: auto;
  border: 10px;
  border-color: aliceblue;
  border-radius: 2%;
}
</style>
