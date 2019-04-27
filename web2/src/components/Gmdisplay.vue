<template>
    <div class="display">
        <div v-html="compile(content)"></div>
    </div>
</template>
<script>
    import marked from 'marked';
    export default {
        name: "Gmdisplay",
        props: {
            content: {
                type: String,
                required: true,
            }
        },
        data() {
            return {
            }
        },
        methods: {
            compile(input) {
                return marked(input)
            }
        }
    }
</script>
<style scoped lang="stylus">

</style>