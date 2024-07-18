<template>
  <div class="home">
    <h1>Welcome to the Homepage</h1>
    <button @click="goToSignup">Sign Up</button>
    <button @click="goToSignin">Sign In</button>
    <button @click="createNote">Create Note</button>
    <button @click="viewMyNotes">View My Notes</button>
  </div>
</template>

<script setup lang="ts">

import { useRouter } from 'vue-router';
import { token } from '../token'; 
import { setNotes } from '../notes'; 
import { getAllNotes } from '../api';

const router = useRouter();

const goToSignup = () => {
  router.push({ name: 'signup', params: { isLogin: 'false' } });
};

const goToSignin = () => {
  router.push({ name: 'signin', params: { isLogin: 'true' } });
};

const createNote = () => {
  if (token.value!==null) {
    router.push({ name: 'createnote' });
  } else {
    console.error('token is not available');
  }
};

const  viewMyNotes = async() => {
  if (token.value!==null) {
    const notes: any = await getAllNotes(token.value)
    setNotes(notes)
    router.push({ name: 'allnotes'});
  } else {
    console.error('token is not available');
  }
};
</script>

<style scoped>
.home {
  text-align: center;
  margin-top: 50px;
}

button {
  margin: 10px;
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
}
</style>
