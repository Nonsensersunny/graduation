<template>
    <div class="links">
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <span>{{ $t('message.common.links') }}</span>
                </div>
            <div v-for="link in links" :key="link.id">
                <el-link :href="link.href" target="_blank">{{ link.desc }}</el-link>
            </div>
        </el-card>
    </div>
</template>

<script>
    export default {
        name: 'HelloWorld',
        props: {

        },
        data() {
            return {
                links: [
                    {
                        href: "nonsenser.top",
                        desc: "disk"
                    }
                ]
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
    .text {
        font-size: 14px;
    }

    .item {
        margin-bottom: 18px;
    }

    .clearfix:before,
    .clearfix:after {
        display: table;
        content: "";
    }
    .clearfix:after {
        clear: both
    }

    .box-card {
        width: 200px;
    }
</style>
