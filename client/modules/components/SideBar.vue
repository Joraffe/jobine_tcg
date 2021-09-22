<template>
  <aside>
    <section>
      {{ text }}
    </section>
  </aside>
</template>


<script>
  export default {
    name: 'SideBar',
    props: {
      text: {
        type: String,
        default: "Hello Cae",
      },
    },
  }
</script>

<style scoped>
  aside {
    width: 20%;
    height: 100%;
    border-right: solid 2px black;
  }
</style>
