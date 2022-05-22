# GoBlockchain
- Source Material: [tensor-programming](https://github.com/tensor-programming/)
- follow me on twitter [@terran_nft](https://twitter.com/terran_nft)

Go Lang Blockchain with cli, wallets, and transactions.

Use the guide found [here](https://www.youtube.com/watch?v=mYlHT9bB6OE)

The guide is 4 years old and hasn't been updated for a while, but after you go through it you should understand the principle mechanics and then reference the files here to get it up and running.

Setup:
    in terminal: 
    
    1 go build
    2 Go BlockChain.exe

Usage:
- getbalance -address ADDRESS - get the balance for an address
- createblockchain -address ADDRESS creates a blockchain and sends genesis reward to address
- printchain - Prints the blocks in the chain
- send -from FROM -to TO -amount AMOUNT -mine - Send amount of coins. Then -mine flag is set, mine off of this node
- createwallet - Creates a new Wallet
- listaddresses - Lists the addresses in our wallet file
- reindexutxo - Rebuilds the UTXO set
- startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining
