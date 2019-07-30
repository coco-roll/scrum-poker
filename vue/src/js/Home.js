import {websocketsend} from './socket.js'
export function cardClick (num) {
  let message = 'type=2&code=' + room_id + '&poker=' + num
  console.log(message)
  websocketsend(message)
  this.$router.push({path: '/card', query: {num: num}})
}

export function httpTest () {
  this.$axios
    .get('http://192.168.83.137:8000/api/test')
    .then(function (response) {
      console.log(response)
    })
    .catch(function (error) { // 请求失败处理
      console.log(error)
    })
}
