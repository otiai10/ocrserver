"use strict"; // ES6
window.onload = () => {
  var http = {
    post: (path, data) => {
      return new Promise((resolve, reject) => {
        var xhr = new XMLHttpRequest();
        xhr.open("POST", path, true);
        xhr.onreadystatechange = () => {
          if (xhr.readyState != XMLHttpRequest.DONE) return;
          if (xhr.status != 200) return reject(xhr.response);
          return resolve(xhr.response);
        };
        xhr.send(data);
      });
    }
  };
  var render = {
    response: document.getElementById("response"),
    _options: document.getElementById("options"),
    success: (res) => {
      render.response.innerHTML = JSON.stringify(JSON.parse(res), null, 2);
    },
    error: (res) => {
      try {
        render.response.innerHTML = JSON.stringify(JSON.parse(res), null, 2);
      } catch (e) {
        render.response.innerHTML = JSON.stringify(res, null, 2);
      }
    },
    options: (json) => {
      render._options.value = JSON.stringify(json, null, 8);
    }
  };
  var generateRequest = () => {
    var files = document.getElementById("file").files;
    var base64 = document.getElementById("base64").value;
    var options = document.getElementById("options").value;
    try {
      options = JSON.parse(options);
    } catch (e) {
      options = {};
    }
    render.options(options);
    var req = {path: "", data: null};
    if (files.length > 0) {
      req.path = "/file";
      req.data = new FormData();
      req.data.append("file", files[0]);
      req.data.append("whitelist", options.whitelist || "");
      req.data.append("trim", options.trim || "");
    } else if (base64) {
      req.path = "/base64";
      var data = {base64: base64};
      data.whitelist = options.whitelist;
      data.trim = options.trim;
      req.data = JSON.stringify(data);
    }
    return req;
  };
  document.getElementById("submit").addEventListener("click", () => {
    var req = generateRequest();
    http.post(req.path, req.data).then(render.success).catch(render.error);
  });
};
