import { createRouter, createWebHistory } from 'vue-router';
import Home from '../../views/home/Home.vue';
import About from '@/views/about/About';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/about',
    name: 'About',
    component: About,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
