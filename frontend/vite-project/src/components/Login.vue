<template>
  <div class="login">
    <h2>{{ isLogin ? 'Login' : 'Signup' }}</h2>
    <form @submit.prevent="isLogin ? login() : signup()">
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

<script>
export default {
  props: {
    isLogin: {
      type: Boolean,
      required: true
    }
  },
  data() {
    return {
      username: '',
      password: ''
    };
  },
  methods: {
    login() {
      console.log('Logging in with:', this.username, this.password);
    },
    signup() {
      console.log('Signing up with:', this.username, this.password);
    }
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
  margin-bottom: .5em;
  display: block;
}
.login input {
  border: 1px solid #CCCCCC;
  padding: .5em;
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
