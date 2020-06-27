let videolabel = function(BFPlayer, configure, labelData, src, saveCB, nextCB, previousCB) {
  let list = function(player, labelData) {
    let data = [];
    let dom = document.querySelector(".label-list-ul");
    let show = function() {
      player.showMark(+this.getAttribute("data-start"), +this.getAttribute("data-end"));
    };
    let hide = function() {
      player.hideMark();
    };
    let add = function(id, start, end, content) {
      data.push({ labelId: configure.labelId, type: 0, startTime: start, endTime: end, sentence: content });
      let e = document.createElement("div");
      e.innerHTML = ` <li class="label-video-list-item">
            <div class="label-list-header"><span>${start}</span><span>${end}</span></div>
            <div class="label-list-content">
                ${content}
            </div>
            <div class="label-list-options">
                <span data-id="${id}" class="iconfont icon-shanchu label-list-delete  label-video-btn"></span>
                <span data-start="${start}" data-end="${end}" class="iconfont icon-kejian label-list-visible label-video-btn"></span>
            </div>
        </li>`;
      // console.log("append to list", e, id, start, end, content, e.children[0]);
      e.children[0].querySelector(".label-list-visible").addEventListener("mouseenter", show);
      e.children[0].querySelector(".label-list-visible").addEventListener("mouseleave", hide);
      e.children[0].querySelector(".label-list-delete").addEventListener("click", function() {
        for (let i = 0; i < data.length; i++) {
          let item = data[i];
          if (item.startTime == start && item.endTime == end && item.sentence == content) {
            data.splice(i, 1);
          }
        }
        this.parentElement.parentElement.remove();
      });
      dom.appendChild(e.children[0]);
    };
    this.add = add;
    this.save = function() {
      return data;
    };
    let ld = function(labelData) {
      console.log(labelData);
      data = [];
      dom.innerHTML = "";
      for (const item of labelData) {
        add(new Date().getTime(), item.startTime, item.endTime, item.sentence);
      }
    };
    ld(labelData);
    this.reload = function(data, src) {
      player.reload(src);
      ld(data);
    };
    return this;
  };
  let sign = function() {
    let dom = {
      panel: document.querySelector(".label-sign"),
      submit: document.querySelector(".label-sign-submit"),
      cancel: document.querySelector(".label-sign-cancel"),
      text: document.querySelector(".label-sign-text"),
      selector: document.querySelector(".label-sign-selector"),
      question: document.querySelector(".label-sign-question")
    };
    let _config;
    this.load = function(config, start, end) {
      start = +start.toFixed(2);
      end = +end.toFixed(2);
      if (config.type == 0) {
        dom.question.innerHTML = "请输入描述";
      } else if (config.type == 1) {
        dom.question.innerHTML = "请填空或清空输入描述";
      } else {
        dom.question.innerHTML = config.question.replace("$", "___");
      }
      if (config.type == 1) {
        dom.text.value = config.question.replace("$", "___");
      } else {
        dom.text.value = "";
      }
      if (config.type == 2) {
        dom.selector.innerHTML = "";
        for (const item of config.selector) {
          let o = document.createElement("option");
          o.value = item;
          o.innerHTML = item;
          dom.selector.appendChild(o);
        }
        dom.selector.style.display = "initial";
        dom.text.value = config.question.replace("$", config.selector[0]);
      } else {
        dom.selector.style.display = "none";
      }
      dom.panel.style.display = "flex";
      _config = config;
      return new Promise((resolve, reject) => {
        let ev = function() {
          //TODO:添加操作通信
          console.log(`Add mark "${dom.text.value}" at ${start} to ${end}`);
          l.add(new Date().getTime(), start, end, dom.text.value);
          dom.submit.removeEventListener("click", ev);
          dom.cancel.removeEventListener("click", dev);
          dom.panel.style.display = "none";
          resolve(dom.text.value);
        };
        let dev = function() {
          dom.submit.removeEventListener("click", ev);
          dom.cancel.removeEventListener("click", dev);
          dom.panel.style.display = "none";
          reject(dom.text.value);
        };
        dom.submit.addEventListener("click", ev);
        dom.cancel.addEventListener("click", dev);
      });
    };
    dom.selector.addEventListener("change", function() {
      dom.text.value = _config.question.replace("$", dom.selector.options[dom.selector.selectedIndex].value);
    });
    return this;
  };
  let s = new sign();
  let player = new BFPlayer(
    document.querySelector(".label-video"),
    src,
    function(start, end) {
      console.log(configure);
      return new Promise((resolve, reject) => {
        s.load(configure, start, end)
          .then(data => {
            resolve(data);
          })
          .catch(data => {
            reject(data);
          });
      });
    },
    saveCB,
    previousCB,
    nextCB
  );

  let l = new list(player, labelData);
  return l;
};
export default videolabel;
