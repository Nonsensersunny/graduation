<template>
    <div class="display">
        <div v-html="compile(content)" v-if="content"></div>
    </div>
</template>
<script>
    import marked from 'marked';
    export default {
        name: "Gmdisplay",
        props: {
            content: {
                type: String,
                required: false,
            }
        },
        data() {
            return {
            }
        },
        methods: {
            compile(input) {
                if (input) {
                    return marked(input)
                } else {
                    return "Oops, error happens..."
                }
            }
        }
    }
</script>
<style scoped lang="stylus">

</style>