// fires when user clicks on extensions icon
chrome.browserAction.onClicked.addListener(function (tab) {
  chrome.browserAction.setPopup({
    popup: "popup.html"
  });
  // chrome.tabs.executeScript(tab.id, {
  //   "file": "content.js"
  // }, function () { // Execute your code
  //   console.log("Script Executed .. "); // Notification on Completion
  // }); // working example

  // chrome.tabs.sendMessage(tab.id, { greeting: "barney" }, function (response) {
  chrome.tabs.executeScript(tab.id, {
    "code": "yt = document.querySelector('.video-stream').currentTime; yt"
  }, function (data) {
    console.log(data[0])
    // alert(data[0])
    document.getElementById("output").textContent = data[0]
  });
  // alert(response.ytTime);
});
