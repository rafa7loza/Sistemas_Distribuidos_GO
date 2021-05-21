function fetch() {
  handleNameInput()
  loadData("students", "studentNameSelect")
  loadData("students", "studentAvgSelect")
  loadData("subjects", "subjectAvgSelect")
}

function handleNameInput() {
  const textInput = document.getElementById("textInput");
  const selectInput = document.getElementById("selectInput");
  const studentExists = document.getElementById("studentNameCheckbox");

  values = studentExists.checked ? [false,true] : [true,false];
  selectInput.hidden = values[0];
  textInput.hidden = values[1];
}

function loadData(name, elementId) {
  const elem = document.getElementById(elementId)
  var result = null;
  var xmlhttp = new XMLHttpRequest();
  xmlhttp.open("GET", "data/"+name+".json", false);
  xmlhttp.send();

  if (xmlhttp.status != 200) {
    console.log("Error loading the data");
    return ;
  }

  var option;
  result = JSON.parse(xmlhttp.responseText);
  result[name].forEach((subject) => {
    console.log(subject);
    option = document.createElement("option");
    option.text = subject;
    elem.add(option)
  });
}
