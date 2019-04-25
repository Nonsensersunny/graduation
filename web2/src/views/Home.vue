<template>
  <div class="home">
    <!--<img alt="Vue logo" src="../assets/logo.png">-->
    <!--<HelloWorld msg="Welcome to Your Vue.js App"/>-->
    <el-container>
      <el-main><Gcontent :content="mockContent" /></el-main>
      <el-aside width="200px">
        <Grank />
        <Gtopic />
        <Glinks />
        <Gdashboard />
      </el-aside>
    </el-container>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import Gcontent from '@/components/Gcontent.vue'
import Grank from '@/components/Grank.vue'
import Gtopic from '@/components/Gtopic.vue';
import Glinks from '@/components/Glinks.vue';
import Gdashboard from '@/components/Gdashboard.vue';


export default {
  name: 'home',
  components: {
    HelloWorld, Gcontent, Grank, Gtopic, Glinks, Gdashboard
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

  }
}
</script>
