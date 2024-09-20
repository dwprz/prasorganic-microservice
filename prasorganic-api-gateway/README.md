
![Logo](https://ik.imagekit.io/pj3r6oe9k/prasorganic-high-resolution-logo-transparent.svg?updatedAt=1726835541390)
# Prasorganic API Gateway

Prasorganic API Gateway is one of the components in the Prasorganic microservices architecture,responsible for managing traffic between clients and various backend services. This gateway is configured using NGINX.


## Tech Stack

[![My Skills](https://skillicons.dev/icons?i=nginx,bash,git&theme=light)](https://skillicons.dev)

## Features

- **Proxy Server and Load Balancer:** The Prasorganic API Gateway uses NGINX as a proxy server, directing client requests to the appropriate server based on the internal domain configuration: restful.local for RESTful services and grpc.local for gRPC services. The load balancer is implemented using upstream to distribute requests across multiple service instances, thereby improving scalability.

- **Caching:** his gateway implements caching to reduce latency and improve performance. By temporarily storing responses from backend services, repeated requests can be processed faster without needing to be forwarded back to the server.

- **Rate Limiter:** To protect backend services from excessive traffic spikes, this gateway implements a rate limiter. This feature regulates the number of requests that can be processed within a given time frame, helping maintain system stability and reducing the risk of DOS attacks.

- 📝  _Actually, I wanted to create a separate load balancer for each service using NGINX, as it is common in microservices to run each service on multiple nodes. However, due to RAM limitations, I combined all the load balancers into one. As a result, any service that needs to communicate with another service uses this API gateway as the main route._

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Dependencies and Their Licenses

- `NGINX:` Licensed under the BSD 2-Clause License. For more information, see the [NGINX License](https://nginx.org/LICENSE).

- `Htpasswd:` Licensed under the Apache License, Version 2.0. For more information, see the [Htpasswd License](https://www.apache.org/licenses/LICENSE-2.0).

- `GNU Make:` Licensed under the GNU General Public License v3.0. For more information, see the [GNU Make License](https://www.gnu.org/licenses/gpl.html).

- `GNU Bash:` Licensed under the GNU General Public License v3.0. For more information, see the [Bash License](https://www.gnu.org/licenses/gpl-3.0.html).

- `Git:` Licensed under the GNU General Public License version 2.0. For more information, see the [Git License](https://opensource.org/license/GPL-2.0).

## Thanks 👍
Thank you for viewing my project.