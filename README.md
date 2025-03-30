# go-moderation-service
Golang bases service that scans user-generated content (comments and reviews) for inappropriate content and flag it for moderation.

# Prerequisites

Ensure that Docker and Docker Compose are installed on your machine.

# Setting up go-moderation-service

Follow these steps to set up the application:

1.Clone the repository:
```
git clone https://github.com/sourav014/go-moderation-service.git
```
2. Set Up Environment Variables

Before starting the application, configure the required environment variables. These can be set inside the docker-compose.yaml file.
Google Cloud Credentials

- Download the Service Account JSON Key from the Google Cloud Console.
- Save it as gcp_credentials_file.json.
- DO NOT commit this file to the repository.
- The file should be mounted when running the Docker container.

SendGrid API Key for Email Notifications

- Set up the following environment variables:
```
SENDGRID_API_KEY: "your-sendgrid-api-key"
SENDGRID_FROM_EMAIL_ADDRESS: "your-email@example.com"
```
- Download the Service Account JSON file from Google Cloud Console and save it inside project root directory as gcp_credentials_file.json. [Reference: https://github.com/sourav014/go-moderation-service/blob/main/gcp_credentials_file.json]

3.Start the application using Docker Compose:
```
docker-compose up -d
```
Alternatively, if you're using a newer version of Docker:
```
docker compose up -d
```
4.Access the application: After the setup is complete, the application will be running on port 8080.

All the API Endpoints are available in the below postman collection URL
```
https://www.postman.com/souravprasad/content-moderation-service/collection/mip37wz/content-moderation-service?action=share&creator=15932967
```
