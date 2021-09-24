<template>
  <div class="card">
    <div class="card-name-container row">
      <div class="card-type col">
        <span class="card-type-content row">
          <span class="card-type-content col">
            <img v-bind:src="typeIcon" />
            <span class="row">{{ cardType }}</span>
          </span>
        </span>
      </div>
      <div class="card-name col">
        <span class="card-name-content row">
          <span class="card-name-content col">{{ name }}</span>
        </span>
      </div>
    </div>
    
    <div class="card-art-container">
      <div class="card-art">
        <img v-bind:src="artUrl" />
      </div>
    </div>

    <div class="card-info-container">
      <div class="card-effect-container">
        <div class="card-effect">{{ effect }}</div>
      </div>
      <div class="card-stats-container">
        <div class="card-stats">
          <div class="card-power pie-segment">{{ power }}</div>
          <div class="card-intelligence pie-segment">{{ intelligence }}</div>
          <div class="card-endurance pie-segment">{{ endurance }}</div>
        </div>
      </div>
    </div>

    <div class="card-extras-container">
      <div class="card-credit">
        <div class="card-artist">
          <img src="/static/images/icons/pen_icon.png"/>  {{ artist }}
        </div>
        <div class="card-copyright">©{{ copyright }}</div>
      </div>
      <div class="card-metadata">
        <div class="card-number">{{ paddedCardNumber }}/{{ cardSetTotalNum }}</div>
        <div class="card-rarity">{{ rarity }}</div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'Card',
    data() {
      return {
        name: 'Petrel',
        cardType: 'familiar',
        effect: 'Draw 2 cards',
        power: 1,
        intelligence: 1,
        endurance: 1,
        rarity: 'R',
        number: 1,
        cardSetTotalNum: 150,
        artist: 'caebae',
        copyright: 'Jobine',
      }
    },
    computed: {
      artUrl() {
        const lower = this.name.toLowerCase()
        const snake_case = lower.replace(' ', '_');
        return `static/images/${this.cardType}/${snake_case}.jpg`;
      },
      typeIcon() {
        return `static/images/icons/${this.cardType}_icon.png`;
      },
      paddedCardNumber() {
        return this.number.toString().padStart(3, '0');
      }
    }
  }
</script>

<style scoped>
  @import url('https://fonts.googleapis.com/css2?family=Raleway:wght@800&display=swap');
  @import url('https://fonts.googleapis.com/css2?family=Nunito:ital,wght@0,400;1,900&display=swap');
  @import url('https://fonts.googleapis.com/css2?family=Carter+One&display=swap');

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
    background-color: rgb(198, 209, 223);
  }

  .card-name-container {
    height: 7.93%;
    margin-top: 2.5%;
    margin-bottom: 1.5%;
    margin-left: 2.5%;
    margin-right: 2.5%;
    border: double 5px black;
    
    position: relative;
  }

  .card-type {
    width: 15%;
    height: 170%;
  
    border-left: solid 3px black;
    border-right: solid 3px black;
    border-bottom: solid 3px black;

    background-color: rgb(156, 179, 211);

    position: absolute;
    transform: translate(-10em, -1.75em);
  }

  .card-type-content {
    display: flex;
    justify-content: center;
    height: 100%;

    font-family: 'Nunito', sans-serif;
    font-size: 1em;
    font-weight: bold;
  }

  .card-name {
    width: 100%;
  }

  .card-name-content {
    display: flex;
    justify-content: center;
    background-color: floralwhite;
    height: 100%;

    font-family: 'Carter One', cursive;
    font-size: 1.5em;
  }

  .card-art-container {
    height: 70%;
    margin-top: 1%;
    margin-bottom: 0.5%;
    margin-left: 5%;
    margin-right: 5%;
    border: double 5px black;
  }

  img {
    max-width: 100%;
    max-height: 100%;
  }

  .card-info-container {
    display: flex;
    width: 92.5%;
    height: 20%;
    margin-top: 2.5%;
    margin-left: 2.5%;
    margin-right: 2.5%;
    background-color: floralwhite;
    border: double 5px black;
  }

  .card-effect-container {
    width: 70%;
    margin: 2.5%;
  }

  .card-effect {
    font-family: 'Nunito', sans-serif;
    font-size: 1em;
    font-style: normal;
    font-weight: bold;
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

  .card-extras-container {
    display: flex;
    justify-content: space-between;
    
    height: 3.5%;

    font-family: 'Nunito', sans-serif;
    font-size: 0.7em;
    font-weight: bold;

    margin-left: 2.5%;
    margin-right: 2.5%
    
  }

  .card-credit {
    display: flex;
    justify-content: space-between;

    width: 25%;
  }

  .card-artist > img {
    height: 0.7em;
  }

  .card-metadata {
    display: flex;
    justify-content: space-evenly;

    width: 15%;
  }
</style>
