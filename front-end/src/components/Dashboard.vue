<template>
  <div>
    <template v-if="account.group == 3">

      <div class="dashboard" v-loading="loading">
        <div class="dashboard-panel panel-number">

          <div class="panel-header">
            总任务数
          </div>
          <div class="panel-content">
            {{ taskCount }}
          </div>
        </div>
        <div class="dashboard-panel panel-number">
          <div class="panel-header">
            总标记用户数
          </div>
          <div class="panel-content">
            {{ userCount }}
          </div>
        </div>
        <div class="dashboard-panel panel-number">
          <div class="panel-header">
            总审核用户数
          </div>
          <div class="panel-content">
            {{ reviewerCount }}
          </div>
        </div>
      </div>
    </template>
    <div class="公告">
      <h3 class="section-title">公告</h3>
      <p>
        2020-2-20 11:30pm 后开始维护。可能会出现服务断开，需要下载数据的请提前下载，感谢大家的配合 😀
      </p>
    </div>
    <div class="guide">
      <h3 class="section-title">图片标记指南</h3>
      <p>
        <b>关键点标记说明：</b><br/>
        <b>下一个标签：</b>点击<i class="iconfont icon-nextweek icon"></i>进行下一个标签的标记。<br />
        必须按逆序删除已经添加的标记，否则删除失败。<br/><br/>
        <b>移动/缩放（快捷键：W）：</b>点击<i class="iconfont icon-yidong icon"></i>可进行移动（鼠标按下移动）与缩放（鼠标滚轮滚动）。<br />
        <b>多边形选择（快捷键：Q）：</b>点击<i class="iconfont icon-duobianxing icon"></i
        >进行多边形选择模式，在图片上点击进行标记。左键单击第一个点完成标记；右键单击第一个点取消标记。<br />此时你仍然可以按下移动/缩放按钮进行操作。<br />
        <b>矩形选择（快捷键：R）：</b>点击<i class="iconfont icon-kuangxuanquyu icon"></i>进行矩形选择模式，在图片上某个点按下并拖动，完成选择。<br />
        <b>编辑（快捷键：E）：</b>点击<i class="iconfont icon-sheji icon"></i
        >进入编辑模式，点击一个已存在的选择区域进行修改，按下点拖动鼠标可移动点，点击图片其他区域结束修改。<br />
        <b>删除（快捷键：D）：</b>点击<i class="iconfont icon-shanchu icon"></i>进入删除模式，点击一个已存在的选择区域删除。<br />
        <b>放大（快捷键：Z）：</b>点击<i class="iconfont icon-fangda icon"></i>按左上角放大。<br />
        <b>放大（快捷键：C）：</b>点击<i class="iconfont icon-suoxiao icon"></i>按左上角缩小。<br />
        <b>保存（快捷键：Ctrl+S）：</b>点击<i class="iconfont icon-baocun icon"></i>保存。<br />
        <b>审核确认（快捷键：Ctrl+C）：</b>点击<i class="iconfont icon-dui icon"></i>确认选中版本（仅在管理员/审核用户模式下）。<br />
        <b>上一个：</b>点击<i class="iconfont icon-ai10 icon"></i>进入上一个标记项目。<br />
        <b>下一个：</b>点击<i class="iconfont icon-ai09 icon"></i>保存描述列表。<br />
        在编辑/选择模式下必须结束标记才可进行下一次标记。
      </p>
      <h3 class="section-title">视频标记指南</h3>
      <p>
        <b>播放（快捷键：SPACE）：</b>点击<i class="bf-icon bf-icon-play icon"></i>或者<i class="bf-icon bf-icon-pause icon"></i>开始播放/暂停。<br />
        <b>设置起点（快捷键：A）：</b>点击<i class="bf-icon bf-icon-xiazai icon"></i>设置标记描述起点。<br />
        <b>设置终点（快捷键：S）：</b>点击<i class="bf-icon bf-icon-shangchuan icon"></i>设置标记描述起点终点。<br />
        <b>快进（快捷键：→）：</b>点击进度条或使用快捷键进行快进操作（5秒）。<br />
        <b>快退（快捷键：←）：</b>点击进度条或使用快捷键进行快进操作（5秒）。<br />
        <b>网页全屏（快捷键：W）：</b>点击<i class="bf-icon bf-icon-mini icon"></i>进入或退出网页全屏。<br />
        <b>屏幕全屏（快捷键：F）：</b>点击<i class="bf-icon bf-icon-fullscreen icon"></i>进入或退出屏幕全屏<b>（屏幕全屏无法完成标记）</b>。<br />
        <b>上一个：</b>点击<i class="bf-icon bf-icon-previous icon"></i>进入上一个标记项目。<br />
        <b>下一个：</b>点击<i class="bf-icon bf-icon-next icon"></i>进入下一个标记项目。<br />
        <b>保存：</b>点击<i class="bf-icon bf-icon-save icon"></i>保存描述列表。<br />
        <b>审核确认：</b>点击<i class="bf-icon bf-icon-wancheng icon"></i>确认选中版本（仅在管理员/审核用户模式下）。<br />
        <b>注意：删除操作也需要点击保存。</b>
      </p>
    </div>
  </div>
</template>

<script>
import { getCount } from "@/models/Task";
export default {
  name: "Home",
  data() {
    return {
      account: this.global.account,
      userCount: 0,
      reviewerCount: 0,
      taskCount: 0,
      loading: false
    };
  },
  mounted() {
    if (this.account.group == 3) {
      this.loading = true;
      getCount(this.account)
        .then(data => {
          this.userCount = data.data.userCount;
          this.reviewerCount = data.data.reviewerCount;
          this.taskCount = data.data.taskCount;
        })
        .catch(data => {
          this.$message({ type: "error", message: "获取仪表信息失败" });
        })
        .then(() => {
          this.loading = false;
        });
    }
  }
};
</script>

<style scoped>
.dashboard {
  display: flex;
  flex-wrap: wrap;
}
.dashboard-panel {
  width: 200px;
  height: 180px;
  text-align: center;
  margin: 10px;
  border-radius: 5px;
  position: relative;
}
.panel-number {
  padding: 15px 5px;
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.11);
  background-color: #f9f9f9;
  color: #696969;
}
.panel-header {
  font-size: 20px;
}
.panel-content {
  font-size: 50px;
  position: absolute;
  left: 0;
  right: 0;
  top: 80px;
}
.guide > p {
  line-height: 2;
}
.icon {
  background-color: #eeee;
  border-radius: 3px;
  padding: 5px;
}
</style>
