<template>
  <div>
    <el-button icon="el-icon-plus" @click="addUserBtn" type="primary">添加用户</el-button>
    <el-button icon="el-icon-refresh-left" @click="showList">刷新列表</el-button>
    <el-input v-model="search" class="search" prefix-icon="el-icon-search" placeholder="搜索用户"></el-input>
    <el-table border :data="userData" stripe style="width: 100%;margin-top:10px;" v-loading="loading">
      <el-table-column prop="username" min-width="100" label="用户名"></el-table-column>
      <el-table-column prop="password" min-width="100" label="密码"></el-table-column>
      <el-table-column :filters="[{ text: '用户', value: 'ROLE_USER' }, { text: '审核', value: 'ROLE_REVIEWER' },{text:'管理',value:'ROLE_ADMIN'}]" :filter-method="filterTag" filter-placement="bottom-end" prop="authorities" label="用户组别" width="180">
        <template slot-scope="scope">
          <el-tag :type="scope.row.authorities =='ROLE_USER' ? 'primary' : 'success'" disable-transitions>{{scope.row.authorities =='ROLE_USER'?'用户':scope.row.authorities =='ROLE_REVIEWER'?"审核":"管理"}}</el-tag>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="160">
        <template slot-scope="scope">
          <el-button v-if="scope.row.authorities !='ROLE_ADMIN'" @click="editUser(scope.row)" size="mini">编辑</el-button>
          <el-button v-if="scope.row.authorities !='ROLE_ADMIN'" @click="deleteUser(scope.row)" type="danger" size="mini">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog title="添加用户" :visible.sync="dialogVisible" width="30%">
      <el-form ref="form" :rules="rules" :hide-required-asterisk="true" :model="form" label-width="80px">
        <el-form-item label="用户名" prop="username">
          <el-input placeholder="用户名" v-model="form.username"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input placeholder="密码" v-model="form.password" show-password></el-input>
        </el-form-item>
        <el-form-item label="用户组" prop="group">
          <el-select v-model="form.group" placeholder="请选择用户组">
            <el-option label="用户" :value="'ROLE_USER'"></el-option>
            <el-option label="审核" :value="'ROLE_REVIEWER'"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="addUser">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import { getUserList, editUser, deleteUser } from "@/models/User.js";
export default {
  data() {
    return {
      dialogVisible: false,
      type: "add",
      editId: 0,
      userData: [],
      tempUserData: [],
      form: {
        username: "",
        password: "",
        group: ""
      },
      rules: {
        username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
        password: [{ required: true, message: "请输入密码", trigger: "blur" }],
        group: [{ required: true, message: "请选择用户组", trigger: "change" }]
      },
      loading: false,
      account: this.global.account,
      search: ""
    };
  },
  watch: {
    search() {
      let v = this.search.toLowerCase();
      if (v == "") {
        this.userData = this.tempUserData;
        return;
      }
      let temp = this.tempUserData.filter(item => item.username.toLowerCase().indexOf(v) != -1);
      this.userData = temp;
    }
  },
  methods: {
    editUser(item) {
      [this.form.username, this.form.password, this.form.group] = [item.username, item.password, item.authorities];
      [this.dialogVisible, this.type, this.editId] = [true, "edit", item.userId];
    },
    deleteUser(item) {
      deleteUser(item.userId, this.account)
        .then(() => {
          this.$message({ type: "success", message: `删除用户成功` });
          this.showList();
        })
        .catch(data => {
          this.$message({ type: "error", message: `删除用户失败,${data.message}(${data.code})` });
        });
    },
    filterTag(val, row) {
      return row.authorities == val;
    },
    addUser() {
      this.$refs.form.validate(valid => {
        if (!valid) {
          return false;
        }
        this.dialogVisible = false;
        this.loading = true;
        editUser(this.form.username, this.form.password, this.form.group, this.type == "add" ? undefined : this.editId, this.account)
          .then(data => {
            this.$message({ type: "success", message: `${this.type == "add" ? "添加" : "修改"}新用户成功` });
            [this.form.username, this.form.password, this.form.group] = ["", "", ""];
          })
          .catch(data => {
            this.$message({ type: "error", message: `${this.type == "add" ? "添加" : "修改"}用户失败,${data.message}(${data.code})` });
          })
          .then(() => {
            this.loading = false;
            this.showList();
          });
      });
    },
    addUserBtn() {
      [this.form.username, this.form.password, this.form.group] = ["", "", ""];
      [this.dialogVisible, this.type] = [true, "add"];
    },
    showList() {
      this.loading = true;
      getUserList(this.account)
        .then(data => {
          this.userData = data.data.userList;
          this.tempUserData = data.data.userList;
        })
        .catch(data => {
          this.$message({ type: "error", message: `获取用户列表失败,${data.message}(${data.code})` });
        })
        .then(() => {
          this.loading = false;
        });
    }
  },
  mounted() {
    this.showList();
  }
};
</script>
<style scoped>
.search {
  width: 200px;
  margin: 0 10px;
}
</style>
