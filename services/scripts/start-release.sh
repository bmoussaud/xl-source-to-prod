echo '{
  "releaseTitle" : "xl-demo-uccm-app-canary",
  "folderId" : "Applications/Folder371e232af7e04f778a82f2843709dc93",
  "variables" : {
    "application" : "xl-demo-uccm-app",
    "stage":"xl-demo-production"
  },
  "releaseVariables" : {
    "application" : "xl-demo-uccm-app",
    "stage":"xl-demo-production"
  },
  "scheduledStartDate" : null,
  "autoStart" : true
}' | http --auth admin:admin POST 'http://localhost:5516/api/v1/templates/Release7d25ce2f75224fae947f1bedd1f0440e/create' \
    'Accept:application/json' \
    'Content-Type:application/json' \
   | jq ".id"