<template>
    <div class="signin">
<!--        <img alt="Vue logo" src="../assets/logo.png">-->
<!--        <HelloWorld msg="Welcome to Your Vue.js App"/>-->
        <el-form :model="user">
            <el-form-item label="Username" required><el-input v-model="user.username"></el-input> </el-form-item>
            <el-form-item label="Password" required><el-input type="password" auto-complete="false" @keyup.enter.native="login" v-model="user.password"></el-input> </el-form-item>
        </el-form>
        <el-button type="primary" @click="login">LOGIN</el-button>
        <el-button type="info" @click="signup">SIGNUP</el-button>
    </div>
</template>

<script>
    // @ is an alias to /src
    import HelloWorld from '@/components/HelloWorld.vue'
    import {RespError} from '@/assets/js/type'

    export default {
        name: 'signin',
        components: {
            HelloWorld
        },
        data() {
            return {
                user: {
                    username: '',
                    password: ''
                }
            }
        },
        methods: {
            async login() {
                try {
                    await this.$store.dispatch("userLogin", this.user);
                    if (this.$store.state.isLogin) {
                        this.$message.success("LOGIN SUCCESS");
                        this.$router.push("/")
                    }
                } catch (e) {
                    if (e instanceof RespError) {
                        this.$message.error("LOGIN FAILED");
                        return;
                    } else {
                        throw e;
                    }
                }
            },
            async signup() {
                try {
                    await this.$store.dispatch("userRegister", this.user);
                    this.$message.success("SIGNUP SUCCESS");
                } catch (e) {
                    console.log(e instanceof RespError)
                    if (e instanceof RespError) {
                        this.$message.error("SIGNUP FAILED");
                        return;
                    } else {
                        throw e;
                    }
                    this.user.username = '';
                }
            }
        }
    }
</script>