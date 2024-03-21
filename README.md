# HumanizerGPT

This project is a Essay/Code generator but has 0% AI detection. Tested with ZeroGPT, Quilbot AI detection, etc. ***DISCLAIMER: The author does not endorse the use of this app for any unethical purposes***

Link to hackathon submission: https://devpost.com/software/humanizergpt

## Tech Stack

- **Frontend**: Vue.js
- **Backend**: Go

## Project Structure

The project is divided into two main directories:

- `frontend`: This directory contains all the Vue.js code for the frontend of the application.
- `main`: This directory contains the Go code for the backend of the application.

## APIs

The backend server is built with Go and it serves static files from the `frontend/dist` directory. It uses the `godotenv` package to load environment variables from a `.env` file.

## How to Run

1. First, you need to build the frontend. Navigate to the `frontend` directory and run `yarn install` to install the dependencies. Then, run `yarn build` to build the frontend.
2. Next, navigate to the `main` directory. If you haven't done so already, run `go get` to download the Go dependencies.
3. Run `go run main/server.go` to start the backend server. The server will serve the static files built in step 1.

Please make sure to update the `.env` file with your specific settings before running the application.


## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
