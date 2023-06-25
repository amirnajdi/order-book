# Order Book
The Order Book is a simple order tracking system for a cryptocurrency website. It receives orders from Kafka, stores them in a database, and provides an API for users to access the order book for each symbol.

## API Endpoint
To retrieve the order book for a specific symbol, you can use the following API endpoint:

`localhost:{APP_PORT}/orders/{symbol}`

You can also include a `limit` query parameter to specify the maximum number of orders to retrieve (the default value is 100).

## Technologies Used
The Order Book system is built using the following technologies:

- Go
- Kafka
- MariaDB
- Docker
- Docker Compose

## Getting Started
To start the project, follow these steps:

1. Copy the `.env.example` file and rename it to `.env`. Update the configuration settings in the `.env` file according to your requirements.
2. Run `docker-compose up -d` to start the necessary containers.
3. Access PHPMyAdmin at the default address `localhost:7070`. 
4. Import the database backup located at `database/order_books.sql.gz` into your MariaDB instance.
5. Run the command `docker exec -it order-book-app-1 /insert-test-data` to insert test data into the system. Replace `order-book-app-1` with the appropriate container name on your system.

## Benchmark
The benchmark was conducted with 1,641,171 records.

![Benchmark Chart](https://imgtr.ee/images/2023/06/21/ZiEuQ.png)

Please note that the benchmark chart provided above showcases the performance of the Order Book system.

Feel free to make any further modifications or additions based on your specific needs.
