<template>
    <div class="rank">
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <span>Rank</span>
                <el-button style="float: right; padding: 3px 0" type="text">MORE>></el-button>
            </div>
            <div v-for="(user, id) in rankedUsers" :key="user.id" class="rank-list text item">
                <span class="index">{{ id + 1 }}</span>
                <span class="grades">{{ user.grades }}</span>
                <span class="username">{{ user.username }}</span>
            </div>

        </el-card>
    </div>
</template>

<script>
    export default {
        name: 'HelloWorld',
        props: {
            msg: String
        },
        data() {
            return {
            }
        },
        computed: {
            rankedUsers() {
                return this.$store.getters.rankedUsers
            }
        },
        methods: {
            async getRankedUsers() {
                await this.$store.dispatch("getRankedUsers")
            }
        },
        created() {
            this.getRankedUsers()
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
    .rank {
        width: 100%;
    }
    .rank-list {
        width: 100%;
        box-shadow: 0px 1px 1px 0px #eee;
        .index {
            margin-right: 15px;
            font-style: italic;
        }
        .grades {
            margin-right: 15px;
        }
        .username {

        }
    }
</style>
