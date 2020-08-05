 <template>
  <div style="width:100%;height:100%;" v-loading="loading">
    <el-page-header class="page-header" @back="goBack" content="标记"></el-page-header>
    <div class="label-core" tabindex="1">
      <div class="label-toolbar">
        <ul>
          <a class="select" title="选择(A)" href="javascript:void(0)">
            <li class="active">
              <span class="iconfont icon-zhizhen"></span>
            </li>
          </a>
          <a class="move" title="移动/缩放(W)" href="javascript:void(0)">
            <li>
              <span class="iconfont icon-yidong"></span>
            </li>
          </a>
          <!--原始版本
          <a
            :style="{ display: type == 3 || (type == 4 && this.labelIndex != this.label.length) ? 'block' : 'none' }"
            class="polygon"
            title="添加多边形区域(Q)"
            href="javascript:void(0)"
          >
            <li>
              <span class="iconfont icon-duobianxing"></span>
            </li>
          </a>
          -->
          <a
            :style="{ display: type==2 || type == 3 || (type == 4 && this.labelIndex != this.label.length) ? 'block' : 'none' }"
            class="polygon"
            title="添加多边形区域(Q)"
            href="javascript:void(0)"
          >
            <li>
              <span class="iconfont icon-duobianxing"></span>
            </li>
          </a>

          <a
            :style="{ display: type == 2 || (type == 4 && this.labelIndex != this.label.length) ? 'block' : 'none' }"
            class="rectangle"
            title="添加正方形区域(R)"
            href="javascript:void(0)"
          >
            <li>
              <span class="iconfont icon-kuangxuanquyu"></span>
            </li>
          </a>
          <div class="label-divider"></div>
          <a :style="{ display: type != 1 && type != 4 ? 'block' : 'none' }" class="edit" title="重新编辑区域(E)" href="javascript:void(0)">
            <li>
              <span class="iconfont icon-sheji"></span>
            </li>
          </a>
          <a @click="classifyClick" v-if="type == 1" class="classify" title="添加分类标签" href="javascript:void(0)">
            <li>
              <span class="iconfont icon-sheji"></span>
            </li>
          </a>
          <a @click="nextLabel" v-if="type == 4" class="classify" title="下一个标签" href="javascript:void(0)">
            <li>
              <span class="iconfont icon-nextweek"></span>
            </li>
          </a>
          <a :style="{ display: type != 1 ? 'block' : 'none' }" class="delete" title="删除区域(D)" href="javascript:void(0)">
            <li>
              <span class="iconfont icon-shanchu"></span>
            </li>
          </a>
          <div class="label-divider"></div>
          <a class="zoomin" title="放大(Z)" href="javascript:void(0)">
            <li>
              <span class="iconfont icon-fangda"></span>
            </li>
          </a>
          <a class="zoomout" title="缩小(X)" href="javascript:void(0)">
            <li>
              <span class="iconfont icon-suoxiao"></span>
            </li>
          </a>

          <div class="label-divider"></div>
          <a
            :style="{ display: account.group == 3 ? 'block' : 'none' }"
            @click="showDelete"
            class="delete" title="删除" href="javascript:void(0)" >
            <li>
              <span class="el-icon-circle-close"></span>
            </li>
          </a>
          <a class="save" title="保存(Ctrl+S)" href="javascript:void(0)">
            <li>
              <span class="iconfont icon-baocun"></span>
            </li>
          </a>
          <a
            :style="{ display: account.group == 2 || account.group == 3 ? 'block' : 'none' }"
            class="confirm"
            @click="showComfirm"
            title="确认该版本(Ctrl+C)"
            href="javascript:void(0)"
          >
            <li>
              <span class="iconfont icon-dui"></span>
            </li>
          </a>
          <div class="label-divider"></div>
          <a @click="previous" class="zoomin" title="上一个" href="javascript:void(0)">
            <li>
              <span class="iconfont icon-ai10"></span>
            </li>
          </a>
          <a @click="next" class="zoomout" title="下一个" href="javascript:void(0)">
            <li>
              <span class="iconfont icon-ai09"></span>
            </li>
          </a>
        </ul>
      </div>
      <div class="label-area">
        <div class="label-main">
          <img v-on:load="loadMark" :src="img" alt srcset />

          <canvas></canvas>
          <svg class="label-svg" xmlns="http://www.w3.org/2000/svg" version="1.1" />

        </div>
      </div>
      <div class="label-popup" style="display: none;">
        <select class="label-label" multiple></select>
        <!-- <el-button type="primary" disabled>多选</el-button> -->
        <!-- <el-select class="multi-label-label" v-model="value1" multiple placeholder="请选择">
          <el-option
            v-for="item in label"
            :key="item.id"
            :label="item.label"
            :value="item.id">
          </el-option>
        </el-select> -->
        
        <div class="label-checkbox">
          <input type="checkbox" class="label-crowd" value="iscrowd" />
          <label for="iscrowd">使用RLE格式</label>
        </div>
        <textarea class="label-desc" rows="5" placeholder="描述"></textarea>
        <button class="label-save-label" @click="print" type="button">保存</button>
        <button class="label-delete-label" type="button">删除</button>

      </div>

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

        <ul class="label-list-ul" ></ul>


        <el-card  body-style="{ padding: '30px'}" shadow="always" class="box-card">
          <div v-if="task.list[task.index].userConfirmId != null" class="yishenghe">
            <h3>已审核</h3>
          </div>
          <div v-else class="weishenghe">
            <h3>未审核</h3>
          </div>
        </el-card>




      </div>

    </div>
    <!-- 类型 -->
    <!-- <el-dialog title="添加分类标签" :visible.sync="classify.dialogVisible" width="30%">
      <el-form :inline="true">
        <el-form-item label="标签" prop="group">
          <el-select v-model="classify.select" multiple collapse-tags style="margin-left: 20px;" placeholder="请选择标签">
            <el-option v-for="item in classify.label" :key="item.value" :label="item.label" :value="item.value"> </el-option>
          </el-select>
        </el-form-item>
      </el-form>

      <span slot="footer" class="dialog-footer">
        <el-button @click="classify.dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="classifyAdd">确 定</el-button>
      </span>
    </el-dialog> -->
  </div>
