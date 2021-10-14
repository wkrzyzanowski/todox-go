<template>
  <v-app>
    <v-main>
      <alert-popup-component />
      <v-container>
        <v-row class="text-center">
          <v-col cols="12">
            <router-link to="/">Home</router-link>
            |
            <router-link to="/about">About</router-link>
            |
            <router-link to="/login">Login</router-link>
          </v-col>
        </v-row>
      </v-container>
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import AlertPopupComponent from './components/alert-popup/AlertPopupComponent.vue';

export default Vue.extend({
  components: { AlertPopupComponent },
  name: 'App',

  data: () => ({}),
});
</script>
