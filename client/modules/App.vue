<template>
  <Header title="Jobine TCG"/>
</template>

<script>
  import Header from 'components/Header.vue';

  export default {
    name: 'App',
    components: {
      Header,
    },
  }
</script>

<style>
  body {
    margin: 0;
  }

  #app {
    border-top: 2px solid black;
    border-bottom: 2px solid black;
  }
</style>
