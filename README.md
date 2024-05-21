# Example setup of service
## Establishing environment
For each service, there should be files for both the application server and the database. Therefore, you will need the following files:
* `Dockerfile`: Contains instructions for building the application server and the database. Currently, we are using PostgreSQL.
* `docker-compose.yml`: Defines the configuration to run multiple containers.
* `.env`: Contains configurations.

You can find the actual contents in [the commit](https://github.com/ryuju0911/example-setup-of-service/commit/40e15fd021a42db0a856d68096e5f0b6f6f1a0c)

Please note:

`.env` is ignored by Git. Therefore, you need to create an example configuration in `.env.sample`, copy it to `.env`, and edit it if necessary.

## Defining service API
Define your service API in `design/design.go` to enable Goa to generate boilerplate code.
You can refer to [the example design](https://github.com/ryuju0911/example-setup-of-service/commit/ea03059b5e9f7da2768ecc29110b780ac1e9fcb8).
However, it's highly recommended to go through the official tutorial.

After designing your API, run `goa gen ${your service name}/design` to generate the code.

Next, execute the `goa example` command to generate a basic implementation of the service.
This will create server files that initialize an HTTP server and client files capable of making requests to the server.

`goa example ${your service name}/design`

However, note that the code generated by `goa example` does not handle database connections.
You'll need to add code to handle this, as demonstrated in [this commit](https://github.com/ryuju0911/example-setup-of-service/commit/d5764fe6f6e938704529e20989f7f03cec704cfb).

## Setting Up Schema Migrations
Before interacting with the database, it's essential to create a table for storing your data. We use [golang-migrate](https://github.com/golang-migrate/migrate) to execute schema migrations.
These migrations are executed each time the containers are brought up, such as when running `docker-compose up`.

There's an important consideration to keep in mind. Because docker-compose doesn't check if the PostgreSQL container is already up before executing migrations,
we need to ensure that it's ready before proceeding. To handle this, we use a script called `wait_for_postgres.sh`, which must be executed before performing migrations.
For implementation details, please refer to [this commit](https://github.com/ryuju0911/example-setup-of-service/commit/fbb9689f77486f396b6af350bedb2fbb3a978a16).

## Implementing Your API
With the server generated, the API skeleton in place, and the table created, you can now proceed to implement your API.
Please refer to [this commit](https://github.com/ryuju0911/example-setup-of-service/commit/dff6b68c027c0b9e48c65eda99d736e56a02976b) for an implementation example.
