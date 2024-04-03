# Solana-Vanity

This program generates Solana wallets until it finds one that starts with a specified prefix. During the process, it provides updates on how many wallets have been generated. Once a matching wallet is found, it outputs the wallet's public and private keys.

## How to Run

1. **Install Go**: Ensure that Go is installed on your system. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

2. **Clone or Download the Program**: Obtain a copy of the program on your machine. If you have git installed, you can clone it using a git command. Alternatively, download the source code as a ZIP file and extract it.

3. **Navigate to the Program Directory**: Use a terminal (or command prompt) to navigate to the directory containing the program files.

4. **Run the Program**: In the terminal, run the command `go run main.go`. Make sure you're in the directory where `main.go` is located.

## Example Output

```
Target prefix: JONO
Status: 1000000 wallets generated in 8.6269127s
Status: 2000000 wallets generated in 16.4320478s
Status: 3000001 wallets generated in 26.6121211s
Success! Wallet found: JONOVzg...
Secret Key: 3Hj1q4jKj4f...
Attempts required: 3256789, Time elapsed: 28m15s
```

*Note: The actual output will vary, especially the wallet address, secret key, attempts required, and time elapsed. These are just examples to illustrate what the output might look like. The "Status" messages show intermediate updates on how many wallets have been generated and the time taken so far.*
