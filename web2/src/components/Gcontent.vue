<template>
    <div class="content">
        <el-tabs type="border-card">
            <el-tab-pane :label="$t('message.list.All')">
                <Glist category="All" />
            </el-tab-pane>
            <el-tab-pane v-for="cat in cats" :key="cat.id" :label="$t('message.list.' + cat.name)">
                <Glist :category="cat.name" />
            </el-tab-pane>
            <!--<el-tab-pane :label="$t('message.list.Share')">-->
                <!--<Glist category="Share" />-->
            <!--</el-tab-pane>-->
            <!--<el-tab-pane :label="$t('message.list.Q&A')">-->
                <!--<Glist category="Q&A" />-->
            <!--</el-tab-pane>-->
            <!--<el-tab-pane :label="$t('message.list.Recruit')">-->
                <!--<Glist category="Recruit" />-->
            <!--</el-tab-pane>-->
            <!--<el-tab-pane :label="$t('message.list.Topic')">-->
                <!--<Glist category="Topic" />-->
            <!--</el-tab-pane>-->
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
        computed: {
            cats() {
                return this.$store.getters.categories
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