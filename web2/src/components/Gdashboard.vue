<template>
    <div class="dashboard">
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <span>{{ $t('message.common.dashboard') }}</span>
                <el-button style="float: right; padding: 3px 0" type="text" @click="goSomeWhere(`/profile/${profile.id}`)">{{ $t('message.common.more') }}>></el-button>
            </div>
            <div class="card-body">
                <el-image
                        style="width: 100px; height: 100px"
                        :src="profile.avatar"
                        fit="fit">
                    <div slot="error" class="image-slot">
                        <i class="el-icon-picture-outline"></i>
                    </div>
                </el-image>
                <p @click="goSomeWhere(`/profile/${profile.id}`)">{{ profile.username }}</p>
                <div class="operation">
                    <el-tooltip :content="$t('message.dashboard.W')" placement="bottom">
                        <el-button circle icon="el-icon-document-add" @click="goSomeWhere('/writer')"></el-button>
                    </el-tooltip>
                    <el-tooltip :content="$t('message.dashboard.D')" placement="bottom">
                        <el-button  circle icon="el-icon-user-solid" @click="goSomeWhere(`/profile/${profile.id}`)"></el-button>
                    </el-tooltip>
                </div>
            </div>
        </el-card>
    </div>
</template>

<script>
    export default {
        name: 'HelloWorld',
        props: {
            activePath: ""
        },
        methods: {
            goSomeWhere(path) {
                this.$router.push(path)
            }
        },
        computed: {
            profile() {
                return this.$store.getters.profile
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
    .text {
        font-size: 14px;
    }

    .item {
        margin-bottom: 18px;
    }

    .clearfix:before,
    .clearfix:after {
        display: table;
        content: "";
    }
    .clearfix:after {
        clear: both
    }

    .box-card {
        width: 200px;
    }
    .card-body {
        text-align: center;
    }
    .operation {
        /*display: flex;*/
        margin: 0 auto;
    }
</style>
