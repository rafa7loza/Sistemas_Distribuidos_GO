"use strict";

function saveUsername() {
  const user = document.getElementById("username").value;
  localStorage.setItem("username", user);
}
