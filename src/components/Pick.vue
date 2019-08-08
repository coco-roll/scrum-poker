<template>
 <div class="basic-table">
    <div class="header box">
      <div class="title">翻牌结果</div>
      <section class="table_box box">
        <span class="box_item">牌值</span>
        <span class="box_item_flag">|</span>
        <span class="box_item">票数</span>
      </section>
    </div>
    <div class="body">
      <div class="container">
        <section class="body_box box" v-for="item in data" :key="item">
          <span class="box_item">{{item.key}}</span>
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
  width:100%;
  background-color: #F7F7F7;
  height: 100%;
}
.box_item{
  flex: 1;
  font-size: 24px;
}
.box_item_flag{
  color: #E6E6E6;
  width: 1px;
  height: 24px;
  display: inline-block;
}
.body_box{
  border-bottom: 1px solid #E6E6E6;
}
.box{
  width: 100%;
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #ffffff;
}
.title{
  background: #ffffff;
  height: 50px;
  line-height: 50px;
  color: #222222;
  font-size: 30px;
  font-weight: bold;
  width: 100%;
  position: absolute;
  top: 0;
}
.table_box{
  position: absolute;
  bottom: 0;
}
.header{
  height: 100px;
  box-sizing: border-box;
  position: fixed;
  top: 0;
  box-shadow: 0 0 10pt rgba(0,0,0,0.06);
}
.body{
  height: calc(100vh - 200px);
  overflow: scroll;
  padding-top: 100px;
}
.container{
  padding: 0 20px;
  background: #ffffff;
}
.footer{
  position: fixed;
  bottom: 0;
  height: 100px;
  flex-direction: column;
  padding-bottom: 10px;
  background: #ffffff;
  box-shadow: 0 0 10pt rgba(0,0,0,0.06);
}
.avg_box{
  height: 50px;
  line-height: 50px;
}
.reset_bt{
  width: calc(100vw - 40px);
  height: 50px;
  line-height: 50px;
  background: #4589F7;
  border-radius: 5px;
  color: #ffffff;
  font-size: 20px;
}
.box_m{
  width: 100%;
  display: flex;
  justify-content: center;
}

</style>
