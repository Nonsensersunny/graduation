<template>
  <div id="app">
    <div class="head"><span>在 线 学 习 系 统</span></div>
    <div id="body">
    <el-menu  background-color="skyblue"  text-color="#000" active-text-color="#fff" :default-active="activePath" router mode="horizontal">
      <el-menu-item index="/">{{ $t('message.common.home') }}</el-menu-item>
      <!--<el-menu-item v-if="isLogin" index="/vote">{{ $t('message.common.votes') }}</el-menu-item>-->
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
 .head{
   height:100px;
   width :100%;
   /*background-color :#009FCC;*/
   text-align :center;
   font-size :26px;
   color:#fff;
   span{
    display: block;
    line-height: 100px;
   }
 }
#body{
  height :800px;
   /*background-image: url(../public/img/img1.jpg);*/
   background-size  :100%;

}
  #app {
    margin: 0 ;
   background-image: url(../public/img/img1.jpg);
   
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
  }
  .el-menu{
    opacity: 0.7;
  }
</style>
