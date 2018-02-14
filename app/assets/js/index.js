"use strict"; // ES6
window.onload = () => {

  var http = {
    post: (path, data) => {
      return new Promise((resolve, reject) => {
        var xhr = new XMLHttpRequest();
        xhr.open("POST", path, true);
        xhr.onreadystatechange = () => {
          if (xhr.readyState == XMLHttpRequest.DONE) return resolve(xhr);
        };
        xhr.send(data);
      });
    }
  };

  var ui = {
    output:    document.getElementById("output"),
    image:     document.querySelector("img#img"),
    btnFile:   document.getElementById("by-file"),
    btnBase64: document.getElementById("by-base64"),
    cancel:    document.getElementById("cancel-input"),
    file:      document.getElementById("file"),
    submit:    document.getElementById("submit"),
    show:      uri => ui.image.setAttribute("src", uri),
    clear:     () => { ui.image.setAttribute("src", ""), ui.file.value = ''; },
  };

  ui.file.addEventListener("change", ev => {
    if (!ev.target.files || !ev.target.files.length) return null;
    const r = new FileReader();
    r.onload = e => ui.show(e.target.result);
    r.readAsDataURL(ev.target.files[0]);
  });
  ui.btnFile.addEventListener("click", () => ui.file.click());
  ui.btnBase64.addEventListener("click", () => {
    const uri = window.prompt("Please paste your base64 image URI");
    if (uri) { ui.clear(); ui.show(uri); }
  });
  ui.cancel.addEventListener("click", () => ui.clear());
  ui.submit.addEventListener("click", () => {
    ui.output.innerText = "{}";
    const req = generateRequest();
    if (!req) return;
    http.post(req.path, req.data).then(xhr => {
      ui.output.innerText = `${xhr.status} ${xhr.statusText}\n-----\n${xhr.response}`;
    });
  })

  var generateRequest = () => {
    var req = {path: "", data: null};
    if (ui.file.files && ui.file.files.length != 0) {
      req.path = "/file";
      req.data = new FormData();
      req.data.append("file", ui.file.files[0]);
    } else if (/^data:.+/.test(ui.image.src)) {
      req.path = "/base64";
      var data = {base64: ui.image.src};
      req.data = JSON.stringify(data);
    } else {
      return window.alert("no image input set");
    }
    return req;
  };
};
