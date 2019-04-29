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
        </el-tabs>
    </div>
</template>
<script>
    import Gmdisplay from '@/components/Gmdisplay.vue'
    export default {
        name: "MarkDown",
        components: {
            Gmdisplay,
        },
        data() {
            return {
                content: ''
            }
        },
        methods: {

        }
    }
</script>
<style>

</style>