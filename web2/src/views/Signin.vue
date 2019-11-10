<template>
    <div class="signin">
<!--        <img alt="Vue logo" src="../assets/logo.png">-->
<!--        <HelloWorld msg="Welcome to Your Vue.js App"/>-->
        <el-form :model="user">
            <el-form-item :label="$t('message.signin.U')" required><el-input v-model="user.username"></el-input> </el-form-item>
            <el-form-item :label="$t('message.signin.P')" required><el-input type="password" auto-complete="false" @keyup.enter.native="login" v-model="user.password"></el-input> </el-form-item>
            <!--<el-link type="primary" @click="isRegister = true">{{ $t('message.signin.R') }}</el-link>-->
            <!--<el-link type="info" style="float: right;">{{ $t('message.signin.F') }}</el-link>-->
        </el-form>
        <div class="operation">
            <el-button type="primary" :disabled="user.username == '' || user.password == ''" @click="login">{{ $t('message.signin.S') }}</el-button>
            <el-button type="info" @click="isRegister = true">{{ $t('message.common.signup') }}</el-button>
        </div>
        <el-dialog
                :title="$t('message.signin.R')"
                :visible.sync="isRegister"
                width="30%"
                center>
            <el-form :model="userForm" status-icon :rules="rules" ref="userForm" label-width="100px">
                <el-form-item :label="$t('message.signin.U')" prop="username">
                    <el-input v-model="userForm.username" :placeholder="$t('message.signin.MS')"></el-input>
                </el-form-item>
                <el-form-item :label="$t('message.signin.P')" prop="password">
                    <el-input type="password" v-model="userForm.password" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item :label="$t('message.signin.CP')" prop="checkPass">
                    <el-input type="password" v-model="userForm.checkPass" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item :label="$t('message.signin.AC')" prop="authCode">
                    <el-input style="width: 70%" v-if="authSent" v-model="userForm.authCode" autocomplete="off"></el-input>
                    <el-button  :disabled="userForm.username == ''" v-else @click="getAuthMail">{{ $t('message.signin.GAC') }}</el-button>
                    <el-tooltip v-if="authSent" :content="$t('message.signin.RS')" placement="bottom">
                        <el-button  :disabled="userForm.username == ''" @click="getAuthMail" style="float: right;"><i class="el-icon-refresh"></i></el-button>
                    </el-tooltip>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="isRegister = false">{{ $t('message.profile.cancel') }}</el-button>
                <el-button type="primary" @click="signup">{{ $t('message.profile.confirm') }}</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
    // @ is an alias to /src
    import HelloWorld from '@/components/HelloWorld.vue'
    import {RespError} from '@/assets/js/type'

    export default {
        name: 'Signin',
        data() {
            var validateUsername = (rule, value, callback) => {
                if (value == '') {
                    callback(new Error(this.$t("message.signin.EUF")));
                } else {
                    if (this.userForm.username != '') {
                        var reg = new RegExp(/^([a-zA-Z0-9._-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-])+/)
                        if (!reg.test(this.userForm.username)) {
                            return callback(new Error(this.$t("message.signin.IM")))
                        }
                        this.$refs.userForm.validateField('password');
                    }
                    callback();
                }
            };
            var validatePassword = (rule, value, callback) => {
                if (value == '') {
                    callback(new Error(this.$t("message.signin.EP")));
                } else {
                    if (this.userForm.password != '') {
                        var reg = new RegExp(/^(\w){6,20}$/);
                        if (!reg.test(this.userForm.password)) {
                            return callback(new Error(this.$t("message.signin.IP")))
                        }
                        this.$refs.userForm.validateField('checkPass');
                    }
                    callback();
                }
            };
            var checkPassword = (rule, value, callback) => {
                if (value == '') {
                    callback(new Error(this.$t("message.signin.EPA")));
                } else if (value != this.userForm.password) {
                    callback(new Error(this.$t("message.signin.PC")));
                } else {
                    callback();
                }
            };
            var checkAuthCode = (rule, value, callback) => {
                if (value == '') {
                    callback(new Error('请输入auth code'));
                } else {
                    if (this.userForm.checkPass !== '') {
                        // this.$refs.userForm.validateField('authCode');
                    }
                    callback();
                }
            };
            return {
                user: {
                    username: '',
                    password: ''
                },
                isRegister: false,
                authSent: false,
                rules: {
                    username: [
                        { validator: validateUsername, trigger: 'blur' }
                    ],
                    password: [
                        { validator: validatePassword, trigger: 'blur' }
                    ],
                    checkPass: [
                        { validator: checkPassword, trigger: 'blur' }
                    ],
                    authCode: [
                        { validator: checkAuthCode, trigger: 'blur' }
                    ]
                },
                userForm: {
                    username: '',
                    password: '',
                    checkPass: '',
                    authCode: ''
                }
            }
        },
        methods: {
            async login() {
                this.$message.success(this.$t('message.signin.login_s'));
                        this.$router.push("/")
                // try {
                //     await this.$store.dispatch("userLogin", this.user);
                //     if (this.$store.state.isLogin) {
                //         this.$message.success(this.$t('message.signin.login_s'));
                //         this.$router.push("/")
                //     }
                // } catch (e) {
                //     if (e instanceof RespError) {
                //         this.$message.error(this.$t('message.signin.login_f'));
                //         return;
                //     } else {
                //         throw e;
                //     }
                // }
            },
            async signup() {
                try {
                    let user = this.userForm
                    let resp = await this.$store.dispatch("userRegister", user);
                    if (resp == 'success') {
                        this.isRegister = false
                        this.user.username = this.userForm.username
                        this.$message.success(this.$t('message.signin.signup_s'));
                    } else {
                        this.$message.error(this.$t('message.signin.signup_f'));
                    }
                } catch (e) {
                    if (e instanceof RespError) {
                        this.$message.error(this.$t('message.signin.signup_f'));
                        return;
                    } else {
                        throw e;
                    }
                    this.user.username = '';
                }
            },
            async getAuthMail() {
                let user = this.userForm
                let resp = await this.$store.dispatch("getAuthMail", user);
                if (resp == 'success') {
                    this.authSent = true
                    this.$message.success(this.$t("message.signin.AMS"))
                } else {
                    this.$message.error(this.$t("message.signin.AMF"))
                }
            }
        }
    }
</script>
<style lang="stylus">
    .signin {
        
        margin: 100px auto;
        padding:30px;
        border-radius: 20px;
        height :250px;
        width: 300px;
        // padding-top: 100px;
        background: rgba(#000,0.8);
    }
    .el-form-item__label{
        color:#fff;
    }
    .operation {
        text-align: center;
        margin-bottom: 200px;
    }
</style>