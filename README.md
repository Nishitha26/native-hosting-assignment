# native-hosting-assignment
# Overview
This project is a Go-based backend service created for hosting a static site that lets users to upload a zip file of a static site, Automatically extract and serve it, Track deployment metadata, Preview deployed sites via generated URLs
## Endpoints
- POST /deploy :
  - Upload zip file and site_name
  - Body: form-data
  - file: (File upload)
  - site_name: (Text)
  - Sample Response:
    {
    "message": "Deployment successful",
    "preview_url": "http://localhost:8080/sites/test-site/",
    "site": "test-site"
}
- GET /deployments – List all deployments
- GET /hello-world – Simple health check
- GET /sites/{site_name}/ – Access deployed site
## How to Run
- git clone https://github.com/Nishitha26/native-hosting-assignment
-go to the project folder
-go mod tidy(to install dependencies)
-go run ./cmd/main.go(Start the development server)
## Improvements With More Time
-Add rollback functionality and versioned deployments
-Store deployments in SQLite
-Add frontend dashboard or Swagger UI
-Add site name conflict checks and zip validation
-Allow site deletion via API
