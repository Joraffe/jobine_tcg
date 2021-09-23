<template>
  <div class="card">
    <div class="card-name-container row">
      <div class="card-type col">
        <span class="card-type-content row">
          <span class="card-type-content col">{{ cardType }}</span>
        </span>
      </div>
      <div class="card-name col">
        <span class="card-name-content row">
          <span class="card-name-content col">{{ name }}</span>
        </span>
      </div>
    </div>
    
    <div class="card-art">{{ artUrl }}</div>
    <div class="card-effect-container">
      <div class="card-effect">{{ effect }}</div>
      <div class="card-stats-container">
        <div class="card-stats">
          <div class="card-power pie-segment">{{ power }}</div>
          <div class="card-intelligence pie-segment">{{ intelligence }}</div>
          <div class="card-endurance pie-segment">{{ endurance }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'Card',
    data() {
      return {
        name: 'Test Card',
        cardType: 'Familiar',
        effect: 'Draw 2 cards',
        power: 1,
        intelligence: 1,
        endurance: 1,
      }
    },
    computed: {
      artUrl() {
        return `/images/${this.cardType}/${this.name}.png`;
      }
    }
  }
</script>

<style scoped>
  @import url('https://fonts.googleapis.com/css2?family=Raleway:wght@800&display=swap');

  .row {
    display: flex;
    flex-direction: row;
    justify-content: center;
  }

  .col {
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  .card {
    display: flex;
    flex-direction: column;

    width: 442.5px;
    height: 645px;
    border: solid 12px darkslategray;
    border-radius: 15px;
    background-color: oldlace;
  }

  .card-name-container {
    height: 10%;
    margin: 2.5%;
    border: double 5px black;

  }

  .card-type {
    width: 20%;
  }

  .card-type-content {
    display: flex;
    justify-content: center;
    background-color: green;
    height: 100%;
  }

  .card-name {
    width: 80%;
  }

  .card-name-content {
    display: flex;
    justify-content: center;
    background-color: red;
    height: 100%;
  }

  .card-art {
    height: 50%;
    margin-top: 2.5%;
    margin-bottom: 2.5%;
    margin-left: 5%;
    margin-right: 5%;
    background-color: powderblue;
    border: double 5px black;
  }

  .card-effect-container {
    display: flex;
    width: 92.5%;
    height: 30%;
    margin: 2.5%;
    background-color: grey;
    border: double 5px black;
  }

  .card-effect {
    width: 70%;
  }

  .card-stats-container {
    width: 25%;
    display: flex;
    flex-direction: column;
    justify-content: end;
    margin: 2.5%;
  }

  .card-stats {
    display: inline-block;
    height: calc(var(--size, 100) * 1px);
    width: calc(var(--size, 100) * 1px);
    border-radius: 50%;
    position: relative;

    border: .15em solid black;
    box-shadow: 0 .075em .2em .05em rgba(0,0,0,.25);
  
    background-origin: border-box;
    background-image: conic-gradient(
      salmon 120deg,      /* 20%               20  * 3.6 =  72 degrees */
      cornflowerblue 0 240deg,   /* 45%         (45 + 20) * 3.6 = 234 degrees */
      darkseagreen 0           /* 35%       last stop is always 360 degrees */
    );
  }

  .card-power {
    font-family: 'Raleway', sans-serif;
    color: white;
    font-size: 1.75em;
    text-shadow: 2px 2px 0 black;

    position: absolute;
    transform: translate(2.3em, .6em);
  }

  .card-intelligence {
    font-family: 'Raleway', sans-serif;
    color: white;
    font-size: 1.75em;
    text-shadow: 2px 2px 0 black;

    position: absolute;
    transform: translate(1.5em, 2em);
  }

  .card-endurance {
    font-family: 'Raleway', sans-serif;
    color: white;
    font-size: 1.75em;
    text-shadow: 2px 2px 0 black;

    position: absolute;
    transform: translate(.8em, .6em);
  }
</style>
