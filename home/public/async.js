const noteApp = {};

noteApp.loadContent = async () => {
    try {
        const response = await fetch('https://www.cloudflare.com/cdn-cgi/trace');
        const data = await response.text();
        const ipaddr = data.match(/ip=(.+)/)[1];
        const secondResp = await fetch(`https://us-central1-lancelot-274021.cloudfunctions.net/listNotesByUserID?user_id=${ipaddr}`);
        if (secondResp.status !== 200) {
            console.log("problem:", response.status);
            return;
        }
        const json = await secondResp.json();
        const { articles } = json;
        noteApp.displayArticles(articles)
    } catch(err) {
        console.log(err);
    }
}

noteApp.displayArticles = (articles) => {
    for (note of articles) {
        const { url, article, comment } = note;
        
        // create <a>:
        const aNode = document.createElement("a")
        aNode.href = url;
        aNode.setAttribute("target", "_blank")
        aNode.innerHTML = url;
        
        // create <blockquote>:
        const headerEle = document.createElement("blockquote");
        const quoteNode = document.createTextNode(comment);
        headerEle.appendChild(quoteNode)
        
        // create <p>:
        const noteEle = document.createElement("p");
        const noteNode = document.createTextNode(article);
        noteEle.appendChild(noteNode)
        
        // create containing element and append to page:
        const parent = document.getElementById("content")
        const div = document.createElement("div");
        div.className = "hold-note";

        div.appendChild(headerEle);
        div.appendChild(noteEle);
        div.appendChild(aNode);

        parent.appendChild(div)
      }
}

noteApp.init = () => {
    noteApp.loadContent();
}

// when the dom is loaded
document.addEventListener("DOMContentLoaded", () => {
    noteApp.init();
})

