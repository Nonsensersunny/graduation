<template>
    <div>
        <div class="bread-nav">
            <el-breadcrumb separator-class="el-icon-arrow-right">
                <el-breadcrumb-item :to="{ path: '/' }">{{ $t('message.common.home') }}</el-breadcrumb-item>
                <el-breadcrumb-item>{{ $t('message.common.profile') }}</el-breadcrumb-item>
                <el-breadcrumb-item>{{ profile.username }}</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <el-card class="box-card">
            <div>{{ $t('message.profile.U') }}:{{ profile.username }}</div>
            <i class="el-icon-male"></i><i class="el-icon-female"></i>
            <div>{{ $t('message.profile.G') }}:{{ profile.grades }}</div>
            <div>{{ $t('message.profile.M') }}:{{ profile.mail }}</div>
            <div>{{ $t('message.profile.D') }}:{{ profile.description }}</div>
        </el-card>
        <el-card class="box-card" v-if="canOperate">
            <el-input v-model="user.username" :placeholder="profile.username"></el-input>
            <el-input v-model="user.description" :placeholder="profile.description"></el-input>
            <el-button @click="updateUserProfile" type="primary">Submit</el-button>
        </el-card>
        <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
            <el-form-item label="密码" prop="pass">
                <el-input type="password" v-model="ruleForm.pass" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="确认密码" prop="checkPass">
                <el-input type="password" v-model="ruleForm.checkPass" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="年龄" prop="age">
                <el-input v-model.number="ruleForm.age"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
                <el-button @click="resetForm('ruleForm')">重置</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>
<script>
    import {User, RespError} from "@/assets/js/type";
    export default {
        name: "Gprofile",
        data() {
            var checkAge = (rule, value, callback) => {
                if (!value) {
                    return callback(new Error('年龄不能为空'));
                }
                setTimeout(() => {
                    if (!Number.isInteger(value)) {
                        callback(new Error('请输入数字值'));
                    } else {
                        if (value < 18) {
                            callback(new Error('必须年满18岁'));
                        } else {
                            callback();
                        }
                    }
                }, 1000);
            };
            var validatePass = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请输入密码'));
                } else {
                    if (this.ruleForm.checkPass !== '') {
                        this.$refs.ruleForm.validateField('checkPass');
                    }
                    callback();
                }
            };
            var validatePass2 = (rule, value, callback) => {
                if (value === '') {
                    callback(new Error('请再次输入密码'));
                } else if (value !== this.ruleForm.pass) {
                    callback(new Error('两次输入密码不一致!'));
                } else {
                    callback();
                }
            };
            return {
                profile: {},
                user: new User(),
                ruleForm: {
                    pass: '',
                    checkPass: '',
                    age: ''
                },
                rules: {
                    pass: [
                        { validator: validatePass, trigger: 'blur' }
                    ],
                    checkPass: [
                        { validator: validatePass2, trigger: 'blur' }
                    ],
                    age: [
                        { validator: checkAge, trigger: 'blur' }
                    ]
                }
            };
        },
        computed: {
            id() {
                return this.$route.params['id']
            },
            canOperate() {
                return this.id == this.$store.getters.profile.id && this.$store.state.isLogin
            }
        },
        methods: {
            async getProfile() {
                if (this.id) {
                    this.profile = await this.$store.dispatch("getProfileById", this.id)
                } else if (this.username) {
                    this.profile = await this.$store.dispatch("getProfileByName", this.username)
                } else {
                    this.$message.error("Oops, something went wrong...");
                    this.$router.push("/")
                }
            },
            async updateUserProfile() {
                try {
                    this.user.id = this.$store.getters.profile.id
                    console.log(this.profile, this.user.description)
                    await this.$store.dispatch("updateUserProfile", this.user)
                    this.$store.dispatch("checkLoginStatus")
                    if (this.$store.state.isLogin) {
                        this.$message.success("Update user profile success")
                    }
                } catch (e) {
                    if (e instanceof RespError) {
                        this.$message.error("Update user profile failed");
                        this.$router.push('/')
                    }
                }
            },
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        alert('submit!');
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },
            resetForm(formName) {
                this.$refs[formName].resetFields();
            }
        },
        created() {
            this.getProfile()
        }
    }
</script>
<style scoped lang="stylus">
    .el-menu-vertical-demo:not(.el-menu--collapse) {
        width: 100px;
        min-height: 400px;
    }
    .bread-nav {
        background-color: #eee;
        padding: 10px;
    }
</style>