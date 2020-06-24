/*
 * BFPlayer
 * Author: q6q64399 (q6q64399@gmail.com)
 */
let BFPlayer = function(area, link, signCB, saveCB, previousCB, nextCB) {
  let dom = {
    frame: area.querySelector(".bfplayer"),
    main: area.querySelector(".bf-main"),
    video: area.querySelector(".bf-video"),
    controller: area.querySelector(".bf-controller"),
    menu: area.querySelector(".bf-menu"),
    loading: area.querySelector(".bf-loading-panel"),
    capture: {
      panel: area.querySelector(".bf-capture"),
      start: area.querySelector(".bf-btn-start"),
      end: area.querySelector(".bf-btn-end")
    },
    pgb: {
      panel: area.querySelector(".bf-pgb-player"),
      buffer: area.querySelector(".bf-pgb-buffer"),
      played: area.querySelector(".bf-pgb-played"),
      select: area.querySelector(".bf-pgb-select"),
      point: area.querySelector(".bf-pgb-point"),
      img: area.querySelector(".bf-pgb-point>div")
    },
    btn: {
      play: area.querySelector(".bf-btn-play"),
      fullbrowser: area.querySelector(".bf-btn-fullbrowser"),
      fullscreen: area.querySelector(".bf-btn-fullscreen"),
      volume: area.querySelector(".bf-btn-volume"),
      save: area.querySelector(".bf-btn-save"),
      previous: area.querySelector(".bf-btn-previous"),
      next: area.querySelector(".bf-btn-next")
    },
    panel: {
      volume: {
        area: area.querySelector(".bf-volume-area"),
        panel: area.querySelector(".bf-volume"),
        value: area.querySelector(".bf-volume-value"),
        point: area.querySelector(".bf-volume-point")
      }
    },
    time: area.querySelector(".bf-label-time"),
    selectTime: area.querySelector(".bf-label-select-time")
  };
  let _config = {
    time: "0:00",
    duration: 0,
    controll: {
      lock: true,
      time: 0,
      pf: undefined
    },
    pgbMousedown: false, //按下时的lock
    pgbLock: false, //防止进度条反复横跳
    pgbLeft: 0, //进度条的scrollLeft,
    fullscreen: false,
    fullbrowser: false,
    select: {
      enable: false,
      start: 0,
      persent: 0
    }
  };
  let event = {
    keydown: function(e) {
      if (e.target.tagName != "INPUT") {
        switch (e.key) {
          case " ":
            dom.video.paused ? dom.video.play() : dom.video.pause();
            break;
          case "ArrowLeft":
            dom.video.currentTime > 2 ? (dom.video.currentTime -= 2) : (dom.video.currentTime = 0);
            //danmaku.jump(Math.floor(dom.video.currentTime * 1000));
            controller.reload();
            break;
          case "ArrowRight":
            dom.video.duration - dom.video.currentTime > 5 ? (dom.video.currentTime += 5) : (dom.video.currentTime = dom.video.duration);
            //danmaku.jump(Math.floor(dom.video.currentTime * 1000));
            controller.reload();
            break;
          case "ArrowUp":
            if (_config.snackbarLock) break;
            dom.video.volume >= 1 ? snackbar("音量已经达到最大值") : (dom.video.volume = +(dom.video.volume + 0.1 >= 1 ? 1 : dom.video.volume + 0.1).toFixed(1));
            controller.reload();
            break;
          case "ArrowDown":
            if (_config.snackbarLock) break;
            dom.video.volume <= 0 ? snackbar("当前是静音状态") : (dom.video.volume = +(dom.video.volume - 0.1 <= 0 ? 0 : dom.video.volume - 0.1).toFixed(1));
            controller.reload();
            break;
          case "A":
          case "a":
            controller.reload();
            event.capture.start();
            break;
          case "S":
          case "s":
            controller.reload();
            event.capture.end();
            break;
          case "W":
          case "w":
            screen.clientToggle();
            break;
          case "F":
          case "f":
            screen.screenToggle();
            break;
          default:
            break;
        }
      }
    },
    mouseleave: function() {
      clearTimeout(_config.controll.pf);
      _config.controll.time = 3;
      controller.timer();
    },
    mousemove: function(e) {
      controller.reload();
    },
    mousemoveall: function(e) {
      if (_config.pgbMousedown) {
        pgbMove(e);
      }
    },
    mousedown: function(e) {
      if (e.button == 2) {
        if (e.target.classList.contains("bf-nomenu")) {
          return;
        }
        _config.menuShow = true;
        let y = e.pageY - this.scrollTop - 54,
          x = e.pageX - this.scrollLeft;
        if (this.scrollHeight - y < 80) y = this.scrollHeight - 80;
        if (this.scrollWidth - x < 230) x = this.scrollWidth - 230;
        dom.menu.style.left = x + "px";
        dom.menu.style.top = y + "px";
        dom.menu.style.display = "inherit";
      } else if (e.button == 0) {
        if (e.target.tagName == "LI") {
          return;
        } else if (dom.menu.style.display != "none") {
          dom.menu.style.display = "none";
        } else if (e.target == dom.video || e.target == dom.controller) {
          dom.video.paused ? dom.video.play() : dom.video.pause();
        }
      }
    },
    mouseup: function(e) {
      if (_config.pgbMousedown) {
        dom.pgb.panel.classList.remove("bf-pgb-player-notransition");
        if (e.target != dom.pgb.img) (dom.pgb.point.style.display = "none"), dom.pgb.panel.classList.remove("bf-pgb-player-enter");
        _config.pgbMousedown = false;
        _config.pgbLock = true;
        //danmaku.jump(Math.floor(dom.pgb.time * 1000));
        dom.video.currentTime = dom.pgb.time;
      }
    },
    pgb: {
      mousedown: function(e) {
        if (!_config.pgbMousedown && e.button == 0) {
          _config.controll.lock = true;
          _config.pgbMousedown = true;
          this.classList.add("bf-pgb-player-notransition");
          _config.pgbLeft = dom.frame.scrollLeft + 20;
          pgbMove(e);
        }
      },
      img: function(e) {
        e.preventDefault();
      },
      mouseenter: function(e) {
        this.classList.add("bf-pgb-player-enter");
        dom.pgb.point.style.display = "inherit";
      },
      mouseleave: function(e) {
        if (_config.pgbMousedown) return;
        this.classList.remove("bf-pgb-player-enter");
        dom.pgb.point.style.display = "none";
      }
    },
    video: {
      timeupdate: function(e) {
        if (_config.pgbLock) {
          _config.pgbLock = false;
          !dom.video.paused && ((_config.controll.lock = false), controller.reload());
        } else if (!_config.pgbMousedown) {
          let percent = this.currentTime / this.duration;
          dom.pgb.played.style.width = percent * 100 + "%";
          dom.time.innerHTML = utils.format(this.currentTime) + " / " + _config.time;
          dom.pgb.point.style.left = (dom.frame.scrollWidth - 40) * percent + "px";
          if (_config.select.enable) {
            _config.controll.lock = true;
            let s = dom.pgb.select.style;
            s.left = (dom.frame.scrollWidth - 40) * _config.select.persent + "px";
            percent = (this.currentTime - _config.select.start) / this.duration;
            if (percent < 0) {
              _config.select.enable = false;
              s.display = "none";
              dom.selectTime.style.display = "none";
              return;
            }
            dom.selectTime.innerHTML = `已选择：${_config.select.start} - ${this.currentTime}`;
            s.width = percent * 100 + "%";
          }
        }
      },
      play() {
        _config.qrcode && event.btn.qrcode();
        let x = dom.btn.play.getElementsByTagName("span")[0].classList;
        x.remove("bf-icon-play");
        x.add("bf-icon-pause");
        _config.controll.lock = false;
        controller.reload();
        //danmaku.start();
      },
      progress() {
        dom.pgb.buffer.style.width = (utils.getBufferTime(this.buffered) / this.duration) * 100 + "%";
      },
      error() {},
      canplay() {
        dom.loading.style.display = "none";
        _config.duration = this.duration;
      },
      waiting() {
        dom.loading.style.display = "block";
      },
      ended() {},
      pause() {
        let x = dom.btn.play.getElementsByTagName("span")[0].classList;
        x.add("bf-icon-play");
        x.remove("bf-icon-pause");
        _config.controll.lock = true;
        controller.show();
        //danmaku.pause();
      },
      volumechange() {
        event.volume.change(this.volume);
      },
      loadstart() {
        event.video.pause();
      },
      durationchange() {
        _config.time = utils.format(this.duration);
        dom.time.innerHTML = "0:00 / " + _config.time;
      }
    },
    btn: {
      play() {
        dom.video.paused ? dom.video.play() : dom.video.pause();
      },
      fullscreen() {
        _config.fullscreen ? screen.exitScreen() : _config.fullbrowser ? screen.exitClient() : screen.screen();
      },
      fullbrowser() {
        _config.fullbrowser ? screen.exitClient() : _config.fullscreen ? screen.exitScreen() : screen.client();
      },
      volume() {
        if (this.lastValue) {
          dom.video.volume = this.lastValue;
          this.lastValue = false;
        } else {
          this.lastValue = dom.video.volume;
          dom.video.volume = 0;
        }
      },
      next() {
        nextCB && nextCB();
      },
      previous() {
        previousCB && previousCB();
      },
      save() {
        saveCB && saveCB();
      }
    },
    capture: {
      start() {
        _config.controll.lock = true;
        _config.select.start = dom.video.currentTime;
        let percent = dom.video.currentTime / dom.video.duration;
        _config.select.persent = percent;
        let s = dom.pgb.select.style;
        s.left = (dom.frame.scrollWidth - 40) * percent + "px";
        s.display = "initial";
        s.width = "0%";
        dom.selectTime.style.display = "initial";
        _config.select.enable = true;
      },
      end() {
        if (!_config.select.enable) {
          snackbar("请先设置初始点");
          return;
        }
        _config.controll.lock = false;
        _config.select.enable = false;
        dom.video.pause();
        signCB(_config.select.start, dom.video.currentTime)
          .then(data => {
            snackbar("添加标记成功");
          })
          .catch(data => {
            snackbar("添加标记失败");
          })
          .then(() => {
            dom.pgb.select.style.display = "none";
            dom.selectTime.style.display = "none";
          });
      }
    },
    tooltip: {
      show: function(e) {
        this.bottom = false;
        this.moveLock = false;
        if (this.classList.contains("bf-selector") && this.querySelector(".bf-selector-panel") != null) {
          this.bottom = true;
        }
        let tip = this.getAttribute("data-tooltip");
        let l = this.querySelector(".bf-tip");
        if (l != null) {
          l.innerHTML = tip;
          return;
        }
        l = document.createElement("div");
        l.innerHTML = tip;
        l.classList.add("bf-tip");
        this.appendChild(l);
        if (this.classList.contains("bf-tooltip-left")) {
          l.style.top = this.offsetHeight / 2 - l.offsetHeight / 2 + "px";
          l.style.left = -l.offsetWidth - 5 + "px";
        } else {
          l.style.top = this.bottom ? 5 : -l.offsetHeight - 5 + "px";
          l.style.left = -l.offsetWidth / 2 + this.offsetWidth / 2 + "px";
        }
        this.tooltip = l;
      },
      hide: function() {
        this.moveLock = false;
        let l = this.querySelector(".bf-tip");
        if (l != null) {
          l.remove();
        }
      },
      move: function() {
        if (!this.moveLock && this.querySelector(".bf-selector-panel") != null) {
          this.moveLock = true;
          this.bottom = true;
        }
        if (this.bottom) {
          this.tooltip.style.top = "15px";
        }
      }
    },
    selector: {
      show: function(e) {
        let list, x;
        let data = this.getAttribute("data-selector"),
          type = this.getAttribute("data-selector-type"),
          l = "",
          temp = document.createElement("div"),
          p = this,
          ev = this.getAttribute("data-selector-event");
        let removeEvent = function() {
          p.removeEventListener("mouseleave", removeEvent);
          list.lock = false;
          setTimeout(function() {
            if (!list.lock) {
              list.remove();
              p.addEventListener("mouseenter", event.selector.show);
            }
          }, 150);
        };
        typeof data != "object" && (data = JSON.parse(data));
        for (let i = 0; i < data.length; i++) {
          l += `<li data-index="${i}" data-value="${data[i].value}" class="bf-selector-item ${data[i].active ? "active" : ""}">${data[i].name}</li>`;
          data[i].active = false;
        }
        list = `<div class="bf-selector-panel"><ul>${l}</ul></div>`;
        temp.innerHTML = list;
        list = temp.childNodes[0];
        this.appendChild(list);
        if (type && type == "fixed") {
          list.style.position = "fixed";
          list.style.top = e.clientY + (p.offsetHeight - e.layerY) + "px";
          list.style.left = e.clientX + (p.offsetWidth / 2 - e.layerX) - list.offsetWidth / 2 + "px";
        } else {
          list.style.top = -list.offsetHeight - 10 + "px";
          list.style.left = -list.offsetWidth / 2 + this.offsetWidth / 2 + "px";
        }
        x = this.querySelectorAll(".bf-selector-item");
        for (let i = 0; i < x.length; i++) {
          x[i].addEventListener(
            "mousedown",
            function() {
              data[+this.getAttribute("data-index")].active = true;
              p.setAttribute("data-selector", JSON.stringify(data));
              p.setAttribute("data-value", this.getAttribute("data-value"));
              event[ev](this.getAttribute("data-value"));
              list.remove();
              p.addEventListener("mouseenter", event.selector.show);
            },
            true
          );
        }

        this.addEventListener("mouseleave", removeEvent);
        list.addEventListener("mouseenter", function() {
          list.lock = true;
          p.removeEventListener("mouseleave", removeEvent);
          list.addEventListener("mouseleave", function() {
            list.remove();
            p.addEventListener("mouseenter", event.selector.show);
          });
        });
        this.removeEventListener("mouseenter", event.selector.show);
      }
    },
    volume: {
      lock: false,
      moveLock: false,
      pr: 0,
      mouseenter: function() {
        dom.panel.volume.area.style.display = "initial";
        event.volume.lock = true;
      },

      mouseleave: function() {
        event.volume.lock = false;
        event.volume.pr = setTimeout(function() {
          if (!event.volume.lock) {
            event.volume.hide();
          }
        }, 180);
      },
      mouseenterTarget: function() {
        event.volume.lock = true;
      },
      mouseleaveTarget: function() {
        event.volume.lock = false;
        !event.volume.moveLock && event.volume.hide();
      },
      hide: function() {
        dom.panel.volume.area.classList.add("bf-volume-hidden");
        event.volume.pr = setTimeout(function() {
          if (!event.volume.lock && !event.volume.moveLock) {
            dom.panel.volume.area.style.display = "none";
            dom.panel.volume.area.classList.remove("bf-volume-hidden");
          }
        }, 150);
      },
      mousedownTarget: function(e) {
        let pos;
        let left = this.offsetLeft + this.parentElement.parentElement.offsetLeft + this.parentElement.parentElement.parentElement.parentElement.offsetLeft;
        let mousemove = function(e) {
            pos = (e.pageX - left) / 80;
            dom.video.volume = pos < 0 ? 0 : pos > 1 ? 1 : pos;
          },
          mouseup = function() {
            _config.controll.lock = dom.video.paused;
            event.volume.moveLock = false;
            document.removeEventListener("mousemove", mousemove, true);
            document.removeEventListener("mouseup", mouseup, true);
            dom.panel.volume.panel.classList.remove("bf-volume-noanimation");
          };
        _config.controll.lock = true;
        event.volume.moveLock = true;
        pos = (e.pageX - left) / 80;
        dom.video.volume = pos < 0 ? 0 : pos > 1 ? 1 : pos;
        dom.panel.volume.panel.classList.add("bf-volume-noanimation");
        document.addEventListener("mousemove", mousemove, true);
        document.addEventListener("mouseup", mouseup, true);
      },
      change: function(v) {
        let point = dom.panel.volume.point,
          panel = dom.panel.volume.area,
          value = dom.panel.volume.value,
          list = dom.btn.volume.children[0].classList;
        panel.style.display = "initial";
        clearTimeout(event.volume.pr);
        list.remove("bf-icon-volumedown"), list.remove("bf-icon-volumemute"), list.remove("bf-icon-volumeoff"), list.remove("bf-icon-volumeup");
        if (v > 0.7) {
          list.add("bf-icon-volumeup");
        } else if (v < 0.3 && v > 0) {
          list.add("bf-icon-volumemute");
        } else if (v == 0) {
          list.add("bf-icon-volumeoff");
        } else {
          list.add("bf-icon-volumedown");
        }
        point.style.left = v * 80 + "px";
        value.style.width = v * 80 + "px";
        event.volume.pr = setTimeout(function() {
          if (!event.volume.lock) {
            event.volume.hide();
          }
        }, 1000);
      },
      bind: function() {
        dom.btn.volume.addEventListener("mouseenter", event.volume.mouseenter);
        dom.btn.volume.addEventListener("mouseleave", event.volume.mouseleave);
        dom.panel.volume.panel.addEventListener("mousedown", event.volume.mousedownTarget);
        dom.panel.volume.area.addEventListener("mouseenter", event.volume.mouseenterTarget);
        dom.panel.volume.area.addEventListener("mouseleave", event.volume.mouseleaveTarget);
      }
    },
    changeRate(v) {
      dom.video.playbackRate = v;
    }
  };
  let bind = function() {
    let x;
    //视频事件 冒泡
    for (let k in event.video) {
      dom.video.addEventListener(k, event.video[k]);
    }
    //按钮 冒泡 阻止继续冒泡
    for (let k in event.btn) {
      dom.btn[k].addEventListener("click", event.btn[k]);
      dom.btn[k].addEventListener("click", function(e) {
        e.stopPropagation();
      });
    }
    //capture 按钮 冒泡
    dom.capture.start.addEventListener("click", event.capture.start);
    dom.capture.end.addEventListener("click", event.capture.end);
    //progressbar 特效 冒泡
    dom.pgb.panel.addEventListener("mouseenter", event.pgb.mouseenter);
    dom.pgb.panel.addEventListener("mouseleave", event.pgb.mouseleave);
    //progressbar 阻止图片拖动事件 捕获
    dom.pgb.img.addEventListener("mousedown", event.pgb.img, true);
    //progressbar 拖动开始 冒泡
    dom.pgb.panel.addEventListener("mousedown", event.pgb.mousedown);
    //progressbar 拖动进行中 捕获
    document.addEventListener("mousemove", event.mousemoveall, true);
    //progressbar 拖动结束 捕获
    document.addEventListener("mouseup", event.mouseup, true);
    //快捷键绑定 冒泡
    dom.frame.addEventListener("keydown", event.keydown);
    //controller 冒泡
    dom.frame.addEventListener("mousemove", event.mousemove);
    dom.frame.addEventListener("mouseleave", event.mouseleave);
    dom.frame.addEventListener("contextmenu", function(e) {
      if (!e.target.classList.contains("bf-nomenu")) {
        e.preventDefault();
      }
    });
    //工具提示 冒泡
    x = dom.frame.querySelectorAll(".bf-tooltip");
    for (let item of x) {
      item.addEventListener("mouseenter", event.tooltip.show);
      item.addEventListener("mousemove", event.tooltip.move);
      item.addEventListener("mouseleave", event.tooltip.hide);
    }
    //选择器 冒泡
    x = dom.frame.querySelectorAll(".bf-selector");
    for (let item of x) {
      item.addEventListener("mouseenter", event.selector.show);
    }
    //点击frame事件
    dom.frame.addEventListener("mousedown", event.mousedown);
    //绑定音量调节
    event.volume.bind();
  };
  let pgbMove = function(e) {
    let w = dom.frame.scrollWidth - 40;
    let percent = (e.pageX - _config.pgbLeft) / w;
    percent = percent >= 1 ? 1 : percent <= 0 ? 0 : percent;
    let t = _config.duration * percent;

    dom.pgb.played.style.width = percent * 100 + "%";
    dom.time.innerHTML = utils.format(t) + " / " + _config.time;
    dom.pgb.point.style.left = percent * w + "px";
    dom.pgb.time = t;
  };
  let snackbar = function(content) {
    if (!_config.snackbarLock) {
      _config.snackbarLock = true;
      let d = dom.frame.querySelector(".bf-snackbar"),
        l;
      if (d != null) {
        clearTimeout(_config.snackbar);
        d.classList.add("bf-snackbar-hide");
        setTimeout(function() {
          d.remove();
        }, 130);
      }
      l = document.createElement("div");
      l.classList.add("bf-snackbar");
      l.appendChild(document.createTextNode(content));
      dom.frame.appendChild(l);
      l.style.left = dom.frame.scrollWidth / 2 - l.scrollWidth / 2 + "px";
      _config.snackbar = setTimeout(function() {
        l.classList.add("bf-snackbar-hide");
        _config.snackbar = setTimeout(function() {
          l.remove();
        }, 130);
      }, 5000);
      _config.snackbarLock = false;
    }
  };
  let screen = {
    isFull: false,
    type: 0,
    screen: function() {
      let elem = dom.frame;
      document.addEventListener("webkitfullscreenchange", screen.exitScreen);
      document.addEventListener("fullscreenchange", screen.exitScreen);
      document.addEventListener("mozfullscreenchange", screen.exitScreen);
      _config.fullscreen = true;
      screen.isFull = true;
      screen.type = 2;
      if (elem.requestFullscreen) {
        elem.requestFullscreen();
      } else if (elem.mozRequestFullScreen) {
        elem.mozRequestFullScreen();
      } else if (elem.webkitRequestFullscreen) {
        elem.webkitRequestFullscreen();
      }
      dom.btn.fullscreen.setAttribute("data-tooltip", "退出全屏");
      snackbar("已进入屏幕全屏");
    },
    client: function() {
      _config.fullbrowser = true;
      screen.isFull = true;
      screen.type = 1;
      dom.frame.classList.add("bfplayer-full");
      dom.btn.fullbrowser.setAttribute("data-tooltip", "退出全屏");
      snackbar("已进入网页全屏");
    },
    exitScreen: function() {
      if (_config.fullscreen) {
        _config.fullscreen = false;
        return;
      }
      screen.isFull = false;
      screen.type = 0;
      dom.btn.fullscreen.setAttribute("data-tooltip", "屏幕全屏");
      if (document.cancelFullScreen) {
        document.cancelFullScreen();
      } else if (document.mozCancelFullScreen) {
        document.mozCancelFullScreen();
      } else if (document.webkitCancelFullScreen) {
        document.webkitCancelFullScreen();
      }
      document.removeEventListener("webkitfullscreenchange", screen.exitScreen);
      document.removeEventListener("fullscreenchange", screen.exitScreen);
      document.removeEventListener("mozfullscreenchange", screen.exitScreen);
      snackbar("已退出屏幕全屏");
    },
    exitClient: function() {
      _config.fullbrowser = false;
      screen.isFull = false;
      screen.type = 0;
      dom.frame.classList.remove("bfplayer-full");
      dom.btn.fullbrowser.setAttribute("data-tooltip", "网页全屏");
      snackbar("已退出网页全屏");
    },
    clientToggle: function() {
      screen.isFull ? screen.type == 1 && screen.exitClient() : screen.client();
    },
    screenToggle: function() {
      screen.isFull ? screen.type == 2 && screen.exitScreen() : screen.screen();
    }
  };

  let controller = {
    show: function() {
      dom.controller.removeAttribute("style");
      dom.capture.panel.removeAttribute("style");
      dom.frame.style.cursor = "auto";
    },
    timer: function() {
      if (_config.controll.lock) {
        return;
      }
      _config.controll.time++;
      if (_config.controll.time >= 3 && !_config.controll.lock) {
        dom.controller.classList.add("bf-controller-hide");
        dom.capture.panel.classList.add("bf-controller-hide");
        setTimeout(function() {
          if (_config.controll.time >= 3 && !_config.controll.lock) {
            _config.controll.time = 0;
            controller.hidden();
          }
        }, 210);
        return;
      }

      _config.controll.pf = setTimeout(function() {
        controller.timer();
      }, 1000);
    },
    hidden: function() {
      dom.controller.classList.remove("bf-controller-hide");
      dom.capture.panel.classList.remove("bf-controller-hide");
      dom.controller.style.display = "none";
      dom.capture.panel.style.display = "none";
      dom.frame.style.cursor = "none";
    },
    reload: function() {
      if (_config.controll.lock) return;
      controller.show();
      clearTimeout(_config.controll.pf);
      _config.controll.time = 0;
      controller.timer();
    }
  };
  let utils = {
    format: function(t) {
      let tt = (Math.floor(+t) / 60).toFixed(3).split(".");
      let s = ((Number(tt[1]) / 1000) * 60).toFixed(0);
      s = s.length == 1 ? "0" + s : s;
      return tt[0] + ":" + s;
    },
    getBufferTime: function(TimeRanges) {
      let t = TimeRanges.length;
      return TimeRanges.end(t - 1);
    }
  };
  this.snackbar = snackbar;
  this.reload = function(link) {
    dom.video.pause();
    dom.pgb.played.style.width = "0";
    dom.video.src = link;
    dom.video.load();
  };
  this.showMark = function(start, end) {
    if (_config.select.enable) {
      snackbar("请先完成当前标记");
      return;
    }
    let percent = start / dom.video.duration;
    let s = dom.pgb.select.style;
    s.left = (dom.frame.scrollWidth - 40) * percent + "px";
    percent = (end - start) / dom.video.duration;
    s.display = "initial";
    s.width = percent * 100 + "%";
    dom.selectTime.innerHTML = `已选择：${start} - ${end}`;
    dom.selectTime.style.display = "initial";
    controller.reload();
    _config.controll.lock = true;
  };
  this.hideMark = function() {
    dom.selectTime.style.display = "none";
    dom.pgb.select.style.display = "none";
    _config.controll.lock = dom.video.paused;
    controller.reload();
  };
  bind();
  dom.video.src = link;
  return this;
};
export default BFPlayer;
