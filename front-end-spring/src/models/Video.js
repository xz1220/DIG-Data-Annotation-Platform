import { ajax, ajaxPost, getUserGroupStr } from "@/models/Service.js";
function getVideo(iid, uid) {
  let group = localStorage.getItem("group");
  let d = getUserGroupStr(group);
  let b = { videoId: iid };
  uid != undefined && (b["userId"] = uid);
  return ajaxPost(`/api/${d}/getVideo`, b);
}
function getVideoList(tid, page) {
  let group = localStorage.getItem("group");
  let d = getUserGroupStr(group);
  let id = localStorage.getItem("id");
  let b = { taskId: tid, page: page, limit: 30 };
  b[d + "Id"] = id;
  return ajaxPost(`/api/${d}/getVideoList`, b);
}
function saveVideoLabel(vid, data, uid) {
  let group = localStorage.getItem("group");
  let d = getUserGroupStr(group);
  let id = localStorage.getItem("id");
  let b = { userId: uid ? uid : id, videoId: vid, data: data };
  return ajaxPost(`/api/${d}/saveVideoLabel `, b);
}
function getPendingUserList(id) {
  let group = localStorage.getItem("group");
  let d = getUserGroupStr(group);
  return ajaxPost(`/api/${d}/getVideoPendingUserList`, { videoId: id });
}
function setVideoFinalVersion(id, uid) {
  let group = localStorage.getItem("group");
  let d = getUserGroupStr(group);
  return ajaxPost(`/api/${d}/setVideoFinalVersion`, { videoId: id, userConfirmId: uid });
}
function getLabelList() {
  return ajax("/api/admin/getVideoLabelList");
}
function deleteLabel(id) {
  return ajaxPost("/api/admin/deleteVideoLabel", {
    labelId: id
  });
}
function addLabel(type, question, selector) {
  return ajaxPost("/api/admin/addVideoLabel", {
    type: type,
    question: question,
    selector: selector
  });
}
function editLabel(id, type, question, selector) {
  return ajaxPost("/api/admin/editVideoLabel", {
    labelId: id,
    type: type,
    question: question,
    selector: selector
  });
}
export default {
  getVideo,
  getVideoList,
  saveVideoLabel,
  getPendingUserList,
  setVideoFinalVersion,
  getLabelList,
  deleteLabel,
  addLabel,
  editLabel
};
