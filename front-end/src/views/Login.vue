<template>
  <el-container>
    <el-main>
      <el-card class="box-card login-card" v-loading="loading">
        <h1>登录</h1>
        <el-form ref="form" :rules="rules" :hide-required-asterisk="true" :model="form" label-width="80px">
          <el-form-item label="用户名" prop="username">
            <el-input placeholder="用户名" v-model="form.username"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input placeholder="密码" v-model="form.password" show-password></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="login">登录</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-main>
  </el-container>
</template>
<script>
import { login } from "@/models/Login.js";
export default {
  name: "Login",
  props: ["path"],
  data: function() {
    return {
      form: {
        username: "",
        password: ""
      },
      rules: {
        username: [{ required: true, message: "请输入用户名", trigger: "blur" }],
        password: [{ required: true, message: "请输入密码", trigger: "blur" }]
      },
      account: this.global.account,
      loading: false
    };
  },
  methods: {
    login: function() {
    let path = "/" + decodeURI(this.path.replace(/,/g, "/"));
      this.$refs.form.validate(valid => {
        if (!valid) {
          return false;
        }
        this.loading = true;
        login(this.form.username, this.form.password)
          .then(data => {
            this.account.group = data.data.user.authorities == "ROLE_ADMIN" ? 3 : data.data.user.authorities == "ROLE_REVIEWER" ? 2 : 1;
            this.account.username = this.form.username;
            this.account.id = data.data.user.userId;
            this.account.token = data.data.token;

            localStorage.setItem("group", this.account.group);
            console.log("group:",this.account.group)
            localStorage.setItem("username", this.form.username);
            localStorage.setItem("id", data.data.user.userId);
            localStorage.setItem("token", data.data.token);

            this.$message({ type: "success", message: "登录成功" });
            this.$router.push(path);
          })
          .catch(data => {
            this.$alert(`${data.message}(${data.code})`, "登录失败", {
              confirmButtonText: "确定",
              type: "warning"
            });
          })
          .then(() => {
            this.loading = false;
          });
      });
    }
  },
  mounted: function() {}
};
</script>
<style scoped>
.login-card {
  width: 400px;
  padding: 0 40px 0 0;
  position: fixed;
  margin: auto;
  left: 0;
  top: 0;
  right: 0;
  bottom: 0;
  height: 300px;
}
.login-card h1 {
  padding: 0 28px;
}
</style>
<style>
</style>
