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
                    <el-tooltip :content="$t('message.common.author')" placement="bottom" class="content-info">
                        <el-link @click="$router.push('/profile/' + content.author)"><i class="el-icon-user"></i> {{ writer.username }}</el-link>
                    </el-tooltip>
                    <el-tooltip :content="$t('message.common.views')" placement="bottom" class="content-info">
                        <span><i class="el-icon-view"></i> {{ content.views }}</span>
                    </el-tooltip>
                    <el-tooltip :content="$t('message.common.category')" placement="bottom" class="content-info">
                        <span><i class="el-icon-folder"></i> {{ content.category }}</span>
                    </el-tooltip>
                    <el-tooltip :content="$t('message.article.T')" placement="bottom">
                        <span><i class="el-icon-watch"></i>{{ new Date(content.time).toLocaleString($store.state.lang) }}</span>
                    </el-tooltip>
                </div>
                <el-tooltip v-if="content.author != $store.getters.profile.id" class="item" effect="dark" :content="$t('message.article.A')" placement="bottom">
                    <i class="el-icon-star-off"  style="float: right; padding: 0px 3px"></i>
                </el-tooltip>
            </div>
            <Gmdisplay :content="content.content" v-if="content" />
            <el-divider v-if="isLogin" content-position="left">
                <h3>{{ $t('message.common.comment') }}</h3>
            </el-divider>
            <div v-for="(comment, index) in comments" :key="comment.id" class="comment-panel">
                <el-card shadow="never">
                    <div>
                        <span style="margin-right: 20px;">#{{ index + 1 }}</span>
                        <span style="margin-right: 20px;"><el-link @click="$router.push('/profile/' + comment.from)">{{ comment.from }}</el-link></span>
                        <span style="font-size: 12px;">{{ new Date(comment.time).toLocaleString($store.state.lang) }}</span>
                    </div>
                    <Gmdisplay :content="comment.content"/>
                </el-card>
            </div>
            <MarkDown v-if="isLogin" :to="content.author" :cid="content.id" />
        </el-card>
    </div>
</template>

<script>
    import Gmdisplay from '@/components/Gmdisplay.vue'
    import MarkDown from '@/components/MarkDown.vue'
    import {RespError} from "@/assets/js/type";

    export default {
        name: 'Garticle',
        components: {
            Gmdisplay, MarkDown
        },
        data() {
            return {
                content: {},
                comments: [],
                writer: {}
            }
        },
        methods: {
            async fetchContent(id) {
                try {
                    let resp = await this.$store.dispatch("getContentById", id);
                    this.content = resp.content
                    this.comments = resp.comments
                    this.writer = resp.writer
                } catch (e) {
                    if (e instanceof RespError) {
                        this.$message.error(this.$t('message.common.UE'))
                    }
                }

            },
        },
        computed: {
            category() {
                return this.$route.params['cat']
            },
            cid() {
                return this.$route.params['id']
            },
            isLogin() {
                return this.$store.state.isLogin
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
    background-image: linear-gradient(skyblue, #fff);
}
    .comment-panel * {
        margin-bottom: 5px;
    }
    .content-info {
        margin-right: 20px;
    }
</style>
