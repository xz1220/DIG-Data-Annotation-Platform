import { ajax, ajaxPost } from "@/models/Service.js";
function deleteDataset(id) {
  return ajaxPost("/api/deleteDataset", `id=${id}`);
}
export { deleteDataset };
