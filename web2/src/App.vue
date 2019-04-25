<template>
  <div id="app">
    <el-menu :default-active="activePath" router mode="horizontal">
      <el-menu-item index="/">Home</el-menu-item>
      <el-menu-item v-if="!isLogin" index="/ebook">Books</el-menu-item>
      <el-menu-item v-if="!isLogin" index="/vote">Votes</el-menu-item>
      <el-menu-item v-if="!isLogin" index="/about">About</el-menu-item>
      <el-menu-item v-if="!isLogin" index="/signin">Signin</el-menu-item>
    </el-menu>
    <router-view/>
  </div>
</template>
<script>
  export default {
    name: 'app',
    data() {
      return {
        activePath: '/'
      }
    },
    watch: {
      $route(newRoute) {
        this.activePath = newRoute.path;
      }
    },
    computed: {
      isLogin() {
        return this.$store.state.isLogin
      }
    }
  }
</script>
<style lang="stylus">
/*#app*/
  /*font-family 'Avenir', Helvetica, Arial, sans-serif*/
  /*-webkit-font-smoothing antialiased*/
  /*-moz-osx-font-smoothing grayscale*/
  /*text-align center*/
  /*color #2c3e50*/

/*#nav*/
  /*padding 30px*/
  /*a*/
    /*font-weight bold*/
    /*color #2c3e50*/
    /*&.router-link-exact-active*/
      /*color #42b983*/
</style>
