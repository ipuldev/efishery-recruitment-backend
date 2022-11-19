<a name="readme-top"></a>
<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li><a href="#prerequisites">Prerequisites</a></li>
    <li><a href="#installation">Installation</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

Efishery commodities is provide authentication & commodities data service

Detail Structure Project :
* Auth Service
  * Database:
    * SQLite
  * Routing:
    * Login
    * Authorize
    * Register
* Fetch Service
  * Routing:
    * Commodities
    * Aggregate

More Detail about architecture program:
_Figma [Figma][Design-url]_
<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Go][Go.com]][Go-url]
* [![NodeJs][Node.com]][Node-url]
* [![Docker][Docker.com]][Docker-url]


<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Prerequisites
* Docker Engine 20.10.18
* Docker Compose version 2.11.2

### Installation
1. ```docker compose up -d```
<br />
Note make sure your docker version is competible
<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage
Testing module
  - Auth service
    - `npm install`
    - `npm test`

  - Fetch service
    - `go test./... -v`

For more information 

_Postman [Documentation](https://postman.com/belivine/workspace/efishery-backend)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] Auth Service
- [x] Fetch Service
- [x] Environment
    - [x] Routing
    - [x] Docker Compose & Image
    - [x] Middleware Auth Forward
- [x] Unit Testing
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[linkedin-url]: https://www.linkedin.com/in/muhammad-saiful-abdulah-079545186/
[Design-url]:https://www.figma.com/file/oVnEMhG7Y6BAq5VXbHYT4r/efishery-reqruitment-design-backend?node-id=0%3A1&t=n5IwguYSyDga205K-1
[GO.com]:https://img.shields.io/badge/GO%201.19-0769AD?style=flat&logo=go&logoColor=white
[GO-Url]:https://go.dev/
[Docker.com]:https://img.shields.io/badge/docker-003f8c?style=flat&logo=docker&logoColor=white
[Docker-Url]:https://www.docker.com/
[Node.com]:https://img.shields.io/badge/node.js-74b858?style=flat&logo=javascript&logoColor=white
[Node-Url]:https://nodejs.org/en/
