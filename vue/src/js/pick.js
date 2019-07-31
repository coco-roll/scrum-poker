import {websocketsend} from './socket.js'
export function restart () {
  let message = 'type=3&code=' + 4037200794235010051
  console.log(message)
  websocketsend(message)
  this.$router.push({path: '/', query: {}})
}
