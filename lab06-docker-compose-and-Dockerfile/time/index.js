'use strict';

const express = require('express');

// Constants
const PORT = process.env.PORT || 8080;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/time', (req, res) => {
  res.send(new Date().toISOString());
});

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);