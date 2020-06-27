<template>
  <el-row>
    <el-col :span="24">
      <div class="header-wrap">
        <h1>
          {{name}}
          <small></small>
        </h1>
        <div class="header-spacer"></div>
        <div class="header-right">
          <el-dropdown @command="logout">
            <span class="el-dropdown-link">
              {{account.username}}
              <i class="el-icon-arrow-down el-icon--right"></i>
            </span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item>Log out</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
      </div>
    </el-col>
  </el-row>
</template>
<script>
import { NAME } from "@/models/Service";
import { logout as lgot } from "@/models/Login";
export default {
  name: "Header",
  data() {
    return {
      account: this.global.account,
      name: NAME
    };
  },
  methods: {
    logout() {
      document.cookie = "";
      lgot(this.account)
        .then(() => {
          this.$router.push("/login/dashboard");
        })
        .catch(() => {
          this.$message({
            type: "error",
            message: "退出登录失败，请重试"
          });
        });
    }
  },
  mounted() {}
};
</script>
<style scoped>
.header-wrap {
  display: flex;
}
.header-spacer {
  flex-grow: 1;
}
.el-dropdown-link {
  cursor: pointer;
  color: #409eff;
}
.el-icon-arrow-down {
  font-size: 12px;
}
h1 {
  margin: 0;
  margin-top: 14px;
  color: #303133;
}
.header-right {
  margin-top: 20px;
}
</style>
