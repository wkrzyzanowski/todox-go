<template>
  <v-form ref="form">
    <v-text-field v-model="username" label="Username" />
    <v-text-field v-model="password" label="Password" type="password" />
    <v-btn @click="login()">Login</v-btn>
  </v-form>
</template>

<script lang="ts">
import Vue from 'vue';
import AuthenticationRestService from '../../service/AuthenticationRestService';

export default Vue.extend({
  name: 'LoginFormComponent',

  data() {
    return {
      username: '',
      password: '',
    };
  },

  methods: {
    login() {
      AuthenticationRestService.login(this.username, this.password)
        .then((response) => {
          console.log(response);
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },

  components: {},
});
</script>
