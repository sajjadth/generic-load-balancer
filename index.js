const express = require("express");
const { createProxyMiddleware } = require("http-proxy-middleware");
require("dotenv").config();

const app = express();

// List of your deployment URLs (proxy instances)
const proxyInstances = process.env.PROXY_INSTANCES.split(",");

let currentInstanceIndex = 0;

// Function to get the next proxy instance
function getNextProxyInstance() {
  const instance = proxyInstances[currentInstanceIndex];
  currentInstanceIndex = (currentInstanceIndex + 1) % proxyInstances.length;
  return instance;
}

// Middleware to handle load balancing to different proxy instances
app.use("/", (req, res, next) => {
  const targetUrl = getNextProxyInstance();
  console.log(`Forwarding request to proxy instance: ${targetUrl}${req.url}`);

  createProxyMiddleware({
    target: targetUrl,
    changeOrigin: true,
    pathRewrite: {
      "^/": "/", // Rewrite the path if needed
    },
    onProxyReq: (proxyReq, req, res) => {
      console.log(
        `Proxying request to: ${targetUrl}${req.url} (Method: ${
          req.method
        }, Headers: ${JSON.stringify(req.headers)})`
      );
    },
  })(req, res, next);
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`Load balancer is running on port ${PORT}`);
});
