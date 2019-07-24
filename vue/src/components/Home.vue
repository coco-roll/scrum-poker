<template>
  <div class="homeback">
    <div>
      <li v-for="num in num_config" :key='num' class="card" :data-id='num' v-on:click="cardClick(num)">
      {{ num == 9999 ? '?' : num }}
      </li>
    </div>
  </div>
</template>

<script>
import {cardClick} from '../js/Home.js'
export default {
  name: 'Home',
  data () {
    return {
      num_config: [1, 2, 3, 4, 6, 9999]
    }
  },
  methods: {
    cardClick
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  h1, h2 {
    font-weight: normal;
  }
  ul {
    list-style-type: none;
    padding: 0;
  }
  li {
    display: inline-block;
  }
  .card {
    background-color: #FF6800;
    border-radius:10px;
    border:1px solid #000;
    width: 30%;
    height:40%;
    font-size: 1.5cm;
    text-align: center;
    line-height:150px;
    margin: 20px;
  }

  .homeback{
    background-color: #97A889;
    height: 100%;
    /*display: flex;*/
    /*justify-content: center;*/
    /*align-items: center;*/
  }
</style>
