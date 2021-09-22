<template>
  <Header title="Jobine TCG"/>
  <section>
    <router-view class="view" name="sidebar"></router-view>
    <router-view class="view" name="content"></router-view>
  </section>
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

  section {
    display: flex;
    height: 100%;
  }
</style>
