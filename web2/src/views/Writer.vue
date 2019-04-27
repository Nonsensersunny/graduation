<template>
    <div class="writer">
        <div class="bread-nav">
            <el-breadcrumb separator-class="el-icon-arrow-right">
                <el-breadcrumb-item :to="{ path: '/' }">Home</el-breadcrumb-item>
                <el-breadcrumb-item>Writer</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <el-card class="box-card">
            <el-input class="title" aria-required="true" v-model="title" placeholder="Title" label="Title"></el-input>
            <div class="category">
                <el-select v-model="selectedCategory" clearable placeholder="Category" aria-label="Category">
                    <el-option
                            v-for="item in category"
                            :key="item"
                            :label="item"
                            :value="item">
                    </el-option>
                </el-select>
            </div>
            <mavon-editor v-model="content" />
            <el-button @click="preview" type="primary">Preview</el-button>
            <el-button @click="createContent" type="success">Submit</el-button>
        </el-card>
    </div>
</template>

<script>
    export default {
        name: 'Writer',
        components: {

        },
        methods: {
            preview() {
                console.log(this.content)
            },
            async createContent() {
                let title = this.title
                let author = this.$store.getters.profile.id
                let category = this.selectedCategory
                let content = this.content
                await this.$store.dispatch("createContent", {title, author, category, content})
            }
        },
        data() {
            return {
                title: '',
                selectedCategory: '',
                content: '',
            }
        },
        computed: {
            category() {
                return ['Share', 'Q&A', 'Topic', 'Vote', 'Recruit']
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
    /*.writer {*/
        /*margin: 0 20px;*/
    /*}*/
    /*.writer * {*/
        /*display: block;*/
    /*}*/
    .title {
        max-width: 300px;
    }
    .bread-nav {
        background-color: #eee;
        padding: 10px;
    }
</style>
