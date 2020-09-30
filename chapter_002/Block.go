package main

import(
	"bytes"
    "crypto/sha256"
    "encoding/binary"
    "time"
    "log"
)

type Block struct {
    TimeStamp   int64
    Hash        []byte
    PrevHash    []byte
    Height      int64
    Data        []byte
}


func NewBlock(height int64, prevBlockHash []byte, data []byte) *Block {
    var block Block
    block = Block{
        TimeStamp:time.Now().Unix(),
        Hash:nil,
        PrevHash:prevBlockHash,
        Height:height,
        Data:data,
    }
    
    // create the hash
    block.SetHash()
    
    return &block
}

func (b  *Block) SetHash() {
    // call Sha256
    timeStampBytes := IntToHex(b.TimeStamp)
    heightBytes := IntToHex(b.Height)
    
    blockBytes := bytes.Join([][]byte{
        heightBytes,
        timeStampBytes,
        b.PrevHash,
        b.Data,
    }, []byte{})
    
    hash := sha256.Sum256(blockBytes)
    b.Hash  = hash[:]
}

func CreateGenesisBlock(data []byte) *Block{
	return NewBlock(1, nil, data)
}

// convert int64 to []byte 
func IntToHex(data int64) []byte {
    buffer := new(bytes.Buffer)
    err := binary.Write(buffer, binary.BigEndian, data)
    if nil != err {
        log.Panic("in transact to []byte failed! %v\n", err)
    }
    return buffer.Bytes()
}
