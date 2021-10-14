<template>
  <v-container>
    <v-snackbar
      v-for="message in this.messageQueue"
      :key="message"
      timeout="2000"
    >
      AAA!
    </v-snackbar>
  </v-container>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'AlertPopupComponent',

  data() {
    return {
      messageQueue: [],
    };
  },

  methods: {},

  components: {},
});
</script>

<style lang="scss">
.position {
  position: absolute;
}
</style>
