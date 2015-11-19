"use strict"; // ES6
window.onload = () => {
  // utils
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
    success: (res) => {
      render.response.className = "ok";
      render.response.innerHTML = JSON.parse(res).result;
    },
    error: (res) => {
      render.response.className = "ng";
      render.response.innerHTML = JSON.parse(res).error;
    }
  };

  // file
  document.getElementById("file-submit").addEventListener("click", () => {
    var data = new FormData();
    var files = document.getElementById("file").files;
    if (files.length != 0) data.append("file", files[0]);
    http.post("/file", data).then(render.success).catch(render.error);
  });
  // base64
  document.getElementById("base64-submit").addEventListener("click", () => {
    var data = {base64: document.getElementById("base64").value};
    http.post("/base64", JSON.stringify(data)).then(render.success).catch(render.error);
  });
};
