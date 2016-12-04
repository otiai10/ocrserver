package controllers

import (
	"html/template"
	"net/http"

	m "github.com/otiai10/marmoset"
	"github.com/otiai10/ocrserver/config"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("").Parse(view)
	if err != nil {
		m.Render(w, true).JSON(http.StatusInternalServerError, m.P{
			"message": err.Error(),
		})
		return
	}
	tmpl.Execute(w, map[string]interface{}{
		"AppName": config.AppName(),
	})
	/* TODO: marmoset.View doesn't work in Heroku instance :(
	marmoset.Render(w).HTML("index", map[string]interface{}{
		"AppName": config.AppName(),
	})
	*/
}

const view = `<!DOCTYPE html>
<html>
  <head>
    <title>{{.AppName}}</title>
    <meta charset="utf-8">
    <link rel="stylesheet" href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css">
    <link rel="shortcut icon" href="/assets/favicon.png">
    <script type="text/javascript">
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
		</script>
  </head>
  <body>
    <div class="container">
      <h1>{{.AppName}}</h1>
      <pre id="response"></pre>
      <div class="row">
        <div class="col-xs-6">
          <h2>File</h2>
          <input id="file" type="file" accept="image/*"><br>
        </div>
        <div class="col-xs-6">
          <h2>or base64</h2>
          <textarea id="base64" class="form-control"></textarea>
        </div>
      </div>
      <div class="row">
        <div class="col-xs-12">
          <h2>options</h2>
          <textarea id="options" class="form-control" rows="4">{
            "whitelist": "OCR1234567890",
            "trim": "\n"
}</textarea>
          <ul>
            <li><code>whitelist (optional): allowed characters</code></li>
            <li><code>trim (optional): characters to trim out</code></li>
          </ul>
        </div>
      </div>
      <div class="pull-right"><button id="submit" class="btn btn-default">submit</button></div>
    </div>
  </body>
</html>`
