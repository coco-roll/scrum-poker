import {websocketsend} from './socket.js'
export function restart () {
  let code = getQueryString('code')
  let message = 'type=3&code=' + code
  console.log(message)
  websocketsend(message)
  this.$router.push({path: '/', query: {code: code}})
}
function getQueryString(name) {
  var reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)', 'i');

  var r = window.location.hash.substr(1).match(reg);
  if (r != null) {
    return unescape(r[2]);
  }
  return null;
}
