<div align="center" id="top"> 
  <img src="./.github/app.gif" alt="Hexagonal Arch Go" />

&#xa0;

  <!-- <a href="https://hexagonalarchgo.netlify.app">Demo</a> -->
</div>

<h1 align="center">Hexagonal Arch Go</h1>

<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/FirerPlayer/hexagonal-arch-go?color=56BEB8">

  <img alt="Github language count" src="https://img.shields.io/github/languages/count/FirerPlayer/hexagonal-arch-go?color=56BEB8">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/FirerPlayer/hexagonal-arch-go?color=56BEB8">

  <img alt="License" src="https://img.shields.io/github/license/FirerPlayer/hexagonal-arch-go?color=56BEB8">

  <!-- <img alt="Github issues" src="https://img.shields.io/github/issues/FirerPlayer/hexagonal-arch-go?color=56BEB8" /> -->

  <!-- <img alt="Github forks" src="https://img.shields.io/github/forks/FirerPlayer/hexagonal-arch-go?color=56BEB8" /> -->

  <!-- <img alt="Github stars" src="https://img.shields.io/github/stars/FirerPlayer/hexagonal-arch-go?color=56BEB8" /> -->
</p>

<!-- Status -->

<!-- <h4 align="center">
	🚧  Hexagonal Arch Go 🚀 Under construction...  🚧
</h4>

<hr> -->

<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#sparkles-features">Features</a> &#xa0; | &#xa0;
  <a href="#rocket-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Starting</a> &#xa0; | &#xa0;
  <a href="#memo-license">License</a> &#xa0; | &#xa0;
  <a href="https://github.com/FirerPlayer" target="_blank">Author</a>
</p>

<br>

## :dart: About

Sure! Here's a simplified version in English:

The Go project with hexagonal architecture involves HTTP requests and messaging to create a modern, scalable application. The hexagonal architecture separates business rules from technical implementation, allowing greater flexibility and easier code maintenance.

In this project, external adapters receive HTTP requests and pass the data to the application layer for business rules to be applied. The external adapter then formats the data and returns it as an HTTP response.

Additionally, messaging is used for asynchronous communication between different parts of the system. This allows for background task processing and efficient information exchange between different microservices.

Using Go for implementation offers high-performance, energy-efficient programming with advanced concurrency and parallelism features, making it ideal for high-load scenarios. The hexagonal architecture and messaging also make the system modular and scalable, making future maintenance and expansion easier.

The following is the structure of the project:

<div align="center">
	<img src="https://github.com/FirerPlayer/hexagonal-arch-go/blob/master/hex-arch.png" alt="Project architecture" />
	
</div>

## :sparkles: Features

:heavy_check_mark: Menssaging with Kafka\
:heavy_check_mark: Hexagonal architecture\
:heavy_check_mark: MySQL\
:heavy_check_mark: POO with Go

## :rocket: Technologies

The following tools were used in this project:

- [Go](https://golang.org)
- [Kafka](https://kafka.apache.org/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## :white_check_mark: Requirements

Before starting :checkered_flag:, you need to have [Git](https://git-scm.com) and [Node](https://nodejs.org/en/) installed.

## :checkered_flag: Starting

```bash
# Clone this project
$ git clone https://github.com/FirerPlayer/hexagonal-arch-go

# Access
$ cd hexagonal-arch-go

# Run the compose
$ docker-compose up -d
```

## :memo: License

This project is under license from MIT. For more details, see the [LICENSE](LICENSE.md) file.

Made with :heart: by <a href="https://github.com/FirerPlayer" target="_blank">Micael</a>

&#xa0;

<a href="#top">Back to top</a>
