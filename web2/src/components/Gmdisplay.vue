<template>
    <div class="display">
        <div v-html="compile(content)" v-if="content"></div>
    </div>
</template>
<script>
    import marked from 'marked'
    import hljs from "highlight.js"
    import javascript from 'highlight.js/lib/languages/javascript'
    import 'highlight.js/styles/github-gist.css'

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
                    return marked(input, {
                        renderer: new marked.Renderer(),
                        highlight: function (code) {
                            return hljs.highlightAuto(code).value
                        },
                        pedantic: false,
                        gfm: true,
                        tables: true,
                        breaks: false,
                        sanitize: false,
                        smartLists: true,
                        smartypants: true,
                        xhtml: false
                    })
                } else {
                    return "Oops, error happens..."
                }
            }
        }
    }
</script>
<style scoped lang="stylus">

</style>