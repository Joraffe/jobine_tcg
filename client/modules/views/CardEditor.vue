<template>
  <section class="row">
    <section class="col">
      <Card />
    </section>
  </section>
</template>

<script>
  import Card from 'components/Card.vue';

  export default {
    name: 'Card Editor',
    components: {
      Card,
    }
  }
</script>

<style scoped>
  .row {
    background-color: cornflowerblue;
    height: 100%;
    width: 80%;
    display: flex;
    flex-direction: row;
    justify-content: center;
  }

  .col {
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

</style>
