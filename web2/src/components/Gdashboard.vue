<template>
    <div class="dashboard">
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <span>{{ $t('message.common.dashboard') }}</span>
                <el-button style="float: right; padding: 3px 0" type="text">{{ $t('message.common.more') }}>></el-button>
            </div>
            <p @click="goSomeWhere(`/profile/${profile.id}`)">{{ profile.username }}</p>
            <el-button @click="goSomeWhere('/writer')">{{ $t('message.dashboard.W') }}</el-button>
            <el-button @click="goSomeWhere(`/profile/${profile.id}`)">{{ $t('message.dashboard.D') }}</el-button>
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
</style>
