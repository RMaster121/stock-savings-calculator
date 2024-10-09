# Stock Savings Calculator

**Stock Savings Calculator** is a Go-based application designed to help you track the value of your stock investments over time. By fetching stock and forex data from [Polygon.io](https://polygon.io), the calculator estimates the value of your investments in your local currency and provides detailed weekly and total reports.

**Remember, that these are past results and do not predict future data. Be careful!**

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- Fetches historical weekly stock and forex data from **Polygon.io**.
- Calculates the value of stock investments over a specified time period.
- Converts investment values to your local currency.
- Generates detailed weekly and cumulative investment reports.

---

## Installation

### 1. Clone the Repository
Start by cloning the project to your local machine:
```sh
git clone https://github.com/yourusername/stock-savings-calculator.git
cd stock-savings-calculator
```

### 2. Install Dependencies
Ensure you have Go installed on your system. Then, run:
```sh
go mod tidy
```

---

## Usage

### 1. Set Up Your API Key
Obtain an API key from [Polygon.io](https://polygon.io) and ensure it’s accessible for the application.

### 2. Run the Application
Use the following command to run the application:
```sh
go run main.go --stock=NVDA --forex-query=PLN --weeks-ago=5 --money=100 --api-key=YOUR_API_KEY
```

### 3. Command Line Arguments
| Argument        | Description                                              | Example             |
|-----------------|----------------------------------------------------------|---------------------|
| `--stock`       | Stock ticker symbol.                                      | `NVDA`              |
| `--forex-query` | Local currency code for conversion (e.g., PLN for Polish Zloty). | `PLN`               |
| `--weeks-ago`   | Number of weeks to go back for investment calculation. Default is 5. | `5`                 |
| `--money`       | Amount of money to invest every week. Default is 100.     | `100`               |
| `--api-key`     | Your API key from Polygon.io.                             | `YOUR_API_KEY`      |

---

## Configuration

All configurations can be passed via command line arguments. Make sure to provide a valid API key from Polygon.io.

---

## Project Structure

```
.vscode/
    └── launch.json
api/
    └── polygon.go
go.mod
main.go
models/
    └── stock_data.go
```

- **`.vscode/launch.json`**: Visual Studio Code configuration for debugging.
- **`api/alphavantage.go`**: Handles fetching and processing of data from Polygon.io.
- **`go.mod`**: Go module definition and dependencies.
- **`main.go`**: Main application entry point.
- **`models/stock_data.go`**: Contains data models used throughout the application.

---

## Contributing

Contributions are welcome!
---

## License

This project is licensed under the MIT License. For more details, refer to the [LICENSE](./LICENSE) file.