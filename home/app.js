fetch("http://localhost:5000/notes") // Call the fetch function passing the url of the API as a parameter
  .then(function(response) {
    console.log(response.status)
    if (response.status !== 200) {
      console.log("Looks like there was a problem. Status Code: " + response.status);
      return;
    }
    // Examine the text in the response
    response.json().then(function(notes) {
        var note;
        for (note of notes) {
            var para = document.createElement("p");
            var quote = document.createTextNode(note['quote']);
            document.createElement("br")
            var comment = document.createTextNode(note['note']);
            para.appendChild(quote);
            para.appendChild(comment);
            document.getElementById("content").appendChild(para);
        }
        });
  })
  .catch(function(err) {
    // This is where you run code if the server returns any errors
    console.log('Fetch Error :-S', err);
  });

