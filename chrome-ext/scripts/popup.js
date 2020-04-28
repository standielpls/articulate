let title;
let url;

function apiAddNotes() {
  let quote = document.getElementById('output').textContent;
  let note = document.getElementById('comments').value;

  let ipaddr = "555";
  fetch('https://www.cloudflare.com/cdn-cgi/trace').then((response) => {
    response.text().then(function (data) {
      ipaddr = data.match(/ip=(.+)/)[1]

      fetch("https://us-central1-lancelot-274021.cloudfunctions.net/createNote", {
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify(
          {
            "user_id": ipaddr,
            "comment": quote,
            "article": note,
            "url": url
          })
      }) // Call the fetch function passing the url of the API as a parameter
        .then(function () {
          document.getElementById("fireAddNotes").textContent = "Added!"
        })
        .catch(function (err) {
          // This is where you run code if the server returns any errors
          console.log("error: ", err)
        });
    })
  });

}


// document.addEventListener("DOMContentLoaded", () => {
//   // document.getElementById("fireAddNotes").addEventListener("click", apiAddNotes);

//   chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
//     const { url, title } = tabs[0];

//     document.getElementById("app-name").textContent = title;
//     // if (url.includes("youtube")) {
//     //   n = document.createElement("a").setAttribute("href", `${url}?t=${timestamp}`)
//     //   document.getElementById("output").appendChild(n)
//     // } else {
//       chrome.tabs.executeScript(
//         {
//           code: "window.getSelection().toString();"
//         },
//         function (selection) {
//           if (!selection[0] || selection[0] === '') {
//             document.getElementById("output").textContent = 'Highlight text for it to appear here!';
//           } else {
//             document.getElementById("output").textContent = selection[0];
//           }
//         }
//       );
//     // };
//   });
// });

function doSomethingWithTimeStamp(domContent) {
  console.log("this is my something with time stamp: ", domContent)
}
chrome.browserAction.onClicked.addListener(function (tab) {
  chrome.tabs.sendMessage(tab.id, {text: 'report_back'}, doSomethingWithTimeStamp);
});