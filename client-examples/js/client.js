const https = require('https');

https.get('https://api.terminjetzt.com/appointments/latest', (res) => {
  let data = '';
  res.on('data', chunk => data += chunk);
  res.on('end', () => console.log(JSON.parse(data)));
});
