<template>
  <div class="login">
    <h2>{{ isLogin ? 'Login' : 'Signup' }}</h2>
    <form @submit.prevent="isLogin ? login() : signUp()">
      <div>
        <label for="username">Username:</label>
        <input type="text" v-model="username" id="username" required>
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" v-model="password" id="password" required>
      </div>
      <button type="submit">{{ isLogin ? 'Login' : 'Signup' }}</button>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import axios from 'axios';
import { setUserId } from '../userId';

interface UserData {
  username: string;
  password: string;
}

const route = useRoute();
const router = useRouter();

const isLogin = ref(route.params.isLogin === 'true');

const username = ref('');
const password = ref('');

const signUp = async () => {
  const userData: UserData = {
    user_name: username.value,
    password: password.value
  };
  try {
    const response = await axios.post('http://localhost:8080/signup', userData);
    const userId: number = response.data;
    setUserId(userId);
    router.push('/');
  } catch (error) {
    console.error('Error signing up:', error);
  }
};

const login = async () => {
  try {
    const userData: UserData = {
      user_name: username.value,
      password: password.value
    };
    const response = await axios.post('http://localhost:8080/signin', userData);
    const userId: number = response.data;
    setUserId(userId);
    router.push('/');
  } catch (error) {
    console.error('Error signing in:', error);
  }
};
</script>

<style scoped>
.login {
  max-width: 300px;
  margin: 0 auto;
  padding: 1em;
  border: 1px solid #ccc;
  border-radius: 1em;
}
.login div {
  margin-bottom: 1em;
}
.login label {
  margin-bottom: 0.5em;
  display: block;
}
.login input {
  border: 1px solid #CCCCCC;
  padding: 0.5em;
  font-size: 1em;
  width: 100%;
  box-sizing: border-box;
}
.login button {
  padding: 0.7em;
  color: #fff;
  background-color: #007BFF;
  border-radius: 5px;
  cursor: pointer;
  width: 100%;
  margin-bottom: 1em;
}
.login button:hover {
  background-color: #0056b3;
}
</style>
