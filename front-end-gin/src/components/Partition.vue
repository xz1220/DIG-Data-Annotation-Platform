<template>
  <div>
    <el-button icon="el-icon-plus" @click="newTaskAdd" type="primary">添加任务</el-button>
    <el-button icon="el-icon-refresh-left" @click="showList">刷新列表</el-button>
    <el-input
      v-model="search"
      @keydown.enter.native="searchClick"
      class="search"
      prefix-icon="el-icon-search"
      placeholder="搜索任务（回车确认）"
    ></el-input>

    <el-table border :data="taskData.list" stripe style="width: 100%;margin-top:10px;" v-loading="loading">
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
      <el-table-column prop="taskName" min-width="100" label="任务名"></el-table-column>
      <el-table-column fixed="right" label="操作" width="530">
        <template slot-scope="scope">
          <el-button v-if="scope.row.taskType != 5" @click="editType(scope.row)" size="mini">选择类型</el-button>
          <el-button @click="editUser(scope.row)" size="mini">选择用户</el-button>
          <el-button @click="editReviewer(scope.row)" size="mini">选择审核</el-button>
          <el-button @click="editLabel(scope.row)" size="mini">选择标签</el-button>
          <el-button @click="splitTask(scope.row)" size="mini">拆分</el-button>
          <el-button @click="deleteTask(scope.row)" type="danger" size="mini">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!--人员选择-->
    <el-dialog
      :title="userSelect.type == 'user' ? '选择用户' : '选择审核'"
      :visible.sync="userSelect.dialogVisible"
      @open="openUserDialog"
      width="400px"
    >
      <el-table ref="table" :data="userSelect.tableData" style="width: 100%" @selection-change="handleSelectionChange" v-loading="userSelect.loading">
        <el-table-column prop="join" type="selection" width="55"></el-table-column>
        <el-table-column label="用户名" prop="username"></el-table-column>
      </el-table>
      <span slot="footer" class="dialog-footer">
        <el-button @click="userSelect.dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="setUserTable">确 定</el-button>
      </span>
    </el-dialog>
    <!--标签选择-->
    <el-dialog title="选择标签" :visible.sync="labelSelect.dialogVisible" @open="openSelectDialog" width="400px">
      <el-table
        ref="labelTable"
        :data="labelSelect.tableData"
        style="width: 100%"
        @selection-change="labelSelectionChange"
        v-loading="labelSelect.loading"
      >
        <el-table-column prop="join" type="selection" width="55"></el-table-column>
        <template v-if="labelSelect.taskItem.taskType != 5">
          <el-table-column label="标签名称" prop="labelName">
            <template slot-scope="scope">
              <el-tag v-if="labelSelect.index[scope.row.labelName]" class="task-tag" effect="plain" size="mini">{{
                labelSelect.index[scope.row.labelName]
              }}</el-tag
              >{{ scope.row.labelName }}
            </template>
          </el-table-column>
        </template>
        <template v-else>
          <el-table-column label="类型" prop="type">
            <template slot-scope="scope">
              <span>{{ scope.row.type == 0 ? "空白" : scope.row.type == 1 ? "填空" : "选择" }}</span>
            </template>
          </el-table-column>
          <el-table-column label="描述" prop="question"></el-table-column>
        </template>
      </el-table>
      <span slot="footer" class="dialog-footer">
        <el-button @click="labelSelect.dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="setLabelTable">确 定</el-button>
      </span>
    </el-dialog>
    <!-- 拆分 -->
    <el-dialog title="任务拆分" :visible.sync="split.dialogVisible" width="30%" v-loading="loading">
      <el-form :inline="true">
        <el-form-item label="拆分数量">
          <el-input-number :min="2" :max="split.item.imageNumber" v-model="split.count" :step="1" step-strictly></el-input-number>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="split.dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="splitSubmit">确 定</el-button>
      </span>
    </el-dialog>
    <!-- 类型 -->
    <el-dialog title="任务类型" :visible.sync="typeSelect.dialogVisible" width="30%" v-loading="typeSelect.loading">
      <el-form :inline="true">
        <el-form-item label="类型" prop="group">
          <el-select v-model="typeSelect.select" placeholder="请选择任务类型">
            <el-option v-for="item in newTask.type" :key="item.value" :label="item.label" :value="item.value"> </el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="typeSelect.dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="typeSubmit">确 定</el-button>
      </span>
    </el-dialog>
    <!-- 添加新任务 -->
    <el-dialog title="导入新任务" :visible.sync="newTask.dialogVisible" width="60%" v-loading="newTask.loading">
      <el-table border :data="newTask.list" stripe style="width: 100%;margin-top:10px;" v-loading="loading">
        <el-table-column prop="taskName" min-width="100" label="任务名"></el-table-column>
        <el-table-column label="设置类型" width="120">
          <template slot-scope="scope">
            <el-select v-if="scope.row.taskType != 5" v-model="scope.row.taskType" placeholder="请选择">
              <el-option v-for="item in newTask.type" :key="item.value" :label="item.label" :value="item.value"> </el-option>
            </el-select>
            <span v-else>视频描述</span>
          </template>
        </el-table-column>
      </el-table>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="newTaskAddSubmit">全部添加</el-button>
      </span>
    </el-dialog>
    <el-pagination
      v-if="page.total > 1"
      @current-change="pageChange"
      class="pagination"
      background
      layout="prev, pager, next,jumper"
      :page-count="page.total"
      :current-page.sync="page.now"
    ></el-pagination>
  </div>
