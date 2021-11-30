const balancers = require('./balancers/client');
const machines = require('./machines/client');

const balancersClient = balancers.Client('http://localhost:8080');
const machinesClient = machines.Client('http://localhost:8080');

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

// Scenario 2: Update machine.
machinesClient.updateMachine(5, false)
    .then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Update machine response:', resp);
    })
    .catch((e) => {
        console.log(`Problem updating machine: ${e.message}`);
    });
