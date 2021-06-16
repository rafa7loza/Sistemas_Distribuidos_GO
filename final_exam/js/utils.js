"use strict";

function saveUsername() {
  const user = document.getElementById("username").value;
  localStorage.setItem("username", user);

  // Store an index which is the index of the last loaded message.
  localStorage.setItem("lastIndex", 0);
  // Also store the chat name
  document.getElementsByName('chat')
    .forEach(item => {
      if(item.checked) {
        console.log(item);
        localStorage.setItem("chatName", item.value);
      }
    });
}
