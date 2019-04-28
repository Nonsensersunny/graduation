<template>
    <div class="dashboard">
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <span>Dashboard</span>
                <el-button style="float: right; padding: 3px 0" type="text">MORE>></el-button>
            </div>
            <p @click="$router.push(`/profile/${profile.id}`)">{{ profile.username }}</p>
            <el-button @click="goWriter">Add {{ activePath }}</el-button>
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
            goWriter() {
                this.$router.push('/writer')
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
