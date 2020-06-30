import { ajax, ajaxPost } from "@/models/Service";
function getTaskList(page) {
  let b = { page: page ? page : 1, limit: 30 };
  return ajaxPost("/api/admin/getTaskList", b);
}
function getList(items, type) {
  let taskId = items.taskId;
  let select = {};
  console.log(items);
  for (const i of type == "User" ? items.userIds : items.reviewerIds) {
    select[i + ""] = true;
  }
  return new Promise((resolve, reject) => {
    ajaxPost("/api/admin/getList" + type, { taskId: taskId })
      .then(data => {
        let user = [];
        for (const item of data.data) {
          user.push({
            userId: item.userId,
            username: item.username,
            join: select[item.userId] == true
          });
        }
        resolve(user);
      })
      .catch(data => {
        reject(data);
      });
  });
}
function getLabelList(type, ids) {
  console.log("TaskType:" + type);
  let select = {};
  for (const item of ids) {
    select[item + ""] = true;
  }
  return new Promise((resolve, reject) => {
    let ps = type != 5 ? ajax("/api/admin/getLabelList") : ajax("/api/admin/getVideoLabelList");

    ps.then(data => {
      let labelList = type != 5 ? data.data.labelList : data.data;
      let list = [];
      for (const item of labelList) {
        type != 5
          ? list.push({
              labelId: item.labelId,
              labelName: item.labelName,
              join: select[item.labelId + ""] == true
            })
          : list.push({
              labelId: item.labelId,
              question: item.question,
              type: item.type,
              join: select[item.labelId + ""] == true
            });
      }
      resolve(list);
    }).catch(data => {
      reject(data);
    });
  });
}
function setReviewer(item, list) {
  let l = [];
  for (const i of list) {
    l.push(i.userId);
  }
  return setConfig(item.taskId, item.taskName, item.userIds, item.labelIds, l);
}
function setUser(item, list) {
  let l = [];
  for (const i of list) {
    l.push(i.userId);
  }
  return setConfig(item.taskId, item.taskName, l, item.labelIds, item.reviewerIds);
}
function setLabel(item, list) {
  let l = [];
  for (const i of list) {
    l.push(i.labelId);
  }
  return setConfig(item.taskId, item.taskName, item.userIds, l, item.reviewerIds);
}
function setConfig(taskId, taskName, userIds, labelIds, reviewerIds) {
  return ajaxPost("/api/admin/updateTask", {
    taskId: taskId,
    taskName: taskName,
    userIds: userIds,
    labelIds: labelIds,
    reviewerIds: reviewerIds,
    taskDesc: ""
  });
}
function splitTask(taskId, quantity) {
  return ajaxPost("/api/admin/splitTask", {
    taskId: taskId,
    quantity: quantity
  });
}
function deleteTask(taskId) {
  return ajaxPost("/api/admin/deleteTask", {
    taskId: taskId
  });
}
function updateTaskType(list) {
  return new Promise((resolve, reject) => {
    let i = 0;
    for (const item of list) {
      ajaxPost("/api/admin/updateTaskType", {
        taskId: item.taskId,
        taskType: item.taskType
      })
        .then(() => {
          if (++i == list.length) {
            resolve();
          }
        })
        .catch(data => {
          console.error(data);
          reject({ code: data.code, message: `处理taskId:'${item.taskId}'出现错误:${data.message}` });
        });
    }
  });
}
function getNewTaskList() {
  return ajax("/api/admin/getNewTaskList");
}
function search(keyword) {
  let b = { keyword };
  return ajaxPost(`/api/admin/searchTask`, b);
}
export default {
  getTaskList,
  getLabelList,
  getList,
  setReviewer,
  setUser,
  setLabel,
  splitTask,
  deleteTask,
  updateTaskType,
  getNewTaskList,
  search
};
