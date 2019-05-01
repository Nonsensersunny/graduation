<template>
    <div class="profile">
        <div class="bread-nav">
            <el-breadcrumb separator-class="el-icon-arrow-right">
                <el-breadcrumb-item :to="{ path: '/' }">{{ $t('message.common.home') }}</el-breadcrumb-item>
                <el-breadcrumb-item>{{ $t('message.common.profile') }}</el-breadcrumb-item>
                <el-breadcrumb-item>{{ profile.username }}</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <el-row :gutter="20" class="box-card">
            <el-col :span="6" class="left-panel">
                <el-card class="left-panel-card">
                    <el-image
                            style="width: 100px; height: 100px"
                            :src="profile.avatar"
                            fit="fit" class="profile-img">
                        <div slot="error" class="image-slot">
                            <i class="el-icon-picture-outline"></i>
                        </div>
                    </el-image>
                    <div class="left-panel-card-info">
                        <div class="profile-div" style="text-align: center; font-size: 20px;margin: 10px;">
                            <span>{{ profile.username }}</span>
                            <span>
                                <i class="el-icon-male" v-if="profile.sex = 1"></i>
                                <i class="el-icon-female" v-else></i>
                                <el-tooltip :content="$t('message.profile.E')" placement="bottom" v-if="canOperate">
                                  <i class="el-icon-edit edit-icon" @click="editable = !editable"></i>
                                </el-tooltip>

                            </span>
                        </div>
                        <div v-if="profile.mail != ''">{{ $t('message.profile.M') }}:{{ profile.mail }}</div>
                        <div v-if="profile.description != ''">{{ $t('message.profile.D') }}:{{ profile.description }}</div>
                    </div>
                </el-card>
            </el-col>
            <el-col :span="18" class="right-panel">
                <el-row class="right-panel-items" :gutter="20">
                    <el-col :span="8">
                        <el-card shadow="hover">
                            <div>{{ $t('message.profile.G') }}:{{ profile.grades }}</div>
                        </el-card>
                    </el-col>
                    <el-col :span="8">
                        <el-card shadow="hover">
                            <div>{{ $t('message.profile.CN') }}:{{ profile.comme_num }}</div>
                        </el-card>
                    </el-col>
                    <el-col :span="8">
                        <el-card shadow="hover">
                            <div>{{ $t('message.profile.QN') }}:{{ profile.quest_num }}</div>
                        </el-card>
                    </el-col>
                    <el-col :span="8">
                        <el-card shadow="hover">
                            <div>{{ $t('message.profile.TN') }}:{{ profile.topic_num }}</div>
                        </el-card>
                    </el-col>
                    <el-col :span="8">
                        <el-card shadow="hover">
                            <div>{{ $t('message.profile.SN') }}:{{ profile.share_num }}</div>
                        </el-card>
                    </el-col>
                    <el-col :span="8">
                        <el-card shadow="hover">
                            <div>{{ $t('message.profile.RN') }}:{{ profile.recui_num }}</div>
                        </el-card>
                    </el-col>
                </el-row>
            </el-col>
        </el-row>
        <el-dialog :title="$t('message.profile.EP')" :visible.sync="editable">
            <el-form :model="user">
                <el-form-item :title="$t('message.profile.avatar')">
                    <el-upload
                            action="''"
                            :http-request="catchAvatar"
                            :show-file-list="false"
                            list-type="picture-card"
                            :on-success="handleAvatarSuccess"
                            :on-progress="handleAvatarProgress"
                            :on-preview="handlePictureCardPreview"
                            :before-upload="beforeAvatarUpload">
                        <img v-if="user.avatar" :src="user.avatar" width="150px" height="150px">
                        <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                    </el-upload>
                </el-form-item>
                <el-form-item>
                    <el-input :placeholder="$t('message.profile.U')" v-model="user.username" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item :label="$t('message.profile.S')">
                    <el-radio v-model="user.sex" :label="1">{{ $t('message.profile.boy') }}</el-radio>
                    <el-radio v-model="user.sex" :label="0">{{ $t('message.profile.girl') }}</el-radio>
                </el-form-item>
                <el-form-item>
                    <el-input placeholder="example@example.com" v-model="user.mail" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-input :placeholder="$t('message.profile.D')" v-model="user.description" autocomplete="off"></el-input>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="cancelChange">{{ $t('message.profile.cancel') }}</el-button>
                <el-button type="primary" @click="submitChange">{{ $t('message.profile.confirm') }}</el-button>
            </div>
        </el-dialog>
    </div>
