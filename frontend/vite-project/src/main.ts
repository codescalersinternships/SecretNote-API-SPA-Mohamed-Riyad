import { createApp } from 'vue';
import App from './App.vue';
// @ts-ignore
import router from './router';
import './style.css'


const app = createApp(App);

app.use(router).mount('#app');