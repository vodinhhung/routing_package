# Routing Algorithm Service

A sophisticated delivery route optimization service built in Go that provides intelligent routing solutions for logistics and delivery operations.

## What is this project?

This is a REST API service that optimizes delivery routes using advanced algorithms to minimize travel distance, time, and cost. The service manages delivery orders with geographic coordinates and calculates the most efficient routes for delivery vehicles starting from a warehouse location.

### Key Features

- **Multi-level Route Optimization**: Uses a sophisticated two-tier optimization approach
  - **Inter-region optimization**: Uses Haversine distance calculation for routing between different geographic regions (VN, US, CA)
  - **Intra-region optimization**: Uses Euclidean distance for efficient routing within each region
- **Dynamic Programming Algorithm**: Implements the Traveling Salesman Problem (TSP) solution using dynamic programming for optimal route calculation
- **Cost Calculation**: Automatically calculates delivery costs based on distance (mileage) and time factors
- **RESTful API**: Provides clean REST endpoints for order management, route creation, and cost optimization
- **MySQL Database**: Persistent storage with GORM ORM for data management

### Routing Algorithm Details

The service implements a sophisticated routing optimization strategy:

1. **Regional Grouping**: Orders are first grouped by region code (VN, US, CA)
2. **Inter-regional Routing**: Uses Haversine distance (great-circle distance) to determine the optimal order of visiting regions from the warehouse
3. **Intra-regional Routing**: Within each region, uses Euclidean distance with dynamic programming TSP solver for optimal order delivery sequence
4. **Cost Optimization**: Calculates total cost using the formula: `cost = mileage × $18.75 + time × $0.30`

## API Endpoints

### Orders
- `POST /orders/new` - Create a new delivery order
- `GET /orders/detail?order_id={id}` - Get order details

### Routes
- `POST /routes/new` - Create an optimized route
- `GET /routes/detail?route_id={id}` - Get route details

### Cost Optimization
- `GET /cost/lowest?order_ids={ids}&warehouse_longitude={lng}&warehouse_latitude={lat}` - Find the lowest cost route for given orders

## Prerequisites

- Go 1.23.4 or later
- MySQL database server
- Make utility

## Database Setup

1. Create a MySQL database and update the connection string in `src/dependency/common.go`:
   ```go
   dsn := "user:password@tcp(127.0.0.1:3306)/your_db_name?parseTime=true"
   ```

2. Run the database migration scripts:
   ```bash
   mysql -u your_user -p your_database < dep/mysql/order/create.sql
   mysql -u your_user -p your_database < dep/mysql/route/create.sql
   ```

## Build and Run

The project uses a Makefile for easy build and execution:

### Available Make Commands

- **Clean build artifacts:**
  ```bash
  make clean
  ```

- **Build the project:**
  ```bash
  make build
  ```

- **Build and run the service:**
  ```bash
  make run
  ```

### Manual Build and Run

You can also build and run manually:

```bash
# Build the application
go build -o bin/main cmd/main.go

# Run the application
./bin/main
```

The service will start on port 8080 and log startup information to the console.

## Project Structure

```
routing_package/
├── bin/                    # Build outputs
├── cmd/
│   └── main.go            # Application entry point
├── config/
│   └── test.json          # Configuration files
├── dep/mysql/             # Database schema files
│   ├── order/create.sql   # Order table schema
│   └── route/create.sql   # Route table schema
├── src/
│   ├── api/               # REST API handlers
│   │   ├── cost.go        # Cost optimization endpoints
│   │   ├── order.go       # Order management endpoints
│   │   ├── route.go       # Route management endpoints
│   │   └── validation.go  # Input validation
│   ├── dependency/        # Database models and connections
│   │   ├── client.go      # Database client
│   │   ├── common.go      # Database initialization
│   │   ├── order.go       # Order model and operations
│   │   └── route.go       # Route model and operations
│   └── internal/          # Core business logic
│       ├── cost.go        # Cost calculation algorithms
│       ├── helper.go      # Routing optimization algorithms
│       ├── order.go       # Order business logic
│       ├── region.go      # Regional coordinate definitions
│       └── route.go       # Route business logic
├── utils/
│   └── utils.go           # Utility functions
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
└── Makefile              # Build automation
```

## Dependencies

- **GORM**: Object-relational mapping library for Go
- **MySQL Driver**: Database connectivity for MySQL
- **Standard Go Libraries**: HTTP server, JSON encoding, mathematical calculations

## Algorithm Complexity

- **Time Complexity**: O(n² × 2ⁿ) for the dynamic programming TSP solution within each region
- **Space Complexity**: O(n × 2ⁿ) for the DP memoization table
- **Practical Performance**: Optimized for real-world delivery scenarios with regional pre-grouping to reduce problem size

## Example Usage

1. Start the service: `make run`
2. Create delivery orders with coordinates and region codes
3. Request route optimization with warehouse location
4. Receive optimized delivery sequence with cost, distance, and time estimates

The service provides intelligent routing that significantly reduces delivery costs and time compared to naive routing approaches.