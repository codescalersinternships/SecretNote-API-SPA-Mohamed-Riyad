import { createRouter, createWebHistory } from 'vue-router';
import Home from '../components/Home.vue';
import CreateNote from '../components/CreateNote.vue';
import Login from '../components/Login.vue';
import ViewNote from '../components/ViewNote.vue';
import AllNotes from '../components/AllNotes.vue';

const routes = [
  { path: '/', component: Home, name: 'home' },
  { path: '/createnote', component: CreateNote, name: 'createnote' },
  { path: '/login', component: Login, name: 'signin' },
  { path: '/signup', component: Login, name: 'signup' },
  { path: '/viewnote/:id', component: ViewNote, name: 'viewnote' },
  { path: '/allnotes', component: AllNotes, name: 'allnotes' },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