</template>
<script>
import { getImg, getPendingUserList, setFinalVersion, saveLabel, deleteImageById} from "@/models/Task.js";
import { LabelCore } from "@/assets/js/label.core.js";
import { HOST as Host, ajaxImg } from "@/models/Service.js";
import { getImgList } from "@/models/Task.js";
export default {
  name: "label",
  props: ["id", "tid", "type", "taskName", "imagePage"],
  data() {
    return {
      img: "",
      label: [],
      select: [],
      value1: [],
      name: "",
      area: [],
      count: 0,
      showComfirmed: true,
      showDeleted:false,
      penddingUserList: [],
      isReviewer: false,
      loading: true,
      penddingUserSelectId: 0,
      labelCore: undefined,
      penddingUserActive: "user",
      account: this.global.account,
      activeIndex: "",
      task: {
        page: +this.imagePage,
        total: 0,
        id: this.tid,
        list: [],
        index: 0,
        name: this.taskName
      },
      classify: {
        label: [],
        select: [],
        dialogVisible: false
      },
      labelIndex: 0,
      options: [{
        value: '选项1',
        label: '黄金糕'
      }, {
        value: '选项2',
        label: '双皮奶'
      }, {
        value: '选项3',
        label: '蚵仔煎'
      }, {
        value: '选项4',
        label: '龙须面'
      }, {
        value: '选项5',
        label: '北京烤鸭'
      }]
    }
  },
  watch: {
    id() {
      this.load(true);
    }
  },
  methods: {
    print(){
      console.log(this.options.value1);
      console.log(this.label);
      console.log(this.value1);
    },
    showDelete(){
      console.log("delete");
      this.loadMark();
      this.$confirm("是否删除该图片？", "提示", {confirmButtonText: "确定", cancelButtonText: "取消", type: "info"})
        .then(() => {
            this.deleteImageById();
          // this.setFinalVersion();
        })
        .catch(() => {
        });

    },
    showComfirm(){
      console.log("func");
      this.$confirm("是否确定该版本为最终版本？", "提示", {confirmButtonText: "确定", cancelButtonText: "取消", type: "info"})
        .then(() => {
          this.setFinalVersion();
          // this.setFinalVersion();
        })
        .catch(() => {
        });
    },
    format(percentage) {
      return percentage === 100 ? '满' : `${percentage}%`;
    },
    load(hasCom) {
      if (this.account.group == 1) {
        this.getImgInfo(hasCom);
      } else if (this.account.group == 3 || this.account.group == 2) {
        if (hasCom) {
          this.getImgInfo(hasCom);
        } else {
          this.getUserLabelList(hasCom);
        }
      } else {
        this.$message({ type: "error", message: `尚未登录` });
      }
      if (this.account.group != 0 && !hasCom) {
        this.getImageList()
          .then()
          .catch();
      }
    },
    //获取图片信息
    getImgInfo(isReload) {
      this.loading = true;
      getImg(this.id, this.account.group == 1 ? this.account.id : this.penddingUserSelectId, this.account)
        .then(data => {
          let src = Host + `/api/image/${this.taskName}/${data.data.image.imageName}`;
          if (this.img == src) {
            this.loading = false;
          }
          let str = JSON.stringify(data.data.labels);
          str = str
            .replace(/labelId/g, "id")
            .replace(/labelName/g, "label")
            .replace(/labelType/g, "type")
            .replace(/labelColor/g, "color");
          this.label = JSON.parse(str);
          this.name = data.data.image.imageName;
          this.area = [];
          this.classify.select = [];
          this.labelIndex = 0;
          let templabel = "";
          for (const item of data.data.datas) {
            let points = [];

            for (const point of item.point) {
              points.push([point.x, point.y]);
            }
            this.area.push({
              labelId: item.labelId,
              labelType: item.labelType,
              desc: item.dataDesc,
              points: points,
              iscrowd: item.iscrowd
            });
            if (this.type == 1) {
              this.classify.select.push(item.labelId);
            }
            if (templabel != item.labelId) {
              this.labelIndex++;
              templabel = item.labelId;
            }
          }
          if (this.img == src) {
            this.labelCore.reload(this.area);
            this.loadLabelIndex();
            return;
          }
          this.img = src;
        })
        .catch(data => {
          this.loading = false;
          this.$alert(`获取图片信息失败，${data.message}(${data.code})。\r\n请重试`, "提示", { confirmButtonText: "确定", type: "warning" });
        });
    },
    //获取审核用户列表
    getUserLabelList(hasCom) {
      getPendingUserList(this.id, this.account)
        .then(d => {
          let data = d.data;
          if (data.length > 0) {
            this.activeIndex = data[0].userId + "";
            this.penddingUserSelectId = data[0].userId;
            this.getImgInfo(hasCom);
            this.penddingUserList = data;
          } else {
            //prettier-ignore
            this.$alert("该任务未分配用户标记", "提示", {confirmButtonText: "确定",callback: action => {this.$router.go(-1);}});
          }
        })
        .catch(data => {
          this.$message({ type: "error", message: `获取用户标记列表失败，${data.message}(${data.code})` });
        });
    },
    //渲染标记区域(调用label.core)
    loadMark() {
      if (this.img == "") {
        return;
      }
      this.loading = false;
      if (this.labelCore != null) {
        this.labelCore.reload(this.area);

        this.loadLabelIndex();
        return;
      }
      this.labelCore = new LabelCore(
        document.querySelector(".label-core"),
        this.label,
        this.value1,
        this.area,
        data => {
          this.saveLabelArea(data);
        },
        data => {
          // this.$confirm("是否确定该版本为最终版本？", "提示", {confirmButtonText: "确定", cancelButtonText: "取消", type: "info"})
          //   .then(() => {
          //     this.setFinalVersion();
          //     // this.setFinalVersion();
          //   })
          //   .catch(() => {
          //   });
        },
        //addCB
        () => {
        },
        //removeCB
        () => {
          if (this.type == 4) {
            this.labelCore.setLabelIndex(--this.labelIndex);
            this.$message({ type: "info", message: `你已经删除上一个标签的所有标记，现在请标记标签:'${this.label[this.labelIndex].label}'` });
          }
        }
      );
      this.loadLabelIndex();
    },
    loadLabelIndex() {
      if (this.type == 4) {
        console.log(this.labelIndex);
        this.labelCore.setLabelIndex(this.labelIndex);
        if (this.labelIndex == this.label.length) {
          this.$message({ type: "success", message: `当前图片已经完成标注` });
          return;
        }
        this.$message({ type: "info", message: `现在请标记标签:'${this.label[this.labelIndex].label}'` });
      }
    },
    nextLabel() {
      this.labelCore.setLabelIndex(++this.labelIndex);
      if (this.labelIndex >= this.label.length) {
        this.labelIndex = this.label.length;
        this.$message({ type: "success", message: `已经完成当前图片标注，请点击保存结束标记` });
        return;
      }
      this.$message({ type: "info", message: `现在请标记标签:'${this.label[this.labelIndex].label}'` });
    },
    //返回
    goBack() {
      this.$confirm("未保存的进度将会丢失，确定返回？", "提示", { confirmButtonText: "确定", cancelButtonText: "取消", type: "warning" })
        .then(() => {
          this.$router.push(`/task/${this.tid}/${this.taskName}/${this.type}`);
        })
        .catch(() => {});
    },
    //选择审核用户event
    penddingUserSelect(index) {
      this.penddingUserSelectId = parseInt(index);
      this.getImgInfo(true);
    },
    //设置最终版本
    setFinalVersion() {
      this.loading = true;
      setFinalVersion(parseInt(this.id), this.penddingUserSelectId, this.account)
        .then(() => {
          this.$message({ type: "success", message: "设置最终版本成功" });
        })
        .catch(data => {
          this.$message({ type: "error", message: `设置最终版本失败，${data.message}(${data.code})` });
        })
        .then(() => {
          this.loading = false;
        });
    },
    deleteImageById(){
      this.loading = true;
      deleteImageById(parseInt(this.id), this.penddingUserSelectId, this.account)
        .then(() => {
          this.$message({ type: "success", message: "删除成功" });
          this.next();
        })
        .catch(data => {
          this.$message({ type: "error", message: `删除失败，${data.message}(${data.code})` });
        })
        .then(() => {
          this.loading = false;
        });

    },
    //保存
    saveLabelArea(data) {
      this.loading = true;
      let postData = [];
      for (const item of data) {
        let points = [],
          i = 0;
        for (const point of item.points) {
          points.push({
            x: point[0],
            y: point[1],
            order: ++i
          });
        }
        postData.push({
          labelId: item.id,
          labelType: item.type,
          dataDesc: item.desc,
          iscrowd: item.iscrowd,
          point: points
        });
      }
      let ps = saveLabel(this.id, postData, this.account.group == 1 ? this.account.id : this.penddingUserSelectId, this.account);
      ps.then(() => {
        this.$message({ type: "success", message: "保存成功" });
      })
        .catch(data => {
          this.$message({ type: "error", message: `保存失败，${data.message}(${data.code})` });
        })
        .then(() => {
          this.loading = false;
        });
    },
    //获取任务图片列表，确定
    getImageList(isNew) {
      return new Promise((resolve, reject) => {

        console.log(this.task.page);
        getImgList(this.task.id, this.task.page)
          .then(data => {
            let list = data.data.images;
            this.task.list = list;
            this.task.total = data.data.totalpages;
            if (!isNew) {
              for (let i = 0; i < list.length; i++) {
                if (list[i].imageId == this.id) {
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
        console.log("这是这个页面最后一张");
        if (task.page == task.total) {
          this.$message({ type: "warning", message: `已经是最后一张` });
          return;
        }
        console.log(task.page);
        task.page++;
        console.log(task.page);
        task.index = 0 ;
        console.log(task.list[task.index].imageId);
        getImgList(this.task.id, this.task.page)
          .then(data => {
            let list = data.data.images;
            this.task.list = list;
            console.log("success");
            console.log(task.list[task.index].imageId);
            console.log(task.list[task.index].userConfirmId);
            this.$router.push(`/img/${task.list[task.index].imageId}/${task.id}/${this.type}/${task.name}/${task.page}`);
          });
        console.log(task.list[task.index].imageId);
        console.log(task.list);
        //this.$router.push(`/img/${task.list[task.index].imageId}/${task.id}/${this.type}/${task.name}/${task.page}`);
      } else {
        console.log(task.list);
        ++task.index;
        console.log(task.list[task.index].imageId);
        this.$router.push(`/img/${task.list[task.index].imageId}/${task.id}/${this.type}/${task.name}/${task.page}`);
      }
    },
    previous() {
      let task = this.task;
      if (task.index == 0) {
        if (task.page == 1) {
          this.$message({ type: "warning", message: `已经是第一张` });
          return;
        }
        console.log(task.page);
        task.page--;
        console.log(task.page);

        getImgList(this.task.id, this.task.page)
          .then(data => {
            let list = data.data.images;
            this.task.list = list;
            task.index = task.list.length - 1;
            console.log("success");
            console.log(task.list[task.index].imageId);
            this.$router.push(`/img/${task.list[task.index].imageId}/${task.id}/${this.type}/${task.name}/${task.page}`);
          });

      } else {
        --task.index;
        this.$router.push(`/img/${task.list[task.index].imageId}/${task.id}/${this.type}/${task.name}/${task.page}`);
      }
    }
  },
  mounted() {
    this.load();
  }
};
</script>
<style src="@/assets/css/label.core.css"></style>




