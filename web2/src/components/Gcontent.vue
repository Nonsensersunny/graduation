<template>
    <div class="content">
        <el-tabs type="border-card">
            <el-tab-pane :label="$t('message.list.All')">
                <Glist category="All" />
            </el-tab-pane>
            <el-tab-pane v-for="cat in cats" :key="cat.id" :label="$store.state.lang == 'en'? cat.name : cat.alias">
                <Glist :category="cat.name" />
            </el-tab-pane>
            <el-tab-pane label="+" v-if="$store.getters.profile.role == 'admin'">
                <el-form label-position="left" label-width="80px" :model="category">
                    <el-form-item :label="$t('message.content.N')">
                        <el-input v-model="category.name"></el-input>
                    </el-form-item>
                    <el-form-item :label="$t('message.content.A')">
                        <el-input v-model="category.alias"></el-input>
                    </el-form-item>
                    <el-form-item :label="$t('message.content.W')">
                        <el-input v-model.number="category.weight" type="number"></el-input>
                    </el-form-item>
                </el-form>
                <el-button :disabled="!(category.weight && category.alias && category.name)" @click="createCategory">{{ $t("message.profile.confirm") }}</el-button>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>

<script>
    import Gdetail from '@/components/Gdetail.vue'
    import Glist from '@/components/Glist.vue'
    import {Category} from "@/assets/js/type";

    export default {
        name: 'HelloWorld',
        components: {
            Gdetail, Glist
        },
        data() {
            return {
                category: new Category()
            }
        },
        computed: {
            cats() {
                return this.$store.getters.categories
            }
        },
        methods: {
            async createCategory() {
                let category = this.category
                this.category.creator = this.$store.getters.profile.id;
                console.log(this.category)
                let resp = await this.$store.dispatch("createCategory", category);
                if (resp != 'success') {
                    this.$notify.error({
                        message: this.$t("message.content.CCF")
                    })
                } else {
                    this.category = new Category()
                    this.$notify.info({
                        message: this.$t("message.content.CCS")
                    })
                }
            }
        },
        created() {

        },
        props: {
            msg: String
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">

</style>