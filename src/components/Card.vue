<template>
  <div class="card" @click="goHome">
    {{ num == 9999 ? '?' : num }}
  </div>
</template>

<script>
export default {
  name: 'Card',
  created () {
    this.num = this.$route.query.num
  },
  methods: {
    goHome () {
      this.$router.push({path: '/', query: {code: this.$route.query.code}})
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .card {
    background-color: #4589F7;
    width:100%;
    height:100%;
    font-size: 4cm;
    text-align: center;
    line-height:100%;
    display: flex;
    justify-content: center;
    align-items: center;
    color: #ffffff;
  }
</style>
