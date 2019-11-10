<template>
  <div class="article">
    <div class="bread-nav">
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/' }">{{ $t('message.common.home') }}</el-breadcrumb-item>
        <el-breadcrumb-item>{{ content.category }}</el-breadcrumb-item>
        <el-breadcrumb-item>{{ content.title }}</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <el-card class="box-card">
      <div slot="header" class="clearfix">
        <h2>{{ content.title }}</h2>
        <div class="profile">
          <el-tooltip
            :content="$t('message.common.author')"
            placement="bottom"
            class="content-info"
          >
            <el-link @click="$router.push('/profile/' + content.author)">
              <i class="el-icon-user"></i>
              {{ writer.username }}
            </el-link>
          </el-tooltip>
          <el-tooltip :content="$t('message.common.views')" placement="bottom" class="content-info">
            <span>
              <i class="el-icon-view"></i>
              {{ content.views }}
            </span>
          </el-tooltip>
          <el-tooltip
            :content="$t('message.common.category')"
            placement="bottom"
            class="content-info"
          >
            <span>
              <i class="el-icon-folder"></i>
              {{ content.category }}
            </span>
          </el-tooltip>
          <el-tooltip :content="$t('message.article.T')" placement="bottom">
            <span>
              <i class="el-icon-watch"></i>
              {{ new Date(content.time).toLocaleString($store.state.lang) }}
            </span>
          </el-tooltip>
        </div>
        <!-- <el-tooltip v-if="content.author != $store.getters.profile.id && isLogin" class="item" effect="dark" :content="favorite? $t('message.article.C') : $t('message.article.A')" placement="bottom">
                    <i :class="favorite? 'el-icon-star-on' : 'el-icon-star-off'"  style="float: right; padding: 0px 3px" @click="toggleFavorite"></i>
        </el-tooltip>-->
        <el-tooltip
          class="item"
          effect="dark"
          :content="favorite? $t('message.article.C') : $t('message.article.A')"
          placement="bottom"
        >
          <i
            :class="favorite? 'el-icon-star-on' : 'el-icon-star-off'"
            style="float: right; padding: 0px 3px"
            @click="toggleFavorite"
          ></i>
        </el-tooltip>
      </div>
      <Gmdisplay :content="content.content" v-if="content" />
      <el-divider v-if="isLogin" content-position="left">
        <h3>{{ $t('message.common.comment') }}</h3>
      </el-divider>
      <el-card style=" background-color:lightgrey;" shadow="never">
        <div
          style="margin:5px 1%"
          v-for="(comment, index) in comments"
          :key="comment.id"
          class="comment-panel"
        >
          <div>
            <span style="margin-right: 20px;">{{ index + 1 }}楼</span>
            <span style="margin-right: 20px;">
              <el-link @click="$router.push('/profile/' + comment.from)">{{ comment.from }}</el-link>
            </span>
            <span
              style="font-size: 12px;"
            >{{ new Date(comment.time).toLocaleString($store.state.lang) }}</span>
          </div>
          <Gmdisplay style="margin-left:40px;" :content="comment.content" />
        </div>
      </el-card>
      <el-divider v-if="isLogin" content-position="left">
        <h3>我的</h3>
      </el-divider>
      <MarkDown style="margin-top:30px ;" v-if="isLogin" :to="content.author" :cid="content.id" />
    </el-card>
  </div>
</template>

<script>
import Gmdisplay from "@/components/Gmdisplay.vue";
import MarkDown from "@/components/MarkDown.vue";
import { RespError } from "@/assets/js/type";

export default {
  name: "Garticle",
  components: {
    Gmdisplay,
    MarkDown
  },
  data() {
    return {
      content: {
        category: "计算机",
        title: "主流编程语言",
        views: 153,
        time: new Date("2019/10/29"),
        content:
          "1111111111111111111111111111111111111111111111111111111111111111111111111111111\n111111111111"
      },
      comments: [
        {
          id: 1,
          index: 0,
          from: "zhangsan",
          time: new Date("2019/10/31"),
          content: "哈哈哈"
        },
        {
          id: 2,
          index: 1,
          from: "lisi",
          time: new Date("2019/11/1"),
          content: "同意"
        },
        {
          id: 3,
          index: 2,
          from: "zhangsan",
          time: new Date("2019/10/31"),
          content: "+1"
        },
        {
          id: 4,
          index: 3,
          from: "zhangsan",
          time: new Date("2019/10/31"),
          content: "xswl"
        },
      ],
      writer: { username: "lilei" },
      favorite: true,
      isLogin: true //新加的
    };
  },
  methods: {
    async fetchContent(id) {
      try {
        let resp;
        if (this.isLogin) {
          let uid = this.$store.getters.profile.id;
          let cid = id;
          resp = await this.$store.dispatch("userGetContentById", { cid, uid });
        } else {
          resp = await this.$store.dispatch("getContentById", id);
        }
        this.content = resp.content;
        this.comments = resp.comments;
        this.writer = resp.writer;
        this.favorite = resp.favorite;
      } catch (e) {
        if (e instanceof RespError) {
          this.$message.error(this.$t("message.common.UE"));
        }
      }
    },
    async toggleFavorite() {
      let uid = this.$store.getters.profile.id;
      let cid = this.content.id;
      let resp;
      if (this.favorite) {
        resp = await this.$store.dispatch("deleteFavorite", { uid, cid });
      } else {
        resp = await this.$store.dispatch("createFavorite", { uid, cid });
      }
      if (resp == "success") {
        this.favorite = !this.favorite;
        this.$notify.info({
          message: this.$t("message.article.OS")
        });
      } else {
        this.$notify.error({
          message: this.$t("message.article.OF")
        });
      }
    }
  },
  computed: {
    category() {
      return this.$route.params["cat"];
    },
    cid() {
      return this.$route.params["id"];
    },
    isLogin() {
      return this.$store.state.isLogin;
    }
  },
  created() {
    this.fetchContent(this.cid);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="stylus">
.box-card {
  background-color: lightgrey;
}

.bread-nav {
  padding: 10px;
}

.el-divider__text {
  background-color: rgba(lightgrey, 0.5);
}

.el-divider {
  background-color: #fff;
}
.el-card .is-never-shadow {
    border:0;
    max-height :200px;
    overflow-y :scroll;
}
.comment-panel * {
  margin-bottom: 5px;
}

.content-info {
  margin-right: 20px;
}
</style>
