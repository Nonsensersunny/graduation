<template>
    <div>
        <el-tabs type="border-card">
            <el-tab-pane label="Write">
                <el-input
                        type="textarea"
                        :autosize="{ minRows: 3}"
                        placeholder="Add your comment here..."
                        v-model="content">
                </el-input>
            </el-tab-pane>
            <el-tab-pane label="Preview">
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