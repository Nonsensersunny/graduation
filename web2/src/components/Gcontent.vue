<template>
    <div class="content">
        <el-tabs type="border-card">
            <el-tab-pane :label="$t('message.list.A')">
                <Glist category="All" />
            </el-tab-pane>
            <el-tab-pane :label="$t('message.list.S')">
                <Glist category="Share" />
            </el-tab-pane>
            <el-tab-pane :label="$t('message.list.Q')">
                <Glist category="Q&A" />
            </el-tab-pane>
            <el-tab-pane :label="$t('message.list.R')">
                <Glist category="Recruit" />
            </el-tab-pane>
            <el-tab-pane :label="$t('message.list.T')">
                <Glist category="Topic" />
            </el-tab-pane>
        </el-tabs>
    </div>
</template>

<script>
    import Gdetail from '@/components/Gdetail.vue'
    import Glist from '@/components/Glist.vue'
    export default {
        name: 'HelloWorld',
        components: {
            Gdetail, Glist
        },
        data() {
            return {
            }
        },
        methods: {

        },
        created() {

        },
        props: {
            msg: String
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">

</style>