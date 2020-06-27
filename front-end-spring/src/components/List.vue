<template>
  <div>
    <el-button icon="el-icon-plus" @click="editUser" type="primary">下载任务</el-button>
    <el-button icon="el-icon-refresh-left" @click="showList">刷新列表</el-button>

    <el-table border :data="taskData" stripe style="width: 100%;margin-top:10px;" v-loading="loading">
      <el-table-column type="expand">
        <template slot-scope="props">
          <el-form label-position="left" inline class="table-expand">
            <el-form-item label="标注人员">
              <template v-if="props.row.users != null && props.row.users.length != 0">
                <span class="person" v-for="item in props.row.users" :key="item.userId">
                  {{ item.username }}({{ item.labeled }}/{{ props.row.imageNumber }})
                </span>
              </template>
              <span v-else>未指派标注人员</span>
            </el-form-item>
            <el-form-item label="审核人员">
              <template v-if="props.row.reviewers != null && props.row.reviewers.length != 0">
                <span class="person" v-for="item in props.row.reviewers" :key="item">
                  {{ item }}
                </span>
              </template>
              <span v-else>未指派审核人员</span>
            </el-form-item>
          </el-form>
        </template>
      </el-table-column>
      <el-table-column
        :filters="[
          { text: '分类', value: '1' },
          { text: '检测', value: '2' },
          { text: '分割', value: '3' },
          { text: '关键点检测', value: '4' },
          { text: '视频描述', value: '5' }
        ]"
        :filter-method="filterTag"
        filter-placement="bottom-end"
        prop="authorities"
        label="类型"
        width="180"
      >
        <template slot-scope="scope">
          {{ getTypeName(scope.row.taskType) }}
        </template>
      </el-table-column>
      <el-table-column prop="taskName" min-width="100" label="名称"></el-table-column>
      <el-table-column label="进度(已完成审核数/总图片数)">
        <template slot-scope="scope">
          {{ scope.row.finish }}/{{ scope.row.imageNumber }}
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" :width="account.group == 3 ? 220 : 80">
        <template slot-scope="scope">
          <el-button @click="enterTask(scope.row)" type="primary" size="mini">进入</el-button>
          <el-button v-if="account.group == 3" @click="downloadTask(scope.row)" size="mini">下载已完成项</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!--人员选择-->
    <el-dialog
      :title="userSelect.type == 'user' ? '下载任务' : '选择审核'"
      :visible.sync="userSelect.dialogVisible"
      @open="openUserDialog"
      width="400px"
    >
      <el-table ref="table" :data="taskData" style="width: 100%" @selection-change="handleSelectionChange">
        <el-table-column prop="join" type="selection" width="55"></el-table-column>
        <el-table-column label="任务名" prop="taskName"></el-table-column>
      </el-table>
      <span slot="footer" class="dialog-footer">
        <el-button @click="userSelect.dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="downloadMultiTask">下 载</el-button>
      </span>
    </el-dialog>

  </div>
</template>
<script>
import { getTaskListN as getList, getDataBlob } from "@/models/Task.js";
export default {
  name: "list",
  data() {
    return {
      taskData: [],
      account: this.global.account,
      loading: false,
      userSelect: {
        dialogVisible: false,
        tableData: [],
        tableSelect: [],
        type: "user",
        taskItem: [],
        loading: false
      },
      page: { now: 1, total: 1 }
    };
  },
  methods: {
    downloadMultiTask(){
      console.log(this.userSelect.tableSelect[1])
      console.log(this.userSelect.tableSelect.length)
      for(var i=0;i<this.userSelect.tableSelect.length;i++){
        this .downloadTask(this.userSelect.tableSelect[i])
      }

    },
    handleSelectionChange(val) {
      this.userSelect.tableSelect = val;
    },
    editUser(item) {
      [this.userSelect.dialogVisible, this.userSelect.type, this.userSelect.taskItem] = [true, "user", item];
    },
    showList() {
      if (this.account.group == 0) {
        return false;
      }
      this.loading = true;
      getList(this.account)
        .then(data => {
          this.taskData = data.data.taskList ? data.data.taskList : data.data;
        })
        .catch(data => {
          console.error(data);
          this.$message({ type: "error", message: "获取任务列表失败" });
        })
        .then(() => {
          this.loading = false;
        });
    },
    enterTask(item) {
      this.$router.push({ path: `/task/${item.taskId}/${item.taskName}/${item.taskType}` });
    },
    downloadTask(item) {
      if (this.account.group != 3) {
        this.$message({ type: "error", message: "权限不足" });
        return;
      }
      this.$message({ type: "info", message: "正在准备数据..." });
      getDataBlob(item.taskId)
        .then(d => {
          let date = new Date();
          let month = date.getMonth() + 1 + "";
          let day = date.getDay() + 1 + "";
          let hours = date.getHours() + "";
          let seconds = date.getSeconds() + "";
          let minutes = date.getMinutes() + "";
          month = month.length == 1 ? "0" + month : month;
          day = day.length == 1 ? "0" + day : day;
          hours = hours.length == 1 ? "0" + hours : hours;
          minutes = minutes.length == 1 ? "0" + minutes : minutes;
          seconds = seconds.length == 1 ? "0" + seconds : seconds;
          //prettier-ignore
          let filename = `${item.taskName}_${date.getFullYear()}-${month}-${day}_${hours}-${minutes}-${seconds}`;
          let blob = d.blob;
          let eleLink = document.createElement("a");
          eleLink.download = filename;
          eleLink.style.display = "none";
          eleLink.href = URL.createObjectURL(blob);
          document.body.appendChild(eleLink);
          eleLink.click();
          document.body.removeChild(eleLink);
        })
        .catch(data => {
          this.$message({ type: "error", message: `数据获取失败: ${data.message}(${data.code})` });
        });
    },
    getTypeName(type) {
      switch (type) {
        case 1:
          return "分类";
        case 2:
          return "检测";
        case 3:
          return "分割";
        case 4:
          return "关键点检测";
        case 5:
          return "视频描述";
        default:
          return "未知类型";
      }
    },
    filterTag(val, row) {
      return row.taskType == val;
    },
    pageChange(p) {
      this.showList();
    }
  },
  mounted() {
    this.showList();
  }
};
</script>

<style>
.table-expand {
  font-size: 0;
}
.table-expand label {
  width: 90px;
  color: #99a9bf;
}
.table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}
.el-form-item .person {
  margin-right: 5px;
}
</style>
