<template>
    <div class="home">
      <h1>Welcome to the Homepage</h1>
      <button @click="goToSignup">Sign Up</button>
      <button @click="goToSignin">Sign In</button>
      <button @click="createNote">Create Note</button>
      <button @click="viewMyNotes">View My Notes</button>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Home',
    methods: {
      goToSignup() {
        console.log('Navigating to Sign Up');
      },
      goToSignin() {
        console.log('Navigating to Sign In');
      },
      createNote() {
        console.log('Creating a Note');
      },
      viewMyNotes() {
        console.log('view notes')
      }
    }
  }
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
  