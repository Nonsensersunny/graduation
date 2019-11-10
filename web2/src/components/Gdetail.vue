<template>
    <div class="detail">
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <h2 class="title">{{ content.title }}</h2>
                <el-button style="float: right; padding: 3px 0" type="text"><i class="el-icon-star-off"></i> Favorite</el-button>
                <div class="info">
                    <el-tag type="success">{{ content.time }}</el-tag>
                    <el-tag type="primary">{{ content.author }}</el-tag>
                    <el-tag type="info"><i class="el-icon-view"></i>  {{ content.views }}</el-tag>
                    <el-tag type="warning">{{ content.category }}</el-tag>
                </div>
            </div>
            <div class="content-wrapper">
                <p>{{ content.content }}</p>
            </div>
        </el-card>
    </div>
</template>

<script>
    export default {
        name: 'Gdetail',
        props: {
            msg: String,
            content: {
                type: Object,
                required: true
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
  .el-card{
    background-color:#000;
  }
</style>
