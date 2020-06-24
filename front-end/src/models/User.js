import { ajax, ajaxPost } from "@/models/Service.js";
function getUserList() {
  return ajax("/api/admin/getUserList");
}
function editUser(name, pass, type, uid) {
  if (uid) {
    return ajaxPost("/api/admin/editUser", { userId: uid, username: name, password: pass, authorities: type });
  } else {
    return ajaxPost("/api/admin/addUser", { username: name, password: pass, authorities: type });
  }
}
function deleteUser(uid) {
  return ajaxPost("/api/admin/deleteUser", { userId: uid });
}
export { getUserList, editUser, deleteUser };
