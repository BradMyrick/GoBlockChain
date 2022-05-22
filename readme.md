Dev: Brad Myrick 
Source Material: https://github.com/tensor-programming/
follow me on twitter @terran_nft
Go Lang Blockchain with cli, wallets, and transactions.
Use the guide found at https://www.youtube.com/watch?v=mYlHT9bB6OE
then use the Badger DB and various updates I've included in the code to get it running correctly.
the guide is 4 years old and has not been updated.
Check back for chain updates. 

Setup:
    in terminal:  go build .
    then:         GoBlockChain.exe

Usage:
 getbalance -address ADDRESS - get the balance for an address
 createblockchain -address ADDRESS creates a blockchain and sends genesis reward to address
 printchain - Prints the blocks in the chain
 send -from FROM -to TO -amount AMOUNT -mine - Send amount of coins. Then -mine flag is set, mine off of this node
 createwallet - Creates a new Wallet
 listaddresses - Lists the addresses in our wallet file
 reindexutxo - Rebuilds the UTXO set
 startnode -miner ADDRESS - Start a node with ID specified in NODE_ID env. var. -miner enables mining
