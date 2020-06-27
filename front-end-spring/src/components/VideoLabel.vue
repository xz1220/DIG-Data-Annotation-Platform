<template>
  <div style="width:100%;height:100%;" v-loading="loading">
    <el-page-header class="page-header" @back="goBack" content="视频标记"></el-page-header>
    <div class="label-video">
      <div class="bfplayer" tabindex="0">
        <div style="display: none;" class="bf-loading"></div>
        <div class="bf-main">
          <div class="bf-loading-panel">
            <div class="bf-loading">
              <div class="bf-ball bf-ball-1"></div>
              <div class="bf-ball bf-ball-2"></div>
              <div class="bf-ball bf-ball-3"></div>
              <div class="bf-ball bf-ball-4"></div>
              <div class="bf-ball bf-ball-5"></div>
            </div>
          </div>
          <video class="bf-video" src=""></video>
        </div>
        <div class="bf-capture">
          <div role="buttom" class="bf-btn bf-btn-start bf-tooltip bf-tooltip-left" data-tooltip="设置起点">
            <span class="bf-icon bf-icon-xiazai"></span>
          </div>
          <div role="buttom" class="bf-btn bf-btn-end bf-tooltip bf-tooltip-left" data-tooltip="设置终点">
            <span class="bf-icon bf-icon-shangchuan"></span>
          </div>
          <div role="buttom" class="bf-btn bf-btn-save bf-tooltip bf-tooltip-left" data-tooltip="保存标记">
            <span class="bf-icon bf-icon-save"></span>
          </div>
          <div
            role="buttom"
            :style="{ display: account.group == 2 || account.group == 3 ? 'block' : 'none' }"
            class="bf-btn bf-tooltip bf-tooltip-left"
            @click="setFinalVersion"
            data-tooltip="确认最终版本"
          >
            <span class="bf-icon bf-icon-wancheng"></span>
          </div>
        </div>
        <div style="display: none;left:0;top:0" class="bf-menu">
          <ul>
            <li>
              <div>
                BFPlayer-Core By q6q64399
              </div>
            </li>
            <li>
              <div>
                1.1.055 (2019/10/13 22:53:25)
              </div>
            </li>
          </ul>
        </div>
        <div class="bf-controller">
          <div class="bf-pgb-player">
            <div style="display: none;" class="bf-pgb-point">
              <div></div>
            </div>
            <div class="bf-pgb-buffer"></div>
            <div class="bf-pgb-played"></div>
            <div class="bf-pgb-select"></div>
          </div>
          <div class="bf-top">
            <div role="buttom" class="bf-btn bf-btn-play"><span class="bf-icon bf-icon-play"></span></div>
            <div role="buttom" class="bf-btn bf-btn-previous bf-tooltip" data-tooltip="上一个"><span class="bf-icon bf-icon-previous"></span></div>
            <div role="buttom" class="bf-btn bf-btn-next bf-tooltip" data-tooltip="下一个"><span class="bf-icon bf-icon-next"></span></div>
            <div role="buttom" class="bf-btn bf-btn-volume"><span class="bf-icon bf-icon-volumeup"></span></div>
            <div class="bf-volume-area" style="display: none;">
              <div class="bf-volume">
                <div class="bf-volume-value" style="width: 100%;"></div>
                <div class="bf-volume-point" style="left:100%;"></div>
              </div>
            </div>
            <span class="bf-label-time">0:00 / 0:00</span>
            <span style="display: none" class="bf-label-select-time">已选择：dd - dd</span>
            <div class="bf-spacer"></div>
            <div
              class="bf-selector bf-selector-quality"
              data-selector-event="changeRate"
              data-selector='[{"name":"0.5x","value":"0.5"},{"name":"0.75x","value":"0.75"},{"name":"1x","value":"1","active":1},{"name":"1.25x","value":"1.25"},{"name":"1.5x","value":"1.5"},{"name":"2x","value":"2"}]'
            >
              倍速
            </div>
            <div role="buttom" class="bf-btn bf-btn-fullbrowser bf-tooltip" data-tooltip="网页全屏"><span class="bf-icon bf-icon-mini"></span></div>
            <div role="buttom" class="bf-btn bf-btn-fullscreen bf-tooltip" data-tooltip="屏幕全屏">
              <span class="bf-icon bf-icon-fullscreen"></span>
            </div>
          </div>
        </div>
      </div>
      <div class="label-right">
        <div class="label-list">
          <el-collapse :style="{ display: account.group != 1 ? 'block' : 'none' }" v-model="penddingUserActive" class="label-user">
            <el-collapse-item class="label-user-item" title="待审核用户列表" name="user">
              <el-menu @select="penddingUserSelect" :default-active="activeIndex" class="el-menu-vertical-demo">
                <el-menu-item v-for="item in penddingUserList" :key="item.userId" :index="item.userId + ''">
                  <span slot="title">{{ item.username }}</span>
                </el-menu-item>
              </el-menu>
            </el-collapse-item>
          </el-collapse>
          <ul class="label-list-ul"></ul>
        </div>
      </div>
      <div style="display: none;" class="label-sign">
        <div class="label-sign-question">
          添加描述
        </div>
        <div class="label-sign-answer">
          <select style="display: none;" class="label-sign-selector"> </select>
          <textarea class="label-sign-text"></textarea>
        </div>
        <div class="label-sign-bottom">
          <button class="label-sign-btn label-sign-submit" type="button">保存</button>
          <button class="label-sign-btn label-sign-cancel" type="button">取消</button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import BFPlayer from "@/assets/js/bfplayer.core";
