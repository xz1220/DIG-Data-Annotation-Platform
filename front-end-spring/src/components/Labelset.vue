<template>
  <div>
    <el-button icon="el-icon-plus" @click="addLabelBtn" type="primary">添加标签</el-button>
    <el-button icon="el-icon-refresh-left" @click="showList">刷新列表</el-button>
    <el-input v-model="search" class="search" prefix-icon="el-icon-search" placeholder="搜索标签"></el-input>
    <el-table border :data="labelData" stripe style="width: 100%;margin-top:10px;" v-loading="loading">
      <el-table-column prop="labelName" min-width="100" label="名称"></el-table-column>
      <el-table-column prop="labelColor" label="颜色">
        <template slot-scope="scope">
          <div class="label-color" :style="'background-color:'+scope.row.labelColor"></div>
          <span>{{scope.row.labelColor}}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="160">
        <template slot-scope="scope">
          <el-button @click="editLabel(scope.row)" size="mini">编辑</el-button>
          <el-button @click="deleteLabel(scope.row)" type="danger" size="mini">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog v-loading="loading" title="添加/编辑标签" :visible.sync="dialogVisible" width="30%">
      <el-form ref="form" :rules="rules" :hide-required-asterisk="true" :model="form" label-width="80px">
        <el-form-item label="标签名称" prop="label">
          <el-input placeholder="标签名称" v-model="form.label"></el-input>
        </el-form-item>
        <el-form-item label="颜色" prop="color">
          <el-color-picker v-model="form.color" color-format="hex" :predefine="defaultColor"></el-color-picker>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="addLabel">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import { getLabelList, deleteLabel, editLabel, addLabel } from "@/models/Label.js";
export default {
  name: "Labelset",
  data() {
    return {
      form: {
        label: "",
        color: "#3498db"
      },
      defaultColor: ["#1abc9c", "#2ecc71", "#3498db", "#9b59b6", "#34495e", "#f1c40f", "#e67e22", "#e74c3c", "#95a5a6"],
      rules: {
        label: [{ required: true, message: "请输入名称", trigger: "blur" }],
        type: [{ required: true, message: "请选择标签类型", trigger: "change" }],
        color: [{ required: true, message: "请选择一个颜色", trigger: "change" }]
      },
      labelData: [],
      dialogVisible: false,
      editId: 0,
      type: "add",
      loading: false,
      account: this.global.account,
      search: "",
      tempLabelData: []
    };
  },
  watch: {
    search() {
      let v = this.search.toLowerCase();
      if (v == "") {
        this.labelData = this.tempLabelData;
        return;
      }
      let temp = this.tempLabelData.filter(item => item.labelName.toLowerCase().indexOf(v) != -1);
      this.labelData = temp;
    }
  },
  methods: {
    showList() {
      this.search = "";
      this.loading = true;
      getLabelList(this.account)
        .then(data => {
          this.labelData = data.data.labelList;
          this.tempLabelData = data.data.labelList;
        })
        .catch(data => {
          this.$message({ type: "error", message: `获取标签列表失败,${data.message}(${data.code})` });
        })
        .then(() => {
          this.loading = false;
        });
    },
    addLabelBtn() {
      this.type = "add";
      this.dialogVisible = true;
      this.form.label = "";
      this.form.color = "#3498db";
    },
    editLabel(item) {
      this.type = "edit";
      this.editId = item.labelId;
      this.form.label = item.labelName;
      this.form.color = item.labelColor;
      this.dialogVisible = true;
    },
    deleteLabel(item) {
      this.loading = true;
      deleteLabel(item.labelId, this.account)
        .then(() => {
          this.$message({ type: "success", message: `删除标签"${item.labelName}"成功` });
        })
        .catch(data => {
          this.$message({ type: "error", message: `删除标签"${item.labelName}"失败,${data.message}(${data.code})` });
        })
        .then(() => {
          this.loading = false;
        });
      this.showList();
    },
    addLabel() {
      this.$refs.form.validate(valid => {
        if (!valid) {
          return false;
        }
        this.loading = true;
        let ps =
          this.type == "add"
            ? addLabel(this.form.label, 0, this.form.color, this.account)
            : editLabel(this.editId, this.form.label, 0, this.form.color, this.account);
        ps.then(() => {
          this.$message({ type: "success", message: `${this.type == "add" ? "添加" : "修改"}标签成功` });
        })
          .catch(data => {
            this.$message({ type: "error", message: `${this.type == "add" ? "添加" : "修改"}标签失败,${data.message}(${data.code})` });
          })
          .then(() => {
            this.dialogVisible = false;
            this.loading = false;
            this.showList();
          });
      });
    }
  },
  mounted() {
    this.showList();
  }
};
</script>
<style scoped>
.label-color {
  width: 10px;
  height: 10px;
  display: inline-block;
  border-radius: 10px;
}

</style>
