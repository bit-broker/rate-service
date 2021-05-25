<p align="center">
  <img alt="Bit-Broker" src="https://avatars.githubusercontent.com/u/80974981?s=200&u=7e396d371614d3a9ce7fc1f7fe4515e255374760&v=4" />
</p>

# Bit-Broker Rate Service

![Github Actions](https://github.com/bit-broker/rate-service/actions/workflows/docker-image.yml/badge.svg)

This repository contains the Rate Service used by Bit-Broker.

The Rate Service implements an HTTP REST API for the CRUD of the configuration and extends the gRPC v2 ratelimit proto (used by Ambassador).

## Deployment

It can be deployed using the following helm chart:

[Rate Service Helm Chart](https://github.com/bit-broker/k8s/tree/main/helm/charts/rate-service)

## Documentation

### REST API

#### Set or Update Configuration
----
  Adds a new configuration or updates an existing one with the unique identifier "UID".

* **URL**

  /api/v1/:uid/config

* **Method:**

  `PUT`

*  **URL Params**

   **Required:**

   `uid=[integer]`

* **Body**

   **Required:**

  ```json
  {
    "enabled": "true|false (Enable/Disable globally)",
    "rate": "N (Number of requests per second)",
    "quota": {
      "max_number": "N (Max requests)",
      "interval_type": "month|day (Per interval)"
    }
  }
  ```

* **Success Response:**

  * **Code:** 200 <br />

* **Error Response:**

  * **Code:** 404 <br />

* **Sample Call:**

  ```curl
  curl --location --request PUT '/api/v1/1/config' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "enabled": true,
    "rate":5,
    "quota":{
      "max_number": 20,
      "interval_type": "month"
    }
  }'
  ```

#### Get Configuration
----
  Returns the existing configuration with the unique identifier "UID".

* **URL**

  /api/v1/:uid/config

* **Method:**

  `GET`

*  **URL Params**

   **Required:**

   `uid=[integer]`

* **Success Response:**

  * **Code:** 200 <br />

* **Error Response:**

  * **Code:** 404 <br />

* **Sample Call:**

  ```curl
  curl --location --request GET '/api/v1/1/config'
  ```

#### Delete Configuration
----
  Deletes the existing configuration with the unique identifier "UID".

* **URL**

  /api/v1/:uid/config

* **Method:**

  `DELETE`

*  **URL Params**

   **Required:**

   `uid=[integer]`

* **Success Response:**

  * **Code:** 200 <br />

* **Error Response:**

  * **Code:** 404 <br />

* **Sample Call:**

  ```curl
  curl --location --request DELETE '/api/v1/1/config'
  ```

### gRPC Proto

[Envoy v2 RateLimit Proto](https://github.com/envoyproxy/envoy/blob/main/api/envoy/service/ratelimit/v2/rls.proto)