</template>
<script>
    import {User, RespError} from "@/assets/js/type";
    import {uploadAvatar} from "@/assets/js/av";

    export default {
        name: "Gprofile",
        data() {
            return {
                profile: {},
                user: new User(),
                editable: false,
                ruleForm: {
                    pass: '',
                    checkPass: '',
                    age: ''
                },
                img_file: ''
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
            cancelChange() {
                this.editable = false
                window.URL.revokeObjectURL(this.user.avatar)
            },
            catchAvatar(file) {
                this.img_file = file.file
                let url = window.URL.createObjectURL(file.file)
                this.user.avatar = url
            },
            handleAvatarProgress(ev, rawFile) {
                console.log(URL.createObjectURL(rawFile.raw))
            },
            async submitChange() {
                let file = this.img_file
                let name = file.name
                uploadAvatar(file, name, async (res) => {
                    this.user.id = this.$store.getters.profile.id
                    this.user.avatar = this.user.avatar? this.user.avatar : this.$store.getters.avatar
                    this.user.username = this.user.username? this.user.username : this.$store.getters.username
                    this.user.description = this.user.description? this.user.description : this.$store.getters.description
                    this.user.mail = this.user.mail? this.user.mail : this.$store.getters.mail
                    this.user.sex = this.user.sex? this.user.sex : this.$store.getters.sex
                    try {
                        this.user.avatar = res.attributes.url
                        await this.$store.dispatch("updateUserProfile", this.user)
                        this.editable = false;
                        this.$store.dispatch("checkLoginStatus")
                        if (this.$store.state.isLogin) {
                            this.$message.success($t('message.profile.UPS'))
                            await this.getProfile()
                        }
                    } catch (e) {
                        if (e instanceof RespError) {
                            this.$message.error($t('message.profile.UPF'));
                        }
                    }
                })

            },
            handleAvatarSuccess(res, file) {
                this.user.avatar = URL.createObjectURL(file.raw);
                console.log(URL.createObjectURL(file.raw))
                console.log(this.user.avatar)
            },
            beforeAvatarUpload(file) {
                const isPNG = file.type === 'image/png';
                const isJPG = file.type === 'image/jpeg';
                const isLt2M = file.size / 1024 / 1024 < 2;

                if (!isJPG && !isPNG) {
                    this.$message.error(this.$t("message.profile.format_e"));
                }
                if (!isLt2M) {
                    this.$message.error(this.$t("message.profile.size_e"));
                }
                return isJPG && isLt2M;
            },
            handlePictureCardPreview(file) {
                this.user.avatar = file.url;
            },
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
    .bread-nav {
        background-color: #eee;
        padding: 10px;
    }
    .profile {
        width: 100%;
    }
    .box-card {
        padding: 20px;
    }
    .box-card .left-panel {
        padding: 20px;
        margin: 0 auto;
        border-right: 1px solid gray;
    }
    .left-panel-card * {
        margin: 0 auto;
    }
    .left-panel-card-info {
        margin: 10px;
        justify-content: center;
    }
    .box-card .right-panel {
        flex: 0.8;

    }
    .box-card .right-panel {
        height: 100%;
    }
    .right-panel-items {
        height: 100%;
    }
    .right-panel-items * {
        height: 100px;
        padding: 10px;
        text-align: center;
    }
    .profile-img {
        text-align: center;
        display: flex;
        justify-content: center;
    }
    .profile-div span {
        margin: 10px;
    }
    .edit-icon {
        opacity: 0.2;
    }
    .edit-icon:hover {
        color: skyblue;
        opacity: 0.8;
    }
</style>