</template>
<script>
import Partition from "@/models/Partition";
import Video from "@/models/Video";
export default {
  name: "partition",
  data() {
    return {
      taskData: this.global.task,
      userSelect: {
        dialogVisible: false,
        tableData: [],
        tableSelect: [],
        type: "user",
        taskItem: [],
        loading: false
      },
      labelSelect: {
        dialogVisible: false,
        tableData: [],
        tableSelect: [],
        index: [], //每个tableData是否被选择的arr,对应index
        taskItem: [],
        loading: false
      },
      typeSelect: {
        dialogVisible: false,
        select: "",
        taskItem: null,
        loading: false
      },
      newTask: {
        list: [],
        dialogVisible: false,
        loading: false,
        type: [{ value: 1, label: "分类" }, { value: 2, label: "检测" }, { value: 3, label: "分割" }, { value: 4, label: "关键点" }]
      },
      account: this.global.account,
      loading: false,
      split: {
        dialogVisible: false,
        count: 2,
        item: ""
      },
      page: { now: 1, total: 1 },
      search: ""
    };
  },
  methods: {
    newTaskAdd() {
      this.newTask.dialogVisible = true;
      this.newTask.loading = true;
      Partition.getNewTaskList()
        .then(data => {
          this.newTask.list = data.data.taskList;
          this.newTask.loading = false;
        })
        .catch(data => {
          this.$message({ type: "error", message: `获取列表失败,${data.message}(${data.code})` });
        })
        .then(() => {
          this.newTask.loading = false;
        });
    },
    newTaskAddSubmit() {
      this.newTask.loading = true;
      Partition.updateTaskType(this.newTask.list)
        .then(() => {
          this.$message({ type: "success", message: `添加成功` });
          this.typeSelect.dialogVisible = false;
          this.showList();
        })
        .catch(data => {
          this.$message({ type: "error", message: `添加失败,${data.message}(${data.code})` });
        })
        .then(() => {
          this.newTask.loading = false;
        });
    },
    typeSubmit() {
      this.typeSelect.loading = true;
      Partition.updateTaskType([{ taskId: this.typeSelect.taskItem.taskId, taskType: this.typeSelect.select }])
        .then(() => {
          this.$message({ type: "success", message: `添加成功` });
          this.typeSelect.dialogVisible = false;
          this.showList();
        })
        .catch(data => {
          this.$message({ type: "error", message: `添加失败,${data.message}(${data.code})` });
        })
        .then(() => {
          this.typeSelect.loading = false;
        });
    },
    showList() {
      this.loading = true;

      Partition.getTaskList(this.page.now)
        .then(data => {
          this.page.total = data.data.totalpages ? data.data.totalpages : 1;
          this.global.task.list = data.data.taskList;
        })
        .catch(data => {
          this.$message({ type: "error", message: `获取列表失败,${data.message}(${data.code})` });
        })
        .then(() => {
          this.loading = false;
        });
    },
    searchClick() {
      this.page.now = 1;
      if (this.search == "") {
        this.showList();
        return;
      }
      this.loading = true;
      Partition.search(this.search)
        .then(data => {
          this.global.task.list = data.data;
        })
        .catch(data => {
          this.$message({ type: "error", message: `搜索失败,${data.message}(${data.code})` });
        })
        .then(() => {
          this.loading = false;
        });
    },
    editUser(item) {
      [this.userSelect.dialogVisible, this.userSelect.type, this.userSelect.taskItem] = [true, "user", item];
    },
    editReviewer(item) {
      [this.userSelect.dialogVisible, this.userSelect.type, this.userSelect.taskItem] = [true, "reviewer", item];
    },
    editLabel(item) {
      [this.labelSelect.dialogVisible, this.labelSelect.taskItem] = [true, item];
    },
    editType(item) {
      [this.typeSelect.dialogVisible, this.typeSelect.taskItem, this.typeSelect.select] = [true, item, item.taskType];
    },
    setUserTable() {
      this.loading = true;
      this.userSelect.dialogVisible = false;
      let ps =
        this.userSelect.type == "user"
          ? Partition.setUser(this.userSelect.taskItem, this.userSelect.tableSelect, this.account)
          : Partition.setReviewer(this.userSelect.taskItem, this.userSelect.tableSelect, this.account);
      ps.then(() => {
        this.$message({ type: "success", message: `保存成功` });
      })
        .catch(data => {
          this.$message({ type: "error", message: `保存失败,${data.message}(${data.code})` });
        })
        .then(() => {
          this.loading = false;
          this.showList();
        });
    },
    openUserDialog() {
      this.$nextTick(() => {
        this.userSelect.loading = true;
        let ps = Partition.getList(this.userSelect.taskItem, this.userSelect.type == "reviewer" ? "Reviewer" : "User");
        ps.then(data => {
          this.userSelect.tableData = data;
          this.$nextTick(() => {
            this.userSelect.tableSelect = [];
            for (const item of data) {
              if (item.join) {
                this.$refs.table.toggleRowSelection(item, true);
              } else {
                this.$refs.table.toggleRowSelection(item, false);
              }
            }
          });
        })
          .catch(data => {
            this.$message({ type: "error", message: `获取帐号列表失败,${data.message}(${data.code})` });
          })
          .then(() => {
            this.userSelect.loading = false;
          });
      });
    },
    handleSelectionChange(val) {
      this.userSelect.tableSelect = val;
    },
    setLabelTable() {
      this.loading = true;
      this.labelSelect.dialogVisible = false;
      console.log(this.labelSelect.tableSelect);
      Partition.setLabel(this.labelSelect.taskItem, this.labelSelect.tableSelect, this.account)
        .then(() => {
          this.$message({ type: "success", message: `保存成功` });
        })
        .catch(data => {
          this.$message({ type: "error", message: `保存失败,${data.message}(${data.code})` });
        })
        .then(() => {
          this.loading = false;
          this.showList();
        });
    },
    labelSelectionChange(val) {
      if (this.labelSelect.taskItem.taskType == 5) {
        if (val.length > 1) {
          this.$message({ type: "error", message: `只可选择一个模板` });
          this.$refs.labelTable.toggleRowSelection(this.labelSelect.tableSelect[0]);
          this.labelSelect.tableSelect = [this.labelSelect.tableSelect[this.labelSelect.tableSelect.length - 1]];
          return false;
        }
      }
      this.labelSelect.index = [];
      if (this.labelSelect.taskItem.taskType == 4) {
        let i = 1;
        for (const item of val) {
          // console.log(item.labelName);
          this.labelSelect.index[item.labelName] = i++;
        }
      }
      this.labelSelect.tableSelect = val;
      // console.log("select :", val);
      // console.log("index :",  this.labelSelect.index );
    },
    openSelectDialog() {
      this.$nextTick(() => {
        this.labelSelect.loading = true;
        let ps = Partition.getLabelList(this.labelSelect.taskItem.taskType, this.labelSelect.taskItem.labelIds);
        ps.then(data => {
          this.labelSelect.tableData = data;
          this.$nextTick(() => {
            this.labelSelect.tableSelect = [];
            for (const item of data) {
              if (item.join) {
                this.$refs.labelTable.toggleRowSelection(item, true);
              } else {
                this.$refs.labelTable.toggleRowSelection(item, false);
              }
            }
          });
        })
          .catch(data => {
            this.$message({ type: "error", message: `获取标签失败,${data.message}(${data.code})` });
          })
          .then(() => {
            this.labelSelect.loading = false;
          });
      });
    },
    splitTask(item) {
      if (item.imageNumber <= 1) {
        this.$alert("该任务只存在一项项目，无法继续分割", "出现错误", { confirmButtonText: "确定" });
        return;
      }
      this.split.item = item;
      this.split.dialogVisible = true;
    },
    splitSubmit() {
      this.loading = true;
      Partition.splitTask(this.split.item.taskId, this.split.count, this.account)
        .then(() => {
          this.$message({ type: "success", message: `拆分成功` });
        })
        .catch(data => {
          this.$message({ type: "error", message: `拆分失败,${data.message}(${data.code})` });
        })
        .then(() => {
          this.split.dialogVisible = false;
          this.showList();
        });
    },
    deleteTask(item) {
      this.$confirm("将永久删除数据, 是否继续?", "提示", { confirmButtonText: "确定", cancelButtonText: "取消", type: "warning" }).then(() => {
        this.loading = true;
        Partition.deleteTask(item.taskId, this.account)
          .then(() => {
            this.$message({ type: "success", message: `删除成功` });
          })
          .catch(data => {
            this.$message({ type: "error", message: `删除成功失败,${data.message}(${data.code})` });
          })
          .then(() => {
            this.loading = false;
            this.showList();
          });
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