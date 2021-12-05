const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        updateMachine: (id, isWorking) => client.put('/machines', { id, isWorking })
    }

};

module.exports = { Client };
