<template>
  <div class="home">
    <!-- <Gcontent :content="mockContent" v-if="!isLogin" class="main-content" /> -->
    <Gcontent :content="mockContent" v-if="isLogin" class="main-content" />
    <el-container v-else>
      <el-main><Gcontent :content="mockContent" /></el-main>
      <el-aside width="230px" style="max-height:700px; overflow-y:scroll;" class="side-board">
        <Gdashboard activePath="share" />
        <Grank />
        <Gtop />
        <Glinks />
      </el-aside>
    </el-container>
    <el-footer>
      <Gfooter />
    </el-footer>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import Gcontent from '@/components/Gcontent.vue'
import Grank from '@/components/Grank.vue'
import Gtop from '@/components/Gtop.vue';
import Glinks from '@/components/Glinks.vue';
import Gdashboard from '@/components/Gdashboard.vue';
import Gfooter from '@/components/Gfooter.vue';

export default {
  name: 'Home',
  components: {
    HelloWorld, Gcontent, Grank, Gtop, Glinks, Gdashboard, Gfooter
  },
  data() {
    return {
      mockContent: {
        title: "Mock title",
        time: new Date().toLocaleString("zh"),
        author: "Mock author",
        content: "Mock content",
        views: 0,
        category: "Mock category"
      }
    }
  },
  computed: {
    isLogin() {
      return this.$store.state.isLogin
    }
  }
}
</script>
<style lang="stylus" scoped>
  .side-board {

  }

  .side-board * {
    margin-top: 10px;
    padding: 10px 0 10px 0;
  }
  .main-content {
    margin: 5px 5px 20px 5px;
  }
  
.el-aside::-webkit-scrollbar {
  display :none;
}
</style>
