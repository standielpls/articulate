const noteApp = {};

// temporarily fetch from cloudflare as user-id
// fetch the data from cf regarding that user-id
// display articles
noteApp.loadContent = async () => {
    try {
        const response = await fetch('https://www.cloudflare.com/cdn-cgi/trace');
        const data = await response.text();
        const ipaddr = data.match(/ip=(.+)/)[1];
        const secondResp = await fetch(`https://us-central1-lancelot-274021.cloudfunctions.net/listNotesByUserID?user_id=${ipaddr}`);
        if (secondResp.status !== 200) {
            noteApp.displayError()
            return;
        }
        const json = await secondResp.json();
        const { articles } = json;
        if (articles == null || articles.length < 1) {
            noteApp.displayEmpty()
            return
        }
        noteApp.displayArticles(articles)
    } catch (err) {
        console.log(err);
    }
}

// display empty state
noteApp.displayEmpty = () => {
    const parent = document.getElementById("content")
    const div = document.createElement("div");
    div.className = "hold-note";

    const noteEle = document.createElement("p");
    const noteNode = document.createTextNode("No articles found, add some using the chrome extension to start!");
    noteEle.appendChild(noteNode)
    div.append(noteEle)
    parent.append(div)
}

// display empty state
noteApp.displayError = () => {
    const parent = document.getElementById("content")
    const div = document.createElement("div");
    div.className = "hold-note";

    const noteEle = document.createElement("p");
    const noteNode = document.createTextNode("Something went wrong!");
    noteEle.appendChild(noteNode)
    div.append(noteEle)
    parent.append(div)
}

// display the articles, creating the dom elements necessary
noteApp.displayArticles = (articles) => {
    for (note of articles) {
        const { url, article, comment } = note;


        let headerEle;
        if (url.includes("youtube")) {
            
            headerEle = document.createElement("a");
            const imgEle = new Image();
            headerEle.appendChild(imgEle)
            headerEle.href = url;

            const newUrl = new URL(url);
            const urlParams = new URLSearchParams(newUrl.search);
            const ytId = urlParams.get('v');
            const thumbnail = `https://img.youtube.com/vi/${ytId}/maxresdefault.jpg`
            imgEle.src = thumbnail;
        } else {
            // create <blockquote>:
            headerEle = document.createElement("blockquote");
            const quoteNode = document.createTextNode(comment);
            headerEle.appendChild(quoteNode)
        }

        // create <p>:
        const noteEle = document.createElement("p");
        const noteNode = document.createTextNode(article);
        noteEle.appendChild(noteNode)

        // create <a>:
        const aNode = document.createElement("a")
        aNode.href = url;
        aNode.setAttribute("target", "_blank")
        aNode.innerHTML = url;

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

