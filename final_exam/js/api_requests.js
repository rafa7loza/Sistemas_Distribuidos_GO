"use strict";

function postMessage() {
  const msg = document.getElementById("message").value;
  const username = localStorage.getItem("username");

  const t = new Date();
  const dateStr = t.getDate() + "-" + (t.getMonth()+1) + "-" + t.getFullYear() + " "
    + t.getHours() + ":" + t.getMinutes();

  const payload = {
    From: username,
    Content: msg,
    Date: dateStr,
  };

  console.log(payload);

  let myProm = new Promise((resolve, reject) => {
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/API/message", true);
    xhr.onload = () => {
      if (xhr.status >= 200 && xhr.status < 300) {
        resolve(xhr.response);
      } else {
        reject({
          status: xhr.status,
          statusText: xhr.statusText
        });
      }
    };
    xhr.onerror = () => {
      reject({
        status: xhr.status,
        statusText: xhr.statusText
      });
    };

    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.send(JSON.stringify(payload));
  });

  myProm.then( ok => {
    console.log(ok);
  }).catch( bad => {
    console.log(bad);
  });
}
