chrome.runtime.onMessage.addListener(function (msg, sender, sendResponse) {
    if (msg.text === 'report_back') {
        sendResponse(document.getElementById("movie-player").getCurrentTime());
    }
});
