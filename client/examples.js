const balancers = require('./balancers/client');

const balancersClient = balancers.Client('http://localhost:8080');

// Scenario 1: Display list of balancers.
balancersClient.listBalancers()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('List of balancers:');
        list.forEach((c) => console.log(c.name));
    })
    .catch((e) => {
        console.log(`Problem listing list of balancers: ${e.message}`);
    });
