<template>
    <div class="ebook">
        <el-button>Chat</el-button>
    </div>
</template>

<script>
    export default {
        name: 'Gebook',
        sockets: {
            connect: function () {
                console.log("socket connected")
            },
            customEmit: function (data) {
                console.log('this method was fired by the socket server. eg: io.emit("customEmit", data)')
            }
        },
        methods: {
            clickBtn(data) {
                this.$socket.emit("emit_method", data)
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">

</style>
