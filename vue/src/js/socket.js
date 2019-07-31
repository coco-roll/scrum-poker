var websock
export function initWebSocket (_this) { // 初始化weosocket
  const wsuri = _this.wsUrl //  这个地址由后端童鞋提供
  websock = new WebSocket(wsuri)
  websock.onmessage = websocketonmessage
  websock.onopen = websocketonopen
  websock.onerror = websocketonerror
  websock.onclose = websocketclose
}
export function websocketsend (Data) { // 数据发送
  websock.send(Data)
}
export function websocketonopen () { // 连接建立之后执行send方法发送数据
  console.log(111)
  let message = 'type=1&code=' + room_id
  websocketsend(message)
}
function websocketonerror () { //  连接建立失败重连
  initWebSocket()
}
export function websocketonmessage (e) {
  console.log(e)
}
function websocketclose (e) { // 关闭
  console.log('断开连接', e)
}
