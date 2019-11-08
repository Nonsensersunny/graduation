<template>
    <div class="footer">
        <div>
            &copy;{{ new Date().getFullYear() }} <a class="resume" href="http://resume.zyven.top">Zyven</a> {{ $t('message.footer.A') }}
        </div>
        <el-link :underline="false" href="http://www.beian.gov.cn">{{ copyright }}</el-link>
    </div>
</template>

<script>
    export default {
        name: 'HelloWorld',
        data() {
            return {
                copyright: "鄂ICP备17023785号-1"
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
.footer {
    text-align: center;
    position: sticky;
    color :#fff;
    margin-bottom: 0;
}
    .resume {
        color: skyblue;
        text-decoration: none;
    }
    .el-link.el-link--default{
        color :#EEEED1;
    }
.resume:hover {
    color: deepskyblue;
}
</style>
