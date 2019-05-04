<template>
  <div class="about">
    <el-tabs v-model="activeName">
      <el-tab-pane :label="$t('message.about.TL')" name="timeline">
        <h1 style="text-align: center;">{{ $t('message.about.TL') }}</h1>
        <el-timeline>
          <el-timeline-item timestamp="2019/3/20" placement="top">
            <el-card>
              <h4>{{ $t('message.about.DA') }}</h4>
              <el-link href="https://www.processon.com/diagraming/5c9a3b59e4b07103434c03aa">2019/3/20 {{ $t('message.about.ERD') }}</el-link><br>
              <el-link href="https://www.processon.com/diagraming/5ca8291be4b029f6dae832a0">2019/3/21 {{ $t('message.about.SAD') }}</el-link><br>
              <el-link href="https://www.processon.com/diagraming/5c9a3acee4b0c7746794db0b">2019/3/21 {{ $t('message.about.UCD') }}</el-link>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2019/3/25" placement="top">
            <el-card>
              <h4>{{ $t('message.about.IP') }}</h4>
              <el-link style="text-decoration: line-through" href="https://github.com/Nonsensersunny/just-do-it" disabled>2019/3/25 {{ $t('message.about.SC') }} <el-tag type="danger" size="mini">{{ $t('message.about.DD') }}</el-tag></el-link>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2019/4/21" placement="top">
            <el-card>
              <h4>{{ $t('message.about.RC') }}</h4>
              <el-link href="https://github.com/Nonsensersunny/graduation">2019/4/21 {{ $t('message.about.SC') }}</el-link>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2019/5/1" placement="top">
            <el-card>
              <h4>{{ $t('message.about.DP') }}</h4>
              <el-link href="http://go.zyven.top">2019/5/1 {{ $t('message.about.LP') }}</el-link>
            </el-card>
          </el-timeline-item>
        </el-timeline>
      </el-tab-pane>
      <el-tab-pane :label="$t('message.about.STK')" name="stack">
        <h1 style="text-align: center;">{{ $t('message.about.STK') }}</h1>
        <div style="text-align: center;">
          <el-image align="center" src="http://lc-bsf0iqqm.cn-n1.lcfile.com/ae39976bc073bffb4424/conslusion.png"></el-image>
        </div>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
  export default {
    name: "About",
    data() {
      return {
        activeName: "timeline"
      }
    }
  }
</script>
<style lang="stylus">
  .about {
    padding: 20px;
  }
</style>