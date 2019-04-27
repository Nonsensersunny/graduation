<template>
    <div class="links">
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <span>Links</span>
                <el-button style="float: right; padding: 3px 0" type="text">MORE>></el-button>
            </div>
            <div v-for="link in links" :key="link.id" class="text item">
                {{ link }}
            </div>
        </el-card>
    </div>
</template>

<script>
    export default {
        name: 'HelloWorld',
        props: {
            links: []
        },
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
