<template>
  <p>登录中...</p>
</template>
<script>
import { checkIsLogin, getUserGroupStr } from "@/models/Service";
export default {
  name: "app",
  props: ["path"],
  data() {
    return {
      account: this.global.account
    };
  },
  mounted: function() {
    let path = "/" + decodeURI(this.path.replace(/,/g, "/"));
    let g = localStorage.getItem("group");
    if (!g) {
      this.$router.push("/login" + path);
    } else {
      checkIsLogin(`/api/${getUserGroupStr(g)}/getTaskList`)
        .then(() => {
          this.account.group = parseInt(localStorage.getItem("group"));
          this.account.username = localStorage.getItem("username");
          this.account.id = localStorage.getItem("id");
          this.account.token = localStorage.getItem("token");
          this.$router.push(path);
        })
        .catch(() => {
          this.$router.push("/login/" + this.path);
        });
    }
  }
};
</script>