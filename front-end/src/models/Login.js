import { ajax, ajaxPost } from "@/models/Service.js";
function logout(account) {
  return ajax("/api/logout",account);
}
function login(username, password) {
  return ajaxPost("/api/login", `username=${username}&password=${password}`);
  //  var login = {'username':username,'password':password};
  //  return ajaxPost("/api/login", login);
}
export {login, logout };
