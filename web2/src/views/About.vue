<template>
  <div class="about">
    <el-tabs v-model="activeName">
      <el-tab-pane :label="$t('message.about.TL')" name="timeline" >
        <h1 style="text-align: center;">{{ $t('message.about.TL') }}</h1>
        <el-timeline>
          <el-timeline-item timestamp="2019/10/5" placement="top">
            <el-card>
              <h4>{{ $t('message.about.DA') }}</h4>
              <el-link href="https://www.processon.com/diagraming/5c9a3b59e4b07103434c03aa">2019/10/5 {{ $t('message.about.ERD') }}</el-link><br>
              <el-link href="https://www.processon.com/diagraming/5ca8291be4b029f6dae832a0">2019/10/6 {{ $t('message.about.SAD') }}</el-link><br>
              <el-link href="https://www.processon.com/diagraming/5c9a3acee4b0c7746794db0b">2019/10/6 {{ $t('message.about.UCD') }}</el-link>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2019/10/15" placement="top">
            <el-card>
              <h4>{{ $t('message.about.IP') }}</h4>
              <el-link style="text-decoration: line-through" href="https://github.com/Nonsensersunny/just-do-it" disabled>2019/10/15 {{ $t('message.about.SC') }} <el-tag type="danger" size="mini">{{ $t('message.about.DD') }}</el-tag></el-link>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2019/10/21" placement="top">
            <el-card>
              <h4>{{ $t('message.about.RC') }}</h4>
              <el-link href="https://github.com/Nonsensersunny/graduation">2019/10/21 {{ $t('message.about.SC') }}</el-link>
            </el-card>
          </el-timeline-item>
          <el-timeline-item timestamp="2019/11/1" placement="top">
            <el-card>
              <h4>{{ $t('message.about.DP') }}</h4>
              <el-link href="http://go.zyven.top">2019/11/1 {{ $t('message.about.LP') }}</el-link>
            </el-card>
          </el-timeline-item>
        </el-timeline>
      </el-tab-pane>
      <el-tab-pane :label="$t('message.about.STK')" name="stack">
        <el-card id="jishu" :body-style="{ padding: '0px' }">
          <img src="http://lc-bsf0iqqm.cn-n1.lcfile.com/ae39976bc073bffb4424/conslusion.png" class="image">
          <div style="padding: 14px;">
            <span>{{ $t('message.about.STK') }}</span>
            <div class="bottom clearfix">
              <time class="time">{{ currentDate }}</time>
            </div>
          </div>
        </el-card>
        <!-- <div style="text-align: center;">
          <el-image align="center" src="http://lc-bsf0iqqm.cn-n1.lcfile.com/ae39976bc073bffb4424/conslusion.png"></el-image>
        </div> -->
      </el-tab-pane>
    </el-tabs>
  </div>
</template>
<script>
  export default {
    name: "About",
    data() {
      return {
        currentDate: new Date(),
        activeName: "timeline"
      }
    }
  }
</script>
<style lang="stylus">
  .about {
    padding: 20px;
  }
  #pane-timeline{
    height :600px;
    overflow-y:scroll;
  }
  #pane-timeline::-webkit-scrollbar{
    display:none;
  }
  .el-card{
    background-color :rgba(255,255,255,1);
  }
  .el-timeline-item__node{
    background-color :#000;
    margin :auto;
  }
  .el-timeline-item__tail{
    border-left :2px solid #000;
    margin :auto;
  }
  .el-timeline-item__timestamp{
    color :#000;
    margin :auto;

  }
    .time {
    font-size: 13px;
    color: #999;
  }
 #jishu{
   width :950px;
   margin :auto;
 }
  .bottom {
    margin-top: 13px;
    line-height: 12px;
    margin-right :25px;
    margin-bottom :15px;
  }
  .image {
    margin :auto;
    display: block;
    border-bottom :1px solid #000;
  }
  .clearfix:before,
  .clearfix:after {
      display: table;
      content: "";
  }
  .clearfix:after {
      clear: both
  }
</style>