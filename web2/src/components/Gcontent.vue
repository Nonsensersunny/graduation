<template>
    <div class="content">
        <el-button @click="getDetail(0)">Get Detail</el-button>
        <!--<Gdetail :content="this.$store.context.content" v-if="getDetail(0)" />-->
    </div>
</template>

<script>
    import Gdetail from '@/components/Gdetail.vue'
    export default {
        name: 'HelloWorld',
        components: {
            Gdetail,
        },
        data() {
            return {
            }
        },
        methods: {
            async getDetail(id) {
                let resp = await this.$store.dispatch("getContentById", id);
                console.log(resp)
            }
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