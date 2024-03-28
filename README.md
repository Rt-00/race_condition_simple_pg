# Race Condition in PostgreSQL and Golang

This project demonstrates a race condition around concurrent transactions that withdraw a quantity of 1 from the stock in a PostgreSQL database.

## How It Works

The project uses Docker Compose to set up a testing environment with a PostgreSQL database. It simulates the situation in which multiple concurrent transactions attempt to remove 1 unit from inventory simultaneously. This situation can lead to unexpected results, such as a negative inventory balance or overlapping transactions. Therefore, this solution can be used in systems with a lot of access competition, as a lock is not used during processing.

## Prerequisites

Make sure you have Docker and Docker Compose installed on your system before running this project.

## How to Use

1. Clone this repository on your local machine:
```
git clone https://github.com/Rt-00/race_condition_simple_pg.git
```    

2. Navigate to the project directory:
```
cd race_condition_simple_pg
```

3. Run Docker Compose to start the PostgreSQL container:
```
docker-compose up -d
```

4. Wait until the PostgreSQL container is up and running.

5. Run the Go script to simulate concurrent transactions:
```
go run main.go
```

6. Observe the results and behaviors of the concurrent transactions.

## Contribution

Contributions are welcome! Feel free to open an issue or submit a pull request with improvements, bug fixes, or any other suggestions you may have.

## License

This project is licensed under the [MIT License](LICENSE).
