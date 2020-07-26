<h1 align="center">
      <img alt="Go_REST_API" title="Go_REST_API" src=".github/logo.png" width="300px" />
</h1>

<h3 align="center">
  Go Rest Api
</h3>

<p align="center">A simple API to test native http packages from Golang 🎓</p>
<p align="center">Made with Golang and PostgreSQL 🚀</p>
<p align="center">Using Docker 🐳</p>

<p align="center">
  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/Lgdev07/crud_api?color=%2304D361">

  <img alt="Made by Lgdev07" src="https://img.shields.io/badge/made%20by-Lgdev07-%2304D361">

  <img alt="License" src="https://img.shields.io/badge/license-MIT-%2304D361">

  <a href="https://github.com/Lgdev07/crud_api/stargazers">
    <img alt="Stargazers" src="https://img.shields.io/github/stars/Lgdev07/crud_api?style=social">
  </a>
</p>

<p align="center">
  <a href="#-installation-and-execution">Installation and execution</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-available-routes">Available Routes</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-contribute">How to contribute</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
</p>

## 🚀 Installation and execution

1. Clone this repository and go to the directory;
2. Rename .env.sample to .env;

<h4> 🔧 Development </h4>

1. Run `docker-compose up`;
2. Make the Requests;

<h4> 🧪 Tests </h4>

1. Run `docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit`;

## 🛣️ Available Routes

<h4> POST </h4>

- '/stores'

<h4> PUT </h4>

- '/stores/{id:[0-9]+}'

<h4> DELETE </h4>

- '/stores/{id:[0-9]+}'

<h4> GET </h4>

- '/stores'

- '/stores/{id:[0-9]+}'

## 🤔 How to contribute

- Fork this repository;
- Create a branch with your feature: `git checkout -b my-feature`;
- Commit your changes: `git commit -m 'feat: My new feature'`;
- Push to your branch: `git push origin my-feature`.

After the merge of your pull request is done, you can delete your branch.

---