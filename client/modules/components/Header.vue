<template>
  <nav>
    <section class="navbar-logo">
      <ul>
        <li>{{ title }}</li>
      </ul>
    </section>
    <section class="navbar-links">
      <ul>
        <li>
          <a href="#">Card Editor</a>
        </li>
        <li>
          <a href="#">Card Collection</a>
        </li>
      </ul>
    </section>
  </nav>
</template>

<script>
  export default {
    name: 'Header',
    props: {
      title: {
        type: String,
        default: "Jobine TCG"
      }
    },
  }
</script>

<style scoped>
  @import url('https://fonts.googleapis.com/css2?family=Roboto:wght@500&display=swap');

  ul {
    list-style-type: none;
    margin: 0;
    padding: 0;
    height: 100%;
  }

  nav {
    display: flex;
    width: 100%;
    height: 50px;
  }

  a {
    display: flex;
    justify-content: center;
    text-decoration: none;
    color: black;
    font-family: 'Roboto', sans-serif;
    font-size: 20px;
  }

  .navbar-logo {
    display: flex;
    justify-content: center;
    width: 15%;
    border-right: 2px solid black;
  }

  .navbar-logo > ul > li {
    display: flex;
    flex-direction: column;
    justify-content: center;
    font-family: 'Roboto', sans-serif;
    font-size: 20px;
    height: 100%;
  }

  .navbar-links {
    display: flex;
    width: 85%;
  }

  .navbar-links > ul {
    display: flex;
    justify-content: start;
    width: 100%;
  }

  .navbar-links > ul > li {
    display: flex;
    flex-direction: column;
    justify-content: center;
    width: 20%;
    height: 100%;
  }

  .navbar-links > ul > li:hover {
    background-color: hsl(165, 100%, 40%);
    cursor: pointer;
  }

</style>
