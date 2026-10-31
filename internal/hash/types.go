package hash

type Hasher struct {
	Time    uint32 `json:"time"`
	Memory  uint32 `json:"memory"`
	Threads uint8  `json:"threads"`
	KeyLen  uint32 `json:"keyLen"`
	SaltLen uint32 `json:"saltLen"`
}

type Hash struct {
	Hash []byte
	Salt []byte
}
