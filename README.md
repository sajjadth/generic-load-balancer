# Generic Proxy Load Balancer

The **Generic Proxy Load Balancer** is a Go-based load balancer that distributes incoming traffic across multiple proxy instances, an efficient and scalable way to manage load distribution. It is designed to work seamlessly with the [Generic NodeJS Proxy](https://github.com/sajjadth/generic-nodejs-proxy) project to create a scalable proxy solution.

## Features

- **Round-Robin Load Balancing:** Distributes requests evenly across all proxy instances using a round-robin algorithm.
- **Proxy Integration:** Designed to integrate with multiple proxy instances, such as the [Generic Proxy](https://github.com/sajjadth/generic-proxy).
- **High Performance**: Built with Go for optimized performance and lower memory usage.
- **Proxy Integration**Designed to integrate with multiple proxy instances.
- **Environment Configuration**: Easily configure proxy instances through environment variables.
- **Zap Logging**: Built-in structured logging with Zap for better observability.
- **Custom Timeout & Transport**: Custom transport for optimized network usage and connection handling.

## Prerequisites

- Go (v1.22.4 or later)

* Git

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository_url>
   cd <repository_directory>
   ```
2. Create a .env file:

   - Copy the provided environment variables template and add your specific values.

   ```bash
    PORT=3000
    PROXY_INSTANCES=http://localhost:5001,http://localhost:5002

   ```

3. Install dependencies:
   ```
   go mod tidy
   ```
4. Run the application:
   ```
   go run main.go
   ```

## Usage

1. Start the load balancer:
   ```bash
   npm start
   ```
2. The load balancer will run on `http://localhost:8080` (or the port specified in your `.env` file) and will forward incoming traffic to the configured proxy instances in a round-robin manner.

## Integration with Generic Proxy

The Generic Proxy Load Balancer is designed to be used in conjunction with the [Generic Proxy](https://github.com/sajjadth/generic-proxy) project. Set up the Generic Proxy instances and configure the `PROXY_INSTANCES` environment variable to point to those instances.

## Migration from Node.js to Go

This project was originally implemented using Node.js and has now been migrated to Go for improved performance and memory efficiency.

The last version of the Node.js implementation can be found under the tag `v1.0-node`.

## Customization

You can modify `main.go` to add custom logic, such as:

- Changing request headers
- Adding authentication
- Logging requests

## Troubleshooting

- Ensure your `PROXY_INSTANCES` is correct in the `.env` file.
- Make sure the proxy instances are reachable from the server.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
