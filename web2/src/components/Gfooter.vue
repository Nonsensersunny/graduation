<template>
    <div class="footer">
        &copy;{{ new Date().getFullYear() }} Zyven {{ $t('message.footer.A') }}
    </div>
</template>

<script>
    export default {
        name: 'HelloWorld',
        props: {
            msg: String
        },
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
.footer {
    text-align: center;
}
</style>
