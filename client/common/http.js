const http = require("http");

const Client = baseUrl => {
    return {
        get: path => new Promise((resolve, reject) => {
            http.get(baseUrl + path, res => {
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
                    reject(error);
                    res.resume();
                }

                res.setEncoding('utf8');
                let rawData = '';
                res.on('data', chunk => {
                    rawData += chunk;
                });
                res.on('end', () => {
                    try {
                        const parsedData = JSON.parse(rawData);
                        resolve(parsedData);
                    } catch (e) {
                        reject(e);
                    }
                });
            }).on('error', e => {
                reject(e);
            });
        }),
        put: async (path, data) => new Promise((resolve, reject) => {
            const putData = JSON.stringify(data);

            const options = {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'Content-Length': Buffer.byteLength(putData)
                }
            };

            const req = http.request(baseUrl + path, options, res => {
                res.setEncoding('utf8');
                let rawData = '';
                res.on('data', chunk => {
                    rawData += chunk;
                });
                res.on('end', () => {
                    try {
                        const parsedData = JSON.parse(rawData);
                        resolve(parsedData);
                    } catch (e) {
                        reject(e);
                    }
                });
            })

            req.on('error', e => {
                reject(e);
            });

            req.write(putData);
            req.end();
        })
    }
};

module.exports = {Client};
