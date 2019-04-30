<template>
    <div class="article">
        <div class="bread-nav">
            <el-breadcrumb separator-class="el-icon-arrow-right">
                <el-breadcrumb-item :to="{ path: '/' }">{{ $t('message.common.home') }}</el-breadcrumb-item>
                <el-breadcrumb-item>{{ content.category }}</el-breadcrumb-item>
                <el-breadcrumb-item>{{ content.title }}</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <h2>{{ content.title }}</h2>
                <div class="profile">
                    <el-tag>{{ $t('message.common.author') }}:{{content.author}}</el-tag>
                    <el-tag>{{ $t('message.common.views') }}:{{content.views}}</el-tag>
                    <el-tag>{{ $t('message.common.category') }}:{{content.category}}</el-tag>
                    <el-tag>{{ $t('message.article.T') }}:{{content.time | date}}</el-tag>
                </div>
                <el-tooltip v-if="content.author != $store.getters.profile.id" class="item" effect="dark" :content="$t('message.article.A')" placement="bottom">
                    <i class="el-icon-star-off"  style="float: right; padding: 0px 3px"></i>
                </el-tooltip>
            </div>
            <Gmdisplay :content="content.content" v-if="content" />
            <MarkDown />
            <div slot="footer" class="clearfix">

            </div>
        </el-card>
    </div>
</template>

<script>
    import Gmdisplay from '@/components/Gmdisplay.vue'
    import MarkDown from '@/components/MarkDown.vue'
    export default {
        name: 'Garticle',
        components: {
            Gmdisplay, MarkDown
        },
        filters: {
          date(val) {
              return new Date(val).toLocaleString("zh-CN")
          }
        },
        data() {
            return {
                content: {}
            }
        },
        methods: {
            async fetchContent(id) {
                this.content = await this.$store.dispatch("getContentById", id);
            },
        },
        computed: {
            category() {
                return this.$route.params['cat']
            },
            cid() {
                return this.$route.params['id']
            }
        },
        created() {
            this.fetchContent(this.cid)
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
.bread-nav {
    background-color: #eee;
    padding: 10px;
}
</style>
