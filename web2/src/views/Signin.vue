<template>
    <div class="signin">
<!--        <img alt="Vue logo" src="../assets/logo.png">-->
<!--        <HelloWorld msg="Welcome to Your Vue.js App"/>-->
        <el-form :model="user">
            <el-form-item :label="$t('message.signin.U')" required><el-input v-model="user.username"></el-input> </el-form-item>
            <el-form-item :label="$t('message.signin.P')" required><el-input type="password" auto-complete="false" @keyup.enter.native="login" v-model="user.password"></el-input> </el-form-item>
        </el-form>
        <div class="operation">
            <el-button type="primary" @click="login">{{ $t('message.signin.S') }}</el-button>
            <el-button type="info" @click="signup">{{ $t('message.common.signup') }}</el-button>
        </div>
    </div>
</template>

<script>
    // @ is an alias to /src
    import HelloWorld from '@/components/HelloWorld.vue'
    import {RespError} from '@/assets/js/type'

    export default {
        name: 'Signin',
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
                        this.$message.success(this.$t('message.signin.login_s'));
                        this.$router.push("/")
                    }
                } catch (e) {
                    if (e instanceof RespError) {
                        this.$message.error(this.$t('message.signin.login_f'));
                        return;
                    } else {
                        throw e;
                    }
                }
            },
            async signup() {
                try {
                    await this.$store.dispatch("userRegister", this.user);
                    this.$message.success(this.$t('message.signin.signup_s'));
                } catch (e) {
                    console.log(e instanceof RespError)
                    if (e instanceof RespError) {
                        this.$message.error(this.$t('message.signin.signup_f'));
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
<style lang="stylus">
    .signin {
        margin: 0 auto;
        width: 300px;
        padding-top: 100px;
        /*background-image: linear-gradient(#eee, #fff);*/
    }
    .operation {
        text-align: center;
        padding-bottom: 200px;
    }
</style>