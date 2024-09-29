# Generic Proxy Load Balancer

The **Generic Proxy Load Balancer** is a Node.js-based load balancer that distributes incoming traffic across multiple proxy instances, providing a simple and effective way to manage load distribution. It is designed to work seamlessly with the [Generic NodeJS Proxy](https://github.com/sajjadth/generic-nodejs-proxy) project to create a scalable proxy solution.

## Features

- **Round-Robin Load Balancing:** Distributes requests evenly across all proxy instances using a round-robin algorithm.
- **Proxy Integration:** Designed to integrate with multiple proxy instances, such as the [Generic NodeJS Proxy](https://github.com/sajjadth/generic-nodejs-proxy).
- **Environment Configuration:** Easily configure proxy instances through environment variables.
- **Automatic Path Rewriting:** Optionally rewrites paths for proxied requests.

## Prerequisites

- Node.js (v14.x or later)
- NPM (Node Package Manager)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/sajjadth/generic-proxy-load-balancer.git
   cd generic-proxy-load-balancer
2. nstall dependencies:
    ```bash
    npm install
3. Create a .env file in the root of your project and configure your environment variables:
    ```bash
    PORT=3000
    PROXY_INSTANCES=http://localhost:5001,http://localhost:5002
    
- `PORT`: The port on which the load balancer will run (default is 3000 if not specified).
- `PROXY_INSTANCES`: A comma-separated list of proxy instance URLs.

## Usage

1. Start the load balancer:
    ```bash
    npm start
2. The load balancer will run on `http://localhost:3000` (or the port specified in your `.env` file) and will forward incoming traffic to the configured proxy instances in a round-robin manner.

## Integration with Generic NodeJS Proxy

The Generic Proxy Load Balancer is designed to be used in conjunction with the [Generic NodeJS Proxy](https://github.com/sajjadth/generic-nodejs-proxy) project. Set up the Generic NodeJS Proxy instances and configure the `PROXY_INSTANCES` environment variable to point to those instances.


## Customization

You can modify `index.js` to add custom logic, such as:
- Changing request headers
- Adding authentication
- Logging requests

## Troubleshooting

- Ensure your `TARGET_URL` is correct in the `.env` file.
- Make sure the target API is reachable from your server.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
