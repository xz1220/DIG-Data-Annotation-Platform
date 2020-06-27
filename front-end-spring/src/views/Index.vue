<template>
  <el-container style="height:100%">
    <el-header>
      <Header></Header>
    </el-header>
    <el-container style="height: calc(100% - 60px);">
      <el-aside style="padding-bottom: 5px;height: 100%;" width="200px">
        <Aside></Aside>
      </el-aside>
      <el-main>
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script >
import Header from "@/components/Header.vue";
import Aside from "@/components/Aside.vue";
export default {
  name: "App",
  components: {
    Header,
    Aside
  },
  mounted: function() {
    
  }
};
</script>
<style scoped>
header {
  border-bottom: 1px #eee solid;
}
</style> 
