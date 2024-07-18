<template>
  <div class="create-note">
    <h2>Create Note</h2>
    <form @submit.prevent="createNote">
      <div>
        <label for="title">Title:</label>
        <input type="text" v-model="title" id="title" required>
      </div>
      <div>
        <label for="content">Content:</label>
        <textarea v-model="content" id="content" class="large-textarea" required></textarea>
      </div>
      <button type="submit">Create Note</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';
import { token } from '../token'; 

const router = useRouter();

const title = ref('');
const content = ref('');

const createNote = async () => {
  const noteData = {
    title: title.value,
    content: content.value,
  };

  try {
    await axios.post('http://localhost:8080/create-note', noteData, {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    });
    router.push('/');
  } catch (error) {
    console.error('Error creating note:', error);
  }
};
</script>

<style scoped>
.create-note {
  max-width: 800px;
  width: 600px;
  margin: 0 auto;
  padding: 1em;
  border: 1px solid #ccc;
  border-radius: 1em;
}
.create-note div {
  margin-bottom: 1em;
}
.create-note label {
  margin-bottom: .5em;
  color: #333333;
  display: block;
}
.create-note textarea {
  height: 300px;
}
.create-note input,
.create-note textarea {
  border: 1px solid #CCCCCC;
  padding: .5em;
  font-size: 1em;
  width: 100%;
  box-sizing: border-box;
}
.create-note button {
  padding: 0.7em;
  color: #fff;
  background-color: #007BFF;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  width: 100%;
}
.create-note button:hover {
  background-color: #0056b3;
}
</style>
