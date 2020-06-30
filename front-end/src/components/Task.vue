<template>
  <div>
    <el-button  @click="showReviewer" size="mini">未标注</el-button>
    <el-alert :title="alert.title" :description="alert.desc" type="info" close-text="知道了"> </el-alert>

    <template v-if="id">
      <div class="img-list" v-loading="loading">
        <!--<el-button type="primary" class="img-list-download" icon="el-icon-download" circle @click="downloadClick"></el-button>-->
        <div
          class="img-item"
          v-for="img in imglist"
          :key="type != 5 ? img.imageId : img.videoId"
          @click="
            goto(
              img.userConfirmId,
              type != 5 ? `/img/${img.imageId}/${id}/${type}/${taskName}/${page.now}` : `/video/${img.videoId}/${id}/${taskName}/${page.now}`
            )
          "
        >
          <el-card shadow="hover" :body-style="{ padding: '0px' }" v-if="review ? !(labeledList[type != 5 ? img.imageId : img.videoId + '']): true">
            <el-image
              class="img-item-img"
              :lazy="true"
              :src="type != 5 ? `${imghost}/api/thumb/${taskName}/${img.imageThumb}` : `${imghost}/api/videos/${taskName}/${img.videoThumb}`"
              fit="cover"
            ></el-image>
            <div class="img-item-span">
              <span>
                {{ type != 5 ? img.imageName : img.videoName }}
                <el-tag type="success" size="mini" v-if="img.userConfirmId != null && img.userConfirmId != 0" class="img-item-finish">已审核</el-tag>
                <el-tag
                  type="primary"
                  size="mini"
                  v-if="labeledList[type != 5 ? img.imageId : img.videoId + '']"
                  class="img-item-labeled"
                >
                  已标记
                </el-tag>
              </span>
            </div>
          </el-card>
        </div>
      </div>
      <el-pagination
        v-if="page.total > 1"
        @current-change="pageChange"
        class="pagination"
        background
        layout="prev, pager, next,jumper"
        :page-count="page.total"
        :current-page.sync="page.now"
      ></el-pagination>
    </template>
    <h1 v-else>Illegal task id! Please refresh this page.</h1>
  </div>
</template>
<script>
import { getImgList } from "@/models/Task.js";
import Video from "@/models/Video.js";
import { HOST as Host } from "@/models/Service.js";
export default {
  name: "task",
  props: ["id", "taskName", "type"],
  data() {
    return {
      imglist: [],
      review:false,
      page: { now: 1, total: 1 },
      account: this.global.account,
      loading: false,
      imghost: Host,
      labeledList: {},
      notice: {
        type: ["分类", "检测", "分割", "关键点", "视频描述"],
        content: [
          "选择一个或多个恰当的标签对图片进行分类。",
          "只允许添加矩形框。",
          "只允许添加多边形。",
          "请按照给定的标签按顺序标记，按左边的‘下一个标签’按钮即可进入下一个标签标记。每个点都是有顺序的。",
          "对视频进行打点，添加描述。"
        ]
      },
      alert: {
        title: "",
        desc: ""
      }
    };
  },
  watch: {
    id: function() {
      this.page.now = 1;
      this.page.total = 1;
      this.show();
    }
  },
  methods: {
    showReviewer(){
          if(this.review){
            this.review=false
          }else {
            this.review=true
          }
    },
    show() {
      this.loading = true;
      let ps = this.type != 5 ? getImgList(this.id, this.page.now) : Video.getVideoList(this.id, this.page.now);

      ps.then(data => {
        this.labeledList = [];
        console.log("Start to doing something")
        for (const item of this.type != 5 ? data.data.labelImageIds : data.data.labelVideoIds) {
          this.labeledList[item + ""] = true;
          console.log("Add labeledImageID to labeledList")
        }
        this.imglist = this.type != 5 ? data.data.images : data.data.videos;
        this.page.total = data.data.totalpages;
      })
        .catch(data => {
          this.$alert(`获取列表失败，${data.message}(${data.code})。\r\n请重试`, "提示", {
            confirmButtonText: "确定",
            type: "warning"
          });
        })
        .then(() => {
          this.loading = false;
        });
    },
    pageChange(p) {
      this.show();
    },
    goto(id, src) {
      if (id == null || this.account.group != 1) {
        this.$router.push(src);
        return;
      }
      this.$alert(`该标注已经锁定`, "提示", {
        confirmButtonText: "确定",
        type: "warning"
      });
    }
  },
  mounted: function() {
    this.alert.title = this.notice.type[this.type - 1];
    this.alert.desc = this.notice.content[this.type - 1];
    this.show();
  }
};
</script>
<style scoped>
.img-list {
  display: flex;
  flex-wrap: wrap;
  width: 100%;
}
.img-item {
  flex-basis: 180px;
  flex-shrink: 0;
  margin: 10px;
  text-decoration: none;
  position: relative;
}
.img-item-img {
  height: 150px;
  width: 180px;
}
.desc {
  font-size: 13px;
  color: #8c8d8e;
  margin: 5px 0 0;
}
.img-item-finish {
  position: absolute;
  top: 10px;
  right: 10px;
}
.img-item-labeled {
  position: absolute;
  top: 10px;
  right: 60px;
}
.img-list-download {
  position: fixed;
  z-index: 1999;
  right: 50px;
  bottom: 30px;
}
.img-item-span {
  padding: 12px;
  width: 180px;
  box-sizing: border-box;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
}
.pagination {
  text-align: right;
  margin-right: 75px;
}
</style>
