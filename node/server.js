const http = require('http');
const MongoClient = require('mongodb').MongoClient;

// Connection URL
const url = 'mongodb://localhost:27017';
// Database Name
const dbName = 'articulate-app';
// Create a new MongoClient
const client = new MongoClient(url, {
    useUnifiedTopology: true
});

const insertNote = async (note) => {
    try {
      await client.connect();
      console.log("Connected correctly to server");
      const db = client.db(dbName);
  
      // Insert a single document
      let r = await db.collection('notes').insertOne(note);
      if (r.insertedCount !== 1) {
          throw new Error('did not insert document');
      }
  
      // Close connection
      client.close();
      return r.insertedId;
    } catch(err) {
      console.log(err.stack);
    }
  };

  const getNotes = async () => {
      try {
          await client.connect();
          console.log('did it');
          const db = client.db(dbName);
          const notes = await db.collection('notes');
        const docs =  await notes.find();
        return docs.toArray();
      } catch(err) {
          console.log(err.stack);
      }
  }



const setupServer = () => {
http
  .createServer(async (req, res) => {
    const url = req.url;
    const method = req.method;
    console.log(method);
    if (url === "/notes" && method === 'GET') {
        const response = await getNotes();
        const json = JSON.stringify({data: response});
        console.log(response, 'response?');
        res.end(json); //end the response
        return;
    }

    if (url === '/notes' && method === 'POST') {
        req.on('data', async (chunk) => {
            const body = await JSON.parse(chunk);
            console.log(body);
        if (body.url === null || body.url === '') {
            res.statusCode = 400;
            res.end('please provide a url');
        }

        if (body.quote === null || body.url === '') {
            res.statusCode = 400;
            res.end('please provide a quote');
        }
        try {
            const note = {quote: body.quote, url: body.url}
            const response = await insertNote(note);
            console.log(response);
            const toSend = await JSON.stringify({__id: response});
            res.statusCode = 200;
            res.end(toSend);
        } catch(err) {
            console.log(err);
        }
    })

    }
  })
  .listen(3000, function() {
    console.log("server start at port 3000"); //the server object listens on port 3000
  });
}

setupServer();