import VideoLabel from "@/assets/js/videolabel.core";
import Video from "@/models/Video.js";
import { HOST as Host } from "@/models/Service.js";
export default {
  name: "VideoLabel",
  props: ["id", "tid", "taskName", "imagePage"],
  data() {
    return {
      loading: true,
      activeIndex: "",
      penddingUserList: [],
      penddingUserSelectId: 0,
      penddingUserActive: "user",
      task: {
        page: +this.imagePage,
        total: 0,
        id: this.tid,
        list: [],
        index: 0,
        name: this.taskName
      },
      videolabel: null,
      account: this.global.account
    };
  },
  watch: {
    id() {
      this.load(true);
    }
  },
  methods: {
    load(reload) {
      if (this.account.group == 1) {
        this.getVideoInfo();
      } else if (this.account.group == 3 || this.account.group == 2) {
        if (reload) {
          this.getVideoInfo();
        } else {
          this.getReviewerList();
        }
      } else {
        this.$message({ type: "error", message: `尚未登录` });
      }
      //获取任务视频列表，next,previous备用
      if (this.account.group != 0 && !reload) {
        this.getList()
          .then()
          .catch();
      }
    },
    getVideoInfo() {
      this.loading = true;
      Video.getVideo(this.id, this.account.group == 1 ? this.account.id : this.penddingUserSelectId).then(data => {
        let src = Host + `/api/video/${this.taskName}/${data.data.video.videoName}`;
        console.log(data);
        this.loading = false;
        if (this.videolabel) {
          this.videolabel.reload(data.data.datas, src);
          return;
        }
        this.videolabel = VideoLabel(
          BFPlayer,
          {
            labelId: data.data.labels[0].labelId,
            question: data.data.labels[0].question,
            type: data.data.labels[0].type,
            selector: typeof data.data.labels[0].selector == "string" ? eval(data.data.labels[0].selector) : data.data.labels[0].selector
          },
          data.data.datas,
          src,
          () => {
            this.loading = true;
            Video.saveVideoLabel(this.id, this.videolabel.save(), this.account.group == 1 ? this.account.id : this.penddingUserSelectId)
              .then(() => {
                this.$message({ type: "success", message: `保存成功` });
              })
              .catch(() => {
                this.$message({ type: "error", message: `保存失败` });
              })
              .then(() => {
                this.loading = false;
              });
          },
          () => {
            this.next();
          },
          () => {
            this.previous();
          }
        );
      });
    },
    getReviewerList() {
      Video.getPendingUserList(this.id).then(d => {
        let data = d.data;
        if (data.length > 0) {
          this.activeIndex = data[0].userId + "";
          this.penddingUserSelectId = data[0].userId;
          this.getVideoInfo();
          this.penddingUserList = data;
        } else {
          //prettier-ignore
          this.$alert("该任务未分配用户标记", "提示", {confirmButtonText: "确定",callback: action => {this.$router.go(-1);}});
        }
      });
    },
    //选择审核用户event
    penddingUserSelect(index) {
      this.penddingUserSelectId = parseInt(index);
      this.getVideoInfo(true);
    },
    setFinalVersion() {
      this.loading = true;
      Video.setVideoFinalVersion(parseInt(this.id), this.penddingUserSelectId)
        .then(() => {
          this.$message({ type: "success", message: `设置成功` });
        })
        .catch(() => {
          this.$message({ type: "error", message: `设置失败` });
        })
        .then(() => {
          this.loading = false;
        });
    },
    goBack() {
      this.$confirm("未保存的进度将会丢失（删除操作也需要保存），确定返回？", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          this.$router.push(`/task/${this.tid}/${this.taskName}/5`);
        })
        .catch(() => {});
    },
    //获取任务图片列表，确定
    getList(isNew) {
      return new Promise((resolve, reject) => {
        Video.getVideoList(this.task.id, this.task.page)
          .then(data => {
            let list = data.data.videos;
            this.task.list = list;
            this.task.total = data.data.totalpages;
            if (!isNew) {
              for (let i = 0; i < list.length; i++) {
                if (list[i].videoId == this.id) {
                  this.task.index = i;
                }
              }
            }
            resolve();
          })
          .catch(() => {
            this.$message({ type: "error", message: `获取列表失败` });
            reject();
          });
      });
    },
    next() {
      let task = this.task;
      if (task.list.length - 1 == task.index) {
        if (task.page == task.total) {
          this.$message({ type: "warning", message: `已经是最后一个` });
          return;
        }
        task.page++;
        this.getVideoInfo().then(() => {
          task.index = 0;
          this.$router.push(`/video/${task.list[task.index].videoId}/${task.id}/${task.name}/${task.page}`);
        });
      } else {
        console.log(task.list);
        ++task.index;
        this.$router.push(`/video/${task.list[task.index].videoId}/${task.id}/${task.name}/${task.page}`);
      }
    },
    previous() {
      let task = this.task;
      if (task.index == 0) {
        if (task.page == 1) {
          this.$message({ type: "warning", message: `已经是第一个` });
          return;
        }
        task.page--;
        this.getVideoInfo().then(() => {
          task.index = task.list.length - 1;
          this.$router.push(`/video/${task.list[task.index].videoId}/${task.id}/${task.name}/${task.page}`);
        });
      } else {
        --task.index;
        this.$router.push(`/video/${task.list[task.index].videoId}/${task.id}/${task.name}/${task.page}`);
      }
    }
  },
  mounted() {
    this.load();
  }
};
</script>
<style src="@/assets/css/bfplayer.core.css"></style>
<style src="@/assets/css/videolabel.core.css"></style>


