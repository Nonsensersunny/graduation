<template>
    <div class="rank">
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <span>{{ $t('message.common.rank') }}</span>
                <!--<el-button style="float: right; padding: 3px 0" type="text">{{ $t('message.common.more') }}>></el-button>-->
            </div>
            <div v-for="(user, id) in rankedUsers" :key="user.id" class="rank-list text item" @click="$router.push(`/profile/${user.id}`)">
                <el-link @click="$router.push('/profile/' + user.id)">
                    <span class="index">{{ id + 1 }}</span>
                    <span class="grades">{{ user.grades }}</span>
                    <span class="username">{{ user.username }}</span>
                </el-link>
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
                return [{id:0,grades:888,username:'lilei'},
                {id:1,grades:756,username:'panda'},
                // {id:2,grades:745,username:'zhangsan'},
                // {id:3,grades:701,username:'zhaojuwen'},
                {id:4,grades:555,username:'hahaha'}]
                // return this.$store.getters.rankedUsers
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
.el-card{
        background-color:rgba(0,0,0,0.7)
        border-radius:10px;
    }
    .clearfix span{
        color:#fff;
    }
    .box-card {
        width: 200px;
    }
    .rank {
        width: 100%;
       
    }
 
    .rank-list {
        width: 100%;
        .index {
            margin-right: 15px;
            font-style: italic;
             color :#fff;
        }
        .grades {
            margin-right: 15px;
            color :#ccc;
        }
        .username {
        color :#ccc;
        }
    }
</style>
