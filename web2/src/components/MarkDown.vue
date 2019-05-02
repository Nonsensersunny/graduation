<template>
    <div>
        <el-tabs type="border-card">
            <el-tab-pane :label="$t('message.markdown.C')">
                <el-input
                        type="textarea"
                        :autosize="{ minRows: 3}"
                        :placeholder="$t('message.markdown.A')"
                        v-model="content">
                </el-input>
            </el-tab-pane>
            <el-tab-pane :label="$t('message.markdown.P')" :disabled="content == ''">
                <Gmdisplay :content="content" />
            </el-tab-pane>
            <el-button type="text" :disabled="content == ''" @click="createComment">{{ $t('message.markdown.S') }}</el-button>
        </el-tabs>
    </div>
</template>
<script>
    import Gmdisplay from '@/components/Gmdisplay.vue'
    import {Comment} from "@/assets/js/type";

    export default {
        name: "MarkDown",
        props: {
            to: Number,
            cid: Number
        },
        components: {
            Gmdisplay,
        },
        data() {
            return {
                content: '',
            }
        },
        methods: {
            async createComment() {
                let comment = new Comment(this.$store.getters.profile.id, this.to, this.cid, this.content)
                let resp = await this.$store.dispatch("createComment", comment)
                if (resp == "success") {
                    this.$message.success(this.$t("message.markdown.CS"))
                    this.content = ''
                } else {
                    this.$message.error(this.$t("message.markdown.CF"))
                }
            }
        },

    }
</script>
<style>

</style>