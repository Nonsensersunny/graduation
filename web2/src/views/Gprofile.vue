<template>
    <div>
        <el-menu default-active="1-4-1" class="el-menu-vertical-demo" :collapse="true">
            <el-submenu index="1">
                <template slot="title">
                    <i class="el-icon-location"></i>
                    <span slot="title">导航一</span>
                </template>
                <el-menu-item-group>
                    <span slot="title">分组一</span>
                    <el-menu-item index="1-1">选项1</el-menu-item>
                    <el-menu-item index="1-2">选项2</el-menu-item>
                </el-menu-item-group>
                <el-menu-item-group title="分组2">
                    <el-menu-item index="1-3">选项3</el-menu-item>
                </el-menu-item-group>
                <el-submenu index="1-4">
                    <span slot="title">选项4</span>
                    <el-menu-item index="1-4-1">选项1</el-menu-item>
                </el-submenu>
            </el-submenu>
            <el-menu-item index="2">
                <i class="el-icon-menu"></i>
                <span slot="title">导航二</span>
            </el-menu-item>
            <el-menu-item index="3" disabled>
                <i class="el-icon-document"></i>
                <span slot="title">导航三</span>
            </el-menu-item>
            <el-menu-item index="4">
                <i class="el-icon-setting"></i>
                <span slot="title">导航四</span>
            </el-menu-item>
        </el-menu>
        <el-card class="box-card">
            <div>{{ profile.username }}</div>
            <div>{{ profile.description }}</div>
        </el-card>
        <el-card class="box-card" v-if="canOperate">
            <el-input v-model="user.username" :placeholder="profile.username"></el-input>
            <el-input v-model="user.description" :placeholder="profile.description"></el-input>
            <el-button @click="updateUserProfile" type="primary">Submit</el-button>
        </el-card>
    </div>
</template>
<script>
    import {User, RespError} from "@/assets/js/type";
    export default {
        name: "Gprofile",
        data() {
            return {
                profile: {},
                user: new User(),
            }
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
                    console.log(this.profile)
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
</style>