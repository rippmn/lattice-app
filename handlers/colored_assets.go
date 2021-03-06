package handlers

import (
	"text/template"
)

var styledColoredTemplate = template.Must(template.New("c_experiment").Parse(`
<html>
<head>
<style>
body {
    font-family: "helveticaneue-light";
    font-size: 16px;
    color: #333;
    position: absolute;
    margin:0;
    width:100%;
    height:100%;
}

dt {
  color:#777;
}

.envs {
  margin:10px
}

.hello {
  position:absolute;
  top:0;
  height:120px;
  left:0;
  right:0;
  text-align:center;
  font-size:80px;
  font-weight:bold;
  line-height:120px;
  color: {{index .Colors 0 }};
}

.my-index {
  position:absolute;
  top:120px;
  height:30px;
  color: #333;
  font-size:30px;
  line-height:30px;
  left:0;
  right:0;
  text-align:center;
  color: {{index .Colors 1}};
}

.index {
  position:absolute;
  top:176px;
  height:120px;
  left:0;
  right:0;
  color: #fff;
  font-size: 80px;
  line-height: 120px;
  background-color:{{index .Colors 2}}; 
  text-align:center;
}

.mid-color {
  position:absolute;
  top:296px;
  height:120px;
  left:0;
  right:0;
  color: #fff;
  font-size: 30px;
  line-height: 120px;
  background-color:{{index .Colors 1}}; 
  text-align: center;
}

.bottom-color {
  position:absolute;
  top:416px;
  bottom:0;
  left:0;
  right:0;
  background-color:{{index .Colors 0}}; 
}


</style>
</head>
<body class="{{.Class}}">
{{.Body}}
</body>
</html>
`))

type ColoredBody struct {
	Body  string
	Class string
	Colors []string
}

