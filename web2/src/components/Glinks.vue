<template>
    <div class="links">
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <span>{{ $t('message.common.links') }}</span>
                <el-button style="float: right; padding: 3px 0" type="text"><i :class="editable? 'el-icon-circle-check' : 'el-icon-edit'" @click="editable = !editable; addLink = !addLink"></i></el-button>
            </div>
            <div v-for="link in links" :key="link.id">
                <el-link :href="link.href" target="_blank" style="margin-bottom: 10px">{{ link.name }}</el-link><i style="float: right;" @click="deleteLink(link.id)" v-if="editable" class="el-icon-delete"></i>
            </div>
            <el-popover
                    v-if="addLink"
                    placement="top"
                    width="240">
                <strong>{{ $t("message.links.A") }}</strong>
                <el-form label-position="left" label-width="60px" :model="link">
                    <el-form-item :label="$t('message.links.LN')">
                        <el-input v-model="link.name"></el-input>
                    </el-form-item>
                    <el-form-item :label="$t('message.links.LA')">
                        <el-input v-model="link.href"></el-input>
                    </el-form-item>
                    <div style="text-align: center">
                        <el-button @click="createLink" :disabled="!(link.href && link.href)">{{ $t('message.profile.confirm') }}</el-button>
                    </div>
                </el-form>
                <i slot="reference" class="el-icon-circle-plus-outline"></i>
            </el-popover>
        </el-card>
    </div>
</template>

<script>
    import {Link} from "@/assets/js/type";

    export default {
        name: 'HelloWorld',
        props: {

        },
        data() {
            return {
                links: [],
                editable: false,
                addLink: false,
                link: new Link()
            }
        },
        computed: {

        },
        methods: {
            async getUserLinks() {
                let id = this.$store.getters.profile.id
                this.links = await this.$store.dispatch("getLinksByUserId", id)
            },
            async deleteLink(id) {
                let resp = await this.$store.dispatch("deleteLink", id);
                if (resp == 'success') {
                    await this.getUserLinks()
                    this.link = new Link()
                    this.addLink = false
                    this.$notify.info({
                        message: this.$t("message.links.LD")
                    })
                } else {
                    this.$notify.error({
                        message: this.$t("message.links.LDF")
                    })
                }
            },
            async createLink() {
                this.link.uid = this.$store.getters.profile.id
                let resp = await this.$store.dispatch("createLink", this.link)
                if (resp == 'success') {
                    await this.getUserLinks()
                } else {
                    this.$notify.error({
                        message: this.$t("message.links.ALF")
                    })
                }
            }
        },
        created() {
            this.getUserLinks();
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
