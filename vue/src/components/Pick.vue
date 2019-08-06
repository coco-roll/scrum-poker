<template>
 <div class="basic-table">
    <div class="title">翻牌结果：</div>
    <div class="header box">
      <section class="table_box box">
        <span class="box_item">牌值</span>
        <span class="box_item">票数</span>
      </section>
    </div>
    <div class="body">
      <div class="container">
        <section class="body_box box" v-for="item in data" :key="item">
          <span class="box_item">{{item.key}}</span>
          <span class="box_item">|</span>
          <span class="box_item">{{item.value}}</span>
        </section>
      </div>
    </div>
    <div class="footer box">
      <div class="box_m">
        <div class="avg_box">平均值： {{this.avg}}</div>
      </div>
      <div class="box_m">
        <div class="reset_bt" @click="restart">重置</div>
      </div>
    </div>
 </div>
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
.basic-table{
  font-size: 30px;
  width:100%;
  background-color: #97A889;
  height: 100%;
}
.box_item{
  flex: 1;
}
.body_box{
  border-bottom: 1px solid #525252;
}
.box{
  width: 100%;
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.header{
  box-sizing: border-box;
  border-bottom: 1px solid #525252;
}
.body{
  height: calc(100vh - 200px);
  overflow: scroll;
}
.footer{
  position: fixed;
  bottom: 0;
  height: 100px;
  flex-direction: column;
  padding-bottom: 10px;
  background: #FF6800;
  /* border-top: 1px solid grey; */
}
.avg_box{
  height: 50px;
}
.reset_bt{
  width: 300px;
  height: 50px;
  border: 1px solid #97A889;
  background: #97A889;
  border-radius: 5px;
  box-sizing: border-box;
}
.box_m{
  width: 100%;
  display: flex;
  justify-content: center;
}

</style>
