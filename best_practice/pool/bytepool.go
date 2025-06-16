package pool

type BytePool struct {
	c   chan []byte
	len int
	cap int
}

func NewBytePool(size, len, cap int) *BytePool {
	return &BytePool{
		c:   make(chan []byte, size),
		len: len,
		cap: cap,
	}
}

func (p *BytePool) Get() (b []byte) {
	select {
	case b = <-p.c:
	default:
		if p.cap > 0 {
			b = make([]byte, p.len, p.cap)
		} else {
			b = make([]byte, p.len)
		}
	}
	return
}

func (p *BytePool) Put(b []byte) {
	select {
	case p.c <- b:
	default:
		//discard
	}
}
