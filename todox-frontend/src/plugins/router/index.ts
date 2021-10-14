import Vue from 'vue';
import VueRouter, { RouteConfig } from 'vue-router';
import HomeView from '../../views/HomeView.vue';
import LoginView from '../../views/LoginView.vue';

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: HomeView,
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginView,
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../../views/AboutView.vue'),
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
