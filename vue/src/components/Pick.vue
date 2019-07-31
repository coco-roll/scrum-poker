<template>
<div>
  <div v-for="item in data" class="card" @click="restart">
  {{item.key}} | {{item.value}}
  </div>
</div>
</template>

<script>
import {restart} from '../js/pick.js'
export default {
  name: 'Pick',
  created () {
    let ret = JSON.parse(this.$route.query.ret)
    let result = []
    for (let i in ret){
      result.push({
        key: i,
        value: ret[i]
      })
    }
    this.data = result
    console.log(this.data)
  },
  methods: {
    restart
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .card {
    background-color: #FF6800;
    width:100%;
    height:100%;
    font-size: 4cm;
    text-align: center;
    line-height:100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }
</style>
