<template>
  <div>
    <el-button icon="el-icon-plus" @click="addLabelBtn" type="primary">添加标签</el-button>
    <el-button icon="el-icon-refresh-left" @click="showList">刷新列表</el-button>
    <el-input v-model="search" class="search" prefix-icon="el-icon-search" placeholder="搜索标签"></el-input>
    <el-table border :data="labelData" stripe style="width: 100%;margin-top:10px;" v-loading="loading">
      <el-table-column prop="type" label="类型">
        <template slot-scope="scope">
          <span>{{ scope.row.type == 0 ? "空白" : scope.row.type == 1 ? "填空" : "选择" }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="question" min-width="100" label="预设描述"></el-table-column>
      <el-table-column prop="selector" label="选择项">
        <template slot-scope="scope">
          <span>{{ scope.row.selector == null ? "无" : parserSelector(scope.row.selector) }}</span>
        </template>
      </el-table-column> 
      <el-table-column fixed="right" label="操作" width="160">
        <template slot-scope="scope">
          <el-button v-if="scope.row.type != 0" @click="editLabel(scope.row)" size="mini">编辑</el-button>
          <el-button v-if="scope.row.type != 0" @click="deleteLabel(scope.row)" type="danger" size="mini">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog v-loading="loading" title="添加/编辑标签" :visible.sync="dialogVisible" width="30%">
      <el-form ref="form" :rules="rules" :hide-required-asterisk="true" :model="form" label-width="80px">
        <el-form-item label="类型" prop="type">
          <el-select v-model="form.type" placeholder="请选择类型">
            <el-option label="填空" :value="1"></el-option>
            <el-option label="选择" :value="2"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="预设描述" prop="question">
          <el-input placeholder="预设描述，用'$'代表空格项" v-model="form.question"></el-input>
        </el-form-item>
        <el-form-item v-if="form.type == 2" label="选择项" prop="selector">
          <el-input placeholder="输入选择项，用','分割" v-model="form.selector"></el-input>
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
import Video from "@/models/Video.js";
export default {
  name: "Labelset",
  data() {
    return {
      form: {
        question: "",
        type: 1,
        selector: ""
      },
      rules: {
        question: [{ required: true, message: "请选择预设描述", trigger: "change" }],
        type: [{ required: true, message: "请选择类型", trigger: "blur" }],
        selector: [{ required: true, message: "请输入选择项", trigger: "change" }]
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
      let temp = this.tempLabelData.filter(item => item.question != null && item.question.toLowerCase().indexOf(v) != -1);
      this.labelData = temp;
    }
  },
  methods: {
    showList() {
      this.search = "";
      this.loading = true;
      Video.getLabelList()
        .then(data => {
          this.labelData = data.data;
          this.tempLabelData = data.data;
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
      this.form = {
        question: "",
        type: 1,
        selector: ""
      };
    },
    editLabel(item) {
      this.type = "edit";
      this.editId = item.labelId;
      this.form.question = item.question;
      this.form.type = item.type;
      this.form.selector = this.parserSelector(item.selector);
      this.dialogVisible = true;
    },
    deleteLabel(item) {
      this.loading = true;
      Video.deleteLabel(item.labelId, this.account)
        .then(() => {
          this.$message({ type: "success", message: `删除标签"${item.question}"成功` });
        })
        .catch(data => {
          this.$message({ type: "error", message: `删除标签"${item.question}"失败,${data.message}(${data.code})` });
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
            ? Video.addLabel(this.form.type, this.form.question, this.parserSelector(this.form.selector, true))
            : Video.editLabel(this.editId, this.form.type, this.form.question, this.parserSelector(this.form.selector, true));
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
    },
    parserSelector(source, toArray) {
      if (toArray) {
        return source.split(",");
      }
      console.log(source);
      if (source == "") {
        return null;
      }
      let t = "";
      for (let i = 0; i < source.length; i++) {
        t += source[i] + (i != source.length - 1 ? "," : "");
      }
      return t;
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

.search {
  width: 200px;
  margin: 0 10px;
}
</style>
