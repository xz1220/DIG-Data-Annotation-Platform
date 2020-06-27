import { ajax, ajaxPost } from "@/models/Service.js";
function getLabelList() {
  return ajax("/api/admin/getLabelList");
}
function deleteLabel(id) {
  return ajaxPost("/api/admin/deleteLabel", {
    labelId: id
  });
}
function editLabel(id, name, type, color) {
  return ajaxPost("/api/admin/editLabel", {
    labelId: id,
    labelName: name,
    labelType: type,
    labelColor: color
  });
}
function addLabel(name, type, color) {
  return ajaxPost("/api/admin/addLabel", {
    labelName: name,
    labelType: type,
    labelColor: color
  });
}
export { getLabelList, deleteLabel, editLabel, addLabel };
