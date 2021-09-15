<template>
  <div class="home">
    <img alt="Vue logo" :src="require('../../assets/logo.png')" />
    <HelloWorld msg="Welcome to Your Vue.js App" />
  </div>
</template>

<script>
import HelloWorld from '@/views/home/components/HelloHome.vue';

export default {
  name: 'Home',
  components: {
    HelloWorld,
  },
};
</script>
