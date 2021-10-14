<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <login-form-component />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import Vue from 'vue';
import LoginFormComponent from '../components/login-form/LoginFormComponent.vue';

export default Vue.extend({
  name: 'LoginView',

  components: { LoginFormComponent },
});
</script>
