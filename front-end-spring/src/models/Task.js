import { ajax, ajaxPost, getUserGroupStr, ajaxFile } from "@/models/Service.js";
function getTaskList(page) {
  let group = localStorage.getItem("group");
  let id = localStorage.getItem("id");
  let d = getUserGroupStr(group);
  let b = { page: page ? page : 1, limit: 30 };
  if (group == 1 || group == 2) {
    b[d + "Id"] = id;
    return ajaxPost(`/api/${d}/getTaskList`, b);
  }
  return ajaxPost(`/api/${d}/getTaskList`, b);
}
function getTaskListN() {
  let group = localStorage.getItem("group");
  let id = localStorage.getItem("id");
  let d = getUserGroupStr(group);
  if (group == 1 || group == 2) {
    let b = {};
    b[d + "Id"] = id;
    return ajaxPost(`/api/${d}/taskList`, b);
  }
  return ajax(`/api/${d}/taskList`);
}
function getImgList(tid, page) {
  let group = localStorage.getItem("group");
  let d = getUserGroupStr(group);
  let id = localStorage.getItem("id");
  let b = { taskId: tid, page: page, limit: 30 };
  b[d + "Id"] = id;
  return ajaxPost(`/api/${d}/getImgList`, b);
}
function getImg(iid, uid) {
  let group = localStorage.getItem("group");
  let d = getUserGroupStr(group);
  let b = { imageId: iid };
  uid != undefined && (b["userId"] = uid);
  return ajaxPost(`/api/${d}/getImg`, b);
}
function getPendingUserList(iid) {
  let group = localStorage.getItem("group");
  let d = getUserGroupStr(group);
  return ajaxPost(`/api/${d}/getPendingUserList`, { imageId: iid });
}
function setFinalVersion(iid, uid) {
  let group = localStorage.getItem("group");
  let d = getUserGroupStr(group);
  return ajaxPost(`/api/${d}/setFinalVersion`, { imageId: iid, userConfirmId: uid });
}

function deleteImageById(iid, uid) {
  return ajaxPost("/api/admin/deleteImageById", { imageId: iid });
}
function saveLabel(iid, data, uid) {
  let group = localStorage.getItem("group");
  let d = getUserGroupStr(group);
  let b = { imageId: iid, data: data };
  uid != undefined && (b["userId"] = uid);
  return ajaxPost(`/api/${d}/saveLabel`, b);
}
function getCount() {
  return ajax("/api/admin/getCount");
}
function getDataBlob(id) {
  let b = { taskId: id };

  return new Promise((resolve, reject) => {
    ajaxFile("/api/admin/downloadDatas", b)
      .then(data => {
        console.log(data);
        var blob = new Blob([JSON.stringify(data.d, null, 2)], { type: "application/json" });
        resolve({ blob: blob, filename: data.f });
      })
      .catch(data => {
        reject(data);
      });
  });
}
export { getTaskList, getTaskListN, getImgList, getImg, getPendingUserList, setFinalVersion, saveLabel, getCount, getDataBlob, deleteImageById};
