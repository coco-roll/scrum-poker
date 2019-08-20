import {websocketsend} from './socket.js'
export function cardClick (num) {
  let message = 'type=2&code=' + window.room_id + '&poker=' + num
  console.log(message)
  websocketsend(message)
  this.$router.push({path: '/card', query: {num: num, code: window.room_id}})
}

export function httpTest () {
  this.$axios
    .get(this.httpUrl+'/api/test')
    .then(function (response) {
      console.log(response)
    })
    .catch(function (error) { // 请求失败处理
      console.log(error)
    })
}

export function enter () {
  if(this.isShow){
    this.isShow = false;
  }else{
    this.isShow = true;
  }
  
}

