const http = require('../common/http');

const Client = (baseUrl) => {
    const client = http.Client(baseUrl);

    return {
        listBalancers: () => client.get('/balancers')
    }
};

module.exports = { Client };
