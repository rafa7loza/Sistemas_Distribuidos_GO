"use strict";

function saveUsername() {
  const user = document.getElementById("username").value;
  localStorage.setItem("username", user);

  // Also store an index which is the index of the last loaded message.
  localStorage.setItem("lastIndex", 0);
}
