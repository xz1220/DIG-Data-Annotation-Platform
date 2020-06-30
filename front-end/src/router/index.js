import Vue from "vue";
import Router from "vue-router";
import Index from "@/views/Index";
import Login from "@/views/Login";
import Check from "@/views/Check";
import Dashboard from "@/components/Dashboard";
import Task from "@/components/Task";
import Label from "@/components/Label";
import VideoLabel from "@/components/VideoLabel";
import LabelsetVideo from "@/components/LabelsetVideo";
import User from "@/components/User";
import Labelset from "@/components/Labelset";
import Partition from "@/components/Partition";
import List from "@/components/List";
import { NAME } from "@/models/Service";
//import Dataset from "@/components/Dataset";
Vue.use(Router);
export default new Router({
  routes: [
    {
      path: "/",
      redirect: "/dashboard"
    },
    {
      path: "/",
      name: "Index",
      component: Index,
      props: true,
      children: [
        {
          path: "dashboard",
          name: "Dashboard",
          component: Dashboard,
          props: true,
          meta: {
            title: "仪表盘 - " + NAME
          }
        },
        {
          path: "task/:id/:taskName/:type",
          name: "Task",
          component: Task,
          props: true,
          meta: {
            title: "任务详情 - " + NAME
          }
        },
        {
          path: "user",
          name: "User",
          component: User,
          props: true,
          meta: {
            title: "账户列表 - " + NAME
          }
        },
        {
          path: "partition",
          name: "Partition",
          component: Partition,
          props: true,
          meta: {
            title: "任务操作 - " + NAME
          }
        },
        {
          path: "list",
          name: "List",
          component: List,
          props: true,
          meta: {
            title: "任务列表 - " + NAME
          }
        },
        /*{
          path: "dataset",
          name: "Dataset",
          component: Dataset,
          props: true
        },*/
        {
          path: "label",
          name: "Labelset",
          component: Labelset,
          props: true,
          meta: {
            title: "标签列表 - " + NAME
          }
        },
        {
          path: "videolabel",
          name: "LabelsetVideo",
          component: LabelsetVideo,
          props: true,
          meta: {
            title: "视频标签列表 - " + NAME
          }
        }
      ]
    },
    {
      path: "/img/:id/:tid/:type/:taskName/:imagePage",
      name: "Label",
      component: Label,
      props: true,
      meta: {
        title: "标记 - " + NAME
      }
    },
    {
      path: "/video/:id/:tid/:taskName/:imagePage",
      name: "VideoLabel",
      component: VideoLabel,
      props: true,
      meta: {
        title: "视频标记 - " + NAME
      }
    },
    {
      path: "/login/:path",
      name: "Login",
      component: Login,
      props: true,
      meta: {
        title: "登录 - " + NAME
      }
    },
    {
      path: "/check/:path",
      name: "Check",
      component: Check,
      props: true,
      meta: {
        title: "Loading..."
      }
    }
  ]
});
