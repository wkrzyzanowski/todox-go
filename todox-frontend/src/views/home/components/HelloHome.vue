<template>
  <div class="hello">Hello world!</div>
</template>

<script>
import HelloRepository from '@/common/services/repositories/HelloRepository';

export default {
  name: 'HelloHome',
  props: {},
  mounted() {
    HelloRepository.getHello().then(
      (response) => console.log(response),
      (err) => console.error(err.message)
    );
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss"></style>
