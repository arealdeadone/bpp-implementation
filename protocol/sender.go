package protocol

import (
	"bpp/network"
	"log"
)

func SendQKey(length uint64) {

	log.Printf("Generating a key of length: %d\r\n", length)
	key := GenerateRandomKey(length)
	log.Printf("The key generated is : %s\r\n", key)

	log.Println("Selecting Random Basis for the key")
	basis := GenerateRandomBasis(length)
	log.Printf("The basis selected are : %s\r\n", basis)

	log.Println("Encoding the generated trits to qutrits")
	encoded := TritsToQutrits(key, basis)
	log.Printf("The encoded qutrits are %s", encoded)

	network.SendMessage(":8002", encoded)
	network.SendMessage(":8002", "STOP")

}
