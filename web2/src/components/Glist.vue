<template>
    <div class="list">
        <div v-for="content in contents" :key="content.id">
            <el-divider content-position="left">{{ new Date(content.time).toLocaleString($store.state.lang) }}</el-divider>
            <el-row class="content-list-item" type="flex" :gutter="20">
                <el-col :span="1" v-if="$store.getters.profile.role == 'admin'">
                <!-- <el-col :span="1" > -->
                    <el-tooltip :content="content.is_key? $t('message.list.CH') : $t('message.list.PH')" placement="bottom">
                        <i :class="content.is_key? 'el-icon-lock' : 'el-icon-unlock'" @click="toggleKeyContent(content.id)"></i>
                    </el-tooltip>
                </el-col>
                <el-col :span="3">
                    <!-- <el-tooltip class="item" effect="dark" :content="$t('message.common.category') + ' : ' + $t('message.list.' + content.category)" placement="bottom"> -->
                    <el-tooltip class="item" effect="dark" :content="( ' 类型: ' + content.category)" placement="bottom">
                        <el-badge :value="content.is_key? $t('message.list.H') : ''" class="item">
                            <!-- <el-tag style="position: relative; bottom: 8px;" type="info">{{ $t('message.list.' + content.category) }}</el-tag> -->
                            <el-tag style="position: relative; bottom: 8px;" type="info">{{  content.category }}</el-tag>
                        </el-badge>
                    </el-tooltip>
                </el-col>
                <el-col :span="3">
                    <el-tooltip class="item" effect="dark" :content="$t('message.common.views') + ' : ' + content.views" placement="bottom">
                        <i class="el-icon-view"><span>      {{ content.views }}</span></i>
                    </el-tooltip>
                </el-col>
                <el-col :span="8">
                    <el-tooltip class="item" effect="dark" :content="$t('message.common.title') + ' : ' + content.title" placement="bottom">
                        <span  @click="goDetail(content.category, content.id)"><i class="el-icon-document"></i> {{ content.title }}</span>
                    </el-tooltip>
                </el-col>
                <el-col :span="6">
                    <el-tooltip class="item" effect="dark" :content="$t('message.common.author') + ' : ' + content.author" placement="bottom">
                        <span><i class="el-icon-user"></i> {{ content.author }}</span>
                    </el-tooltip>
                </el-col>
                <el-col :span="3">
                    <el-tooltip class="item" effect="dark" :content="$t('message.common.comment') + ' : ' + content.author" placement="bottom">
                        <span><i class="el-icon-chat-dot-square"></i> {{ content.author }}</span>
                    </el-tooltip>
                </el-col>
            </el-row>
            <el-divider></el-divider>
        </div>
        
    </div>
</template>

<script>
    export default {
        name: 'Glist',
        filters: {
            date(val) {
                return new Date(val).toLocaleString(_this.$store.state.lang)
            }
        },
        data() {
            return {
                contents: [{index:0,id:1,author:'lilei',comment:'hahaha',category:'计算机',views:152,time:new Date('2019'),title:'主流编程语言'},
                {index:1,id:2,author:'zhangdan',comment:'hahaha',category:'计算机',views:152,time:new Date('2019'),title:'论Python 的优点'}],
            }
        },
        computed: {
          c() {

          }
        },
        methods: {
          async getContentByCat(cat) {
              if (cat === 'All') {
                  this.contents = await this.$store.dispatch("getRankedContent");
              } else {
                  this.contents = await this.$store.dispatch("getContentByCat", cat)
              }
          },
            goDetail(cat, id) {
              this.$router.push(`/content/${cat}/${id}`)
            },
            async toggleKeyContent(id) {
                let resp = await this.$store.dispatch("toggleKeyContent", id);
                if (resp == 'success') {
                    await this.getContentByCat(this.category)
                    this.$notify.info({
                        message: this.$t("message.list.OS")
                    })
                } else {
                    this.$notify.error({
                        message: this.$t("message.list.OF")
                    })
                }
            }
        },
        created() {
            // this.getContentByCat(this.category) 这是后来注释的
        },
        props: {
            msg: String,
            category: {
                type: String,
                require: true,
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
    /*.text {*/
        /*font-size: 14px;*/
    /*}*/

    /*.item {*/
        /*margin-bottom: 18px;*/
    /*}*/

    /*.clearfix:before,*/
    /*.clearfix:after {*/
        /*display: table;*/
        /*content: "";*/
    /*}*/
    /*.clearfix:after {*/
        /*clear: both*/
    /*}*/

    .box-card {
        width: 180px;
    }
    /*.el-row {*/
        /*margin-bottom: 20px;*/
        /*&:last-child {*/
             /*margin-bottom: 0;*/
         /*}*/
    /*}*/
    /*.content-list-item {*/
        /*padding: 5px;*/
        /*margin: 10px;*/
        /*border-radius: 5px;*/
        /*box-shadow: 0px 1px 1px 0px #eeeeee;*/
        /*!*border: 1px solid #82848a;*!*/
    /*}*/
</style>
