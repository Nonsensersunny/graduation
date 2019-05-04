<template>
    <div class="top">
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <span>{{ $t('message.common.top') }}</span>
                <!--<el-button style="float: right; padding: 3px 0" type="text">{{ $t('message.common.more') }}>></el-button>-->
            </div>
            <div v-for="content in contents" :key="content.id" class="text item" @click="goDetail(content.category, content.id)">
                <el-tag type="primary">{{ $t('message.list.' + content.category) }}</el-tag>
                <span style="margin-left: 10px; overflow: hidden">{{ content.title }}</span>
            </div>
        </el-card>
    </div>
</template>

<script>
    export default {
        name: 'Gtop',
        props: {

        },
        data() {
            return {
                contents: [],
            }
        },
        methods: {
            async getTopContent() {
                this.contents = await this.$store.dispatch("getTopContent")
            },
            goDetail(cat, id) {
                this.$router.push(`/content/${cat}/${id}`)
            }
        },
        created() {
            this.getTopContent()
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
