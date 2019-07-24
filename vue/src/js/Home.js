export function cardClick (num) {
  this.$router.push({path: '/card', query: {num: num}})
}
