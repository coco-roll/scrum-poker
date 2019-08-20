<template>
  <div class="homeback">
    <!-- 规则弹框 -->
    <div class="rule-link-wrap" style="display: inline-block;">
        <a href="javascript:void(0);" class="d-rule-link rule-link" @click="enter">翻牌说明</a>
        <div class="rule-popup-box" id="rule-popup-box-01" v-show="isShow">
            <div class="tit">
                <h3>翻牌说明</h3>
            </div>
            <div class="con">
                <ul>
                    <li>1、本工具的页面头部url会带上小组的id，如果需要进不同的小组，只需要更改url上的id即可</li>
                    <li>2、本工具可供多人投票使用，当小组内所有人都投票完后会自动出现投票结果</li>
                    <li>3、在投票结果页面点击重新开始即可进行新的一轮投票</li>
                </ul>

            </div>
        </div>
    </div>
    <!-- 规则弹框结束 -->

    <div>
      <li v-for="num in num_config" :key='num' class="card" :data-id='num' v-on:click="cardClick(num)">
      {{ num == 9999 ? '?' : num }}
      </li>
    </div>
  </div>
</template>

<script>
import {cardClick, httpTest, enter, leave} from '../js/Home.js'
export default {
  name: 'Home',
  data () {
    return {
      num_config: [1, 2, 3, 4, 6, 9999],
      isShow: false,
    }
  },
  methods: {
    cardClick,
    httpTest,
    leave,
    enter
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
    background-color: #4589F7;
    border-radius:10px;
    border:1px solid #000;
    width: 30%;
    height:40%;
    font-size: 1.5cm;
    text-align: center;
    line-height:150px;
    margin: 20px;
    color: #ffffff;
  }

  .homeback{
    height: 100%;
    /*display: flex;*/
    /*justify-content: center;*/
    /*align-items: center;*/
  }

  .rule-popup-box{ position: absolute; top: 10%; right: 20%; z-index: 100; width: 50%; padding: 20px; background-color: #fff; border: 3px solid #c4c4c4; }
  .rule-popup-box .tit { border-bottom: 1px solid #dedede; }
  .rule-popup-box .tit h3 { padding-bottom: 10px; color: #333; font-size: 16px; }
  .rule-popup-box .con { padding: 10px 0 0 0; }
  .rule-popup-box .con ul { padding-bottom: 5px; }
  .rule-popup-box .con ul li { color: #333; padding: 2px 0 2px 20px; text-indent: -20px;list-style:none;}

</style>
