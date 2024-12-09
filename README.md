# Weather Check App

This is a weather query application built with Go (Golang). The app integrates with two weather APIs and provides the average temperature for a given location. To avoid overloading the external APIs, incoming requests are batched and served after a 5-second waiting period. The system logs all weather queries to a SQLite database.

## Features

- **Single Endpoint**: `/weather?q=<location>`
    - Accepts a query parameter `location` (e.g., `Istanbul`).
    - Returns the average temperature based on data from two external weather services.

- **Request Batching**:
    - Requests for the same location within 5 seconds are batched together, allowing the app to only make a single API request.
    - If there are more than 10 requests for the same location during the waiting period, the system makes the request immediately.

- **Temperature Calculation**:
    - The app receives temperature data from two different weather services (WeatherAPI and WeatherStack).
    - The final temperature is the average of the two temperatures with a slight decrease applied.

- **Logging**:
    - Weather queries, including location, service temperatures, request count, and timestamps, are logged to a SQLite database for future reference.

- **Max Wait Time**:
    - Requests will wait for up to 5 seconds for batching. If no other requests are received in that time, the query is processed immediately.

## APIs Used

1. **WeatherAPI**
    - API Key: `147d644004414106a2f75650232001`
    - Endpoint: `http://api.weatherapi.com/v1/forecast.json?key=<API_KEY>&q=<location>&days=1&aqi=no&alerts=no`
    - Temperature Data: `current.temp_c`

2. **WeatherStack**
    - API Key: `838c0d5e8fcc1dbbc66e8c1c0a14c6e5`
    - Endpoint: `http://api.weatherstack.com/current?access_key=<API_KEY>&query=<location>`
    - Temperature Data: `current.temperature`

## Database

- **SQLite Database**: `weather.sqlite`
- **Table**: `weather_queries`
    - Columns:
        - `id` (auto-increment)
        - `location` (text)
        - `service_1_temperature` (float)
        - `service_2_temperature` (float)
        - `request_count` (integer)
        - `created_at` (timestamp)

## Installation

### Prerequisites

Ensure you have the following installed:
- Go 1.17 or later
- SQLite (for local development, or you can use a different database)
- API keys for WeatherAPI and WeatherStack

### Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/weather-check-app.git
    cd weather-check-app
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Create the SQLite database file:

   The application will automatically create the database if it doesn't exist, but you can manually initialize it by running:

    ```bash
    touch weather.sqlite
    ```

4. Set up your environment variables or API keys in the code. For example, add your keys to the services like `weatherapi.com` and `weatherstack.com` in the code itself.

## Usage

### Running the Application

To start the application, use the following command:

```bash
go run main.go
