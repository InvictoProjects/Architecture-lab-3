const http = require('http');

const request = (url, options, callback) => http.request(url, options, res => {
    const statusCode = res.statusCode;
    const contentType = res.headers['content-type'];

    let error;
    if (statusCode !== 200) {
        error = new Error('Request Failed.\n' +
            `Status Code: ${statusCode}`);
    } else if (!/^application\/json/.test(contentType)) {
        error = new Error('Invalid content-type.\n' +
            `Expected application/json but received ${contentType}`);
    }
    if (error) {
        callback(error, null);
        res.resume();
    }

    res.setEncoding('utf8');
    let rawData = '';
    res.on('data', chunk => {
        rawData += chunk;
    });
    res.on('end', () => callback(rawData));
});

const get = url => new Promise((resolve, reject) => {
    const req = request(url, { method: 'GET' }, data => {
        try {
            const parsedData = JSON.parse(data);
            resolve(parsedData);
        } catch (e) {
            reject(e);
        }
    });

    req.on('error', e => {
        reject(e);
    });

    req.end();
});

const put = (url, data) => new Promise((resolve, reject) => {
    const putData = JSON.stringify(data);

    const options = {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Content-Length': Buffer.byteLength(putData)
        }
    };

    const req = request(url, options, data => {
        try {
            const parsedData = JSON.parse(data);
            resolve(parsedData);
        } catch (e) {
            reject(e);
        }
    });

    req.on('error', e => {
        reject(e);
    });

    req.write(putData);
    req.end();
});

const Client = baseUrl => ({
    get: path => get(baseUrl + path),
    put: (path, data) => put(baseUrl + path, data)
});

module.exports = { Client };
