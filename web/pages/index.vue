<template>
  <section class="container">
    <div>
      <logo />
      <h1 class="title">
        web
      </h1>
      <h2 class="subtitle">
        My neat Nuxt.js project
      </h2>
      <div class="links">
        <a
          href="https://nuxtjs.org/"
          target="_blank"
          class="button--green"
        >Documentation</a>
        <a
          href="https://github.com/nuxt/nuxt.js"
          target="_blank"
          class="button--grey"
        >GitHub</a>
      </div>
      <el-form :model="user">
        <el-form-item label="Username" required><el-input v-model="user.username"></el-input> </el-form-item>
        <el-form-item label="Password" required><el-input type="password" auto-complete="false" @keyup.enter.native="login" v-model="user.password"></el-input> </el-form-item>
      </el-form>
      <el-button type="primary" @click="login">LOGIN</el-button>
      <el-button type="info" @click="signup">SIGNUP</el-button>
    </div>
  </section>
</template>

<script>
import Logo from '~/components/Logo.vue'
import {RespError} from "~/assets/type.js"

export default {
  components: {
    Logo
  },
  data() {
    return {
      user: {
        username: "",
        password: "",
      }
    }
  },
  methods: {
    async login() {
      console.log("LOGIN")
      try {
        await this.$store.dispatch("userLogin", this.user);
      } catch (e) {
        if (e instanceof RespError) {
          this.$message.error("LOGIN FAILED");
          return;
        } else {
          throw e;
        }
        this.user.username = '';
        this.$message.success("LOGIN SUCCESS");
      }
    },
    signup() {

    }
  }
}
</script>

<style>
.container {
  margin: 0 auto;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.title {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
