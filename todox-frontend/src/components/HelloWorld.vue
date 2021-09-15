<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          :src="require('../assets/logo.svg')"
          class="my-3"
          contain
          height="200"
        />
        <span>
          {{ message }}
        </span>
      </v-col>
    </v-row>
    <v-row class="text-center">
      <v-col cols="12">
        <v-btn @click="makeCall()">Hello from API!</v-btn>
        |
        <v-btn @click="clearMessage()">Clear API!</v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import Vue from 'vue';
import HelloRepository from '../repository/HelloRepository';

export default Vue.extend({
  name: 'HelloWorld',
  methods: {
    makeCall(): void {
      HelloRepository.getHello()
        .then((response) => (this.message = response.data))
        .catch((err) => {
          console.error(err);
        });
    },
    clearMessage(): void {
      this.message = '';
    },
  },
  data: () => ({
    message: '',
  }),
});
</script>
