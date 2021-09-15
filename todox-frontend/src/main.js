import { createApp } from 'vue';
import App from './App.vue';
import router from './common/router';
import store from './common/store';
import VueAxios from 'vue-axios';
import axios from 'axios';

createApp(App).use(store).use(router).use(VueAxios, axios).mount('#app');
