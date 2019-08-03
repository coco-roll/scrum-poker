<template>
 <table class="basic-table" @click="restart">
    <thead>
      <tr>
        <th>牌值</th>
        <th>票数</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in data">
        <td>{{item.key}}</td>
        <td>{{item.value}}</td>
      </tr>
        <tr>
          <td>平均值</td>
          <td>{{this.avg}}</td>
        </tr>
    </tbody>
 </table>
</template>

<script>
import {restart} from '../js/pick.js'
export default {
  name: 'Pick',
  created () {
    let ret = JSON.parse(this.$route.query.ret)
    console.log(ret)
    let result = [], allNum = 0, count = 0
    for (let i in ret){
      result.push({
        key: i,
        value: ret[i]
      })

      allNum = i * ret[i]
      count = count + ret[i]
    }
    this.data = result
    this.avg = allNum / count
    console.log(this.data, this.avg)
  },
  methods: {
    restart
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  table {
    font-size: 1.0cm;
    width:100%;
    background-color: #FF6800;
  }
  table,td,th {
       border: 1px solid black;
  }
</style>
