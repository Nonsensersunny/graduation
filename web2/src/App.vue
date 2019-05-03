<template>
  <div id="app">
    <el-menu background-color="skyblue" text-color="#fff" active-text-color="gray" :default-active="activePath" router mode="horizontal">
      <el-menu-item index="/">{{ $t('message.common.home') }}</el-menu-item>
      <el-menu-item v-if="isLogin" index="/ebook">{{ $t('message.common.books') }}</el-menu-item>
      <el-menu-item v-if="isLogin" index="/vote">{{ $t('message.common.votes') }}</el-menu-item>
      <el-menu-item v-if="isLogin" index="/about">{{ $t('message.common.about') }}</el-menu-item>
      <el-menu-item v-if="!isLogin" index="/signin">{{ $t('message.common.signin') }}</el-menu-item>
      <el-menu-item>
        <el-dropdown @command="selectLang">
          <span>{{ selectedLang }}<i class="el-icon-arrow-down el-icon--right"></i></span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item v-for="(key, val, id) in langs" :command="val" :key="id">{{ key }}</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </el-menu-item>
      <el-menu-item v-if="isLogin" @click="logout">{{ $t('message.common.logout') }}</el-menu-item>
    </el-menu>
    <router-view class="app-body" />
  </div>
</template>
<script>
  export default {
    name: 'app',
    beforeCreate() {
      this.$store.dispatch("checkLoginStatus")
    },
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
      },
      selectedLang() {
        return this.langs[this.$store.getters.selectedLang]
      },
      langs() {
        return this.$store.getters.langs
      }
    },
    methods: {
      selectLang(lang) {
        this.$store.commit("selectLang", lang);
      },
      logout() {
        this.$store.dispatch("userLogout")
        this.$router.push("/signin")
      },
      async getCategories() {
        await this.$store.dispatch("getCategories")
      }
    },
    created() {
      this.getCategories()
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
  #app {
    margin: 0 10% 0 10%;
    background-image: linear-gradient(#eef, #fff);
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
  }
</style>
