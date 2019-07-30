<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
import Home from './components/Home'
import {initWebSocket} from './js/socket.js'
export default {
  name: 'App',
  mounted () {
    window.room_id = this.$route.query.code
    initWebSocket()
  },
  components: {
    Home
  }
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
