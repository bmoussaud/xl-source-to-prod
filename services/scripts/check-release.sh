export release_id="Applications/Folder371e232af7e04f778a82f2843709dc93/Release47d1a4a73630443ebe99c34c8c6ef637"

http --auth admin:admin  GET "http://localhost:5516/api/v1/templates/${release_id}" 'Accept:application/json' | jq ".status"



