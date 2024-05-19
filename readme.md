# GreenAPI Web Interface

## Description

This project is a simple web interface for interacting with the GreenAPI using various endpoints. It provides an easy-to-use interface for executing API methods and viewing their responses.

## Endpoints

The following endpoints are available:

1. **GET /api/getSettings**:

   - Description: Retrieves the settings of the GreenAPI instance.
   - Request Method: POST
   - Request Body:
     ```json
     {
       "idInstance": "string",
       "apiTokenInstance": "string"
     }
     ```
   - Response: Returns the settings of the GreenAPI instance.

2. **GET /api/getStateInstance**:

   - Description: Retrieves the state of the GreenAPI instance.
   - Request Method: POST
   - Request Body:
     ```json
     {
       "idInstance": "string",
       "apiTokenInstance": "string"
     }
     ```
   - Response: Returns the state of the GreenAPI instance.

3. **GET /api/sendMessage**:

   - Description: Sends a message using the GreenAPI instance.
   - Request Method: POST
   - Request Body:
     ```json
     {
       "idInstance": "string",
       "apiTokenInstance": "string",
       "chatId": "string",
       "message": "string"
     }
     ```
   - Response: Returns the result of sending the message.

4. **GET /api/sendFileByUrl**:
   - Description: Sends a file by URL using the GreenAPI instance.
   - Request Method: POST
   - Request Body:
     ```json
     {
       "idInstance": "string",
       "apiTokenInstance": "string",
       "chatId": "string",
       "urlFile": "string"
     }
     ```
   - Response: Returns the result of sending the file.

## URL

- https://test-api-form.onrender.com/index
