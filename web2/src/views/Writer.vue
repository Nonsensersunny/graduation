<template>
  <div class="writer">
    <div class="bread-nav">
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/' }">{{ $t('message.common.home') }}</el-breadcrumb-item>
        <el-breadcrumb-item>{{ $t('message.writer.W') }}</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <el-card class="box-card">
      <div class="category">
        <span>标题：</span>
        <el-input
          style="width:30%;margin-rigth:25px;"
          class="title"
          aria-required="true"
          v-model="title"
          :placeholder="$t('message.common.title')"
          label="Title"
        ></el-input>

        <span style="margin-left:15px;">类型：</span>
        <el-select
          style="width:15%;"
          v-model="selectedCategory"
          clearable
          :placeholder="$t('message.common.category')"
          aria-label="Category"
        >
          <!-- <el-option
                            v-for="item in category"
                            :key="item.id"
                            :label="$store.state.lang == 'en'? item.name : item.alias"
                            :value="item.name">
          </el-option>-->
          <el-option
            v-for="item in category"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          ></el-option>
        </el-select>
      </div>
      <mavon-editor v-model="content" />
      <!-- <el-link
        @click="createContent"
        type="primary"
        v-if="title != '' && selectedCategory != '' && content != ''"
      >{{ $t('message.writer.S') }}</el-link> -->
      <el-link
        @click="createContent"
        type="primary"
        
      >{{ $t('message.writer.S') }}</el-link>
    </el-card>
  </div>
</template>

<script>
import { RespError } from "@/assets/js/type";
export default {
  name: "Writer",
  components: {},
  methods: {
    preview() {
      console.log(this.content);
    },
    async createContent() {
      try {
        let title = this.title;
        let author = this.$store.getters.profile.id;
        let category = this.selectedCategory;
        let content = this.content;
        let resp = await this.$store.dispatch("createContent", {
          title,
          author,
          category,
          content
        });
        if (resp === "success") {
          this.$message.success(this.$t("message.writer.AS"));
          this.$router.push("/");
        } else {
          this.$message.info(this.$t("message.writer.AF"));
        }
      } catch (e) {
        if (e instanceof RespError) {
          this.$message.error(e.error);
        }
      }
    }
  },
  data() {
    return {
      title: "",
      selectedCategory: "",
      content: ""
    };
  },
  computed: {
    category() {
      var temp = [{ value: "IT", label: "IT" }];
      return temp;
      // return this.$store.getters.categories
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
/* .writer { */
/* margin: 0 20px; */
/* } */
/* .writer * { */
/* display: block; */
/* } */
.title {
  max-width: 300px;
}

.bread-nav {
  background-color: lightgrey;
  padding: 10px;
}

.box-card {
  background-color: lightgrey;
  height: 700px;
}
.v-note-wrapper{
    min-height:450px;
}
.box-card * {
  margin-bottom: 20px;
}

</style>
