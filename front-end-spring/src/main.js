// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue";
import router from "./router";
import ElementUI from "element-ui";
import App from "@/App.vue";
import global from "@/common.js";
import "element-ui/lib/theme-chalk/index.css";
import "@/assets/css/global.css";

Vue.config.productionTip = false;
Vue.use(ElementUI);
/* eslint-disable no-new */
Vue.prototype.global = global;
router.beforeEach((to, from, next) => {
  /* 路由发生变化修改页面title */
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()
})
console.log(global);
new Vue({
  el: "#app",
  router,
  components: { App },
  template: "<App/>"
});
