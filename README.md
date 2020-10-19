# Sales App Demo
The project focuses on a sales storage/retrieval business use case.<br>
Since there was a strong typed language requirement by the client,choices <br>
for Golang and Typescript were made.

## How it Works 

### Preview Table
The application consists of a preview table that shows paginated orders.<br>
Users can view and sort all orders using the controls here.

### Report Upload
Sample files (csv) are included in this repo under the prod/test-data folder.

### Date Filter

UI elements to the right of the screen show for a given date range:
#### The top five profitable item types.
#### The Total Profit.

## Background
### Front End
Application files for the frontend are found in the directory sfront.<br>
Main technologies involved here are the framework VueJs with the logic writter in typescript.<br>
Compiled assets for the client are available in the prod/dist folder.
Any decent web server should be able to serve these.

### Backend
Application files for the backend are found in the directory api.<br>
Main technologies involved here are the framework buffalo and the language Golang.<br>
compiled assets for the backend are availabel in the prod/bin folder.<br>
Compiled assets here are made into an executable, in this case this was a windows target machine.

# Limitations & Possible Solutions
## File Uploads are extremely slow for large files.
These uploads can reduce the quality of the user experience by causing timeouts.

### Use Of Web Workers (Front End)
Web workers and file chuncking techniques with (Browser File Api) can help prevent blocking the main thread.

### Use of workers in buffalo (Backend)
This will solve timeouts and the requirement for persistent connections.

### Use of Server Events
The client would simply subscribe to the server events. However this still requires a persistent connection.


