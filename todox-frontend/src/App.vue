<template>
  <v-app>
    <v-main>
      <v-container>
        <v-row class="text-center">
          <v-col cols="12">
            <router-link to="/">Home</router-link>
            |
            <router-link to="/about">About</router-link>
          </v-col>
        </v-row>
      </v-container>
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'App',

  data: () => ({}),
});
</script>
