<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
import Home from './components/Home'
import {initWebSocket} from './js/socket.js'
import {emptyString} from './js/tool.js'
export default {
  name: 'App',
  mounted () {
    if (emptyString(window.room_id) && emptyString(this.$route.query.code)) {
      console.log('跳转URL')
      redirectHome(this)
    } else if (!emptyString(this.$route.query.code)) {
      window.room_id = this.$route.query.code
      initWebSocket(this)
    }
  },
  components: {
    Home
  }
}

export function redirectHome(_this) {
  console.log(_this)
  _this.$axios
    .get(_this.httpUrl+'/api/url')
    .then(function (response) {
      console.log(response)
      let code = response.data.data.code
      window.room_id = code
      initWebSocket(_this)
      _this.$router.push({path: '/', query: {code: code}})
    })
    .catch(function (error) { // 请求失败处理
      console.log(error)
    })
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
}
html, body {
  margin: 0;
  padding: 0;
  height: 100%;
}
</style>
