function main() {
  handleNameInput()
}

function handleNameInput() {
  const textInput = document.getElementById("textInput")
  const selectInput = document.getElementById("selectInput")
  const studentExists = document.getElementById("studentNameCheckbox")

  values = studentExists.checked ? [false,true] : [true,false]
  selectInput.hidden = values[0]
  textInput.hidden = values[1]
}
