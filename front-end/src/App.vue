<template>
  <transition>
    <router-view />
  </transition>
</template>
<script>
export default {
  name: "app",
  data() {
    return {
      account: this.global.account
    };
  },
  mounted: function() {
    if (this.$route.fullPath.indexOf("/login") == -1) {
      this.$router.push("/check/" + this.$route.fullPath.replace("/", "").replace(/\//g, ","));
    }
  }
};
</script>
<style scoped>
</style>
