package tcp

import (
	"encoding/binary"
	"log"
	"net"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/eliothedeman/bangarang/event"
	"github.com/eliothedeman/bangarang/provider"
)

func init() {
	provider.LoadEventProviderFactory("tcp", NewTCPProvider)
}

// provides events from tcp connections
type TCPProvider struct {
	encoding string
	pool     *event.EncodingPool
	laddr    *net.TCPAddr
	listener *net.TCPListener
}

func NewTCPProvider() provider.EventProvider {
	return &TCPProvider{}
}

// the config struct for the tcp provider
type TCPConfig struct {
	Encoding    string `json:"encoding"`
	Listen      string `json:"listen"`
	MaxDecoders int    `json:"max_decoders"`
}

func (t *TCPProvider) Init(i interface{}) error {
	c := i.(*TCPConfig)

	// make sure we have a valid address
	addr, err := net.ResolveTCPAddr("tcp4", c.Listen)
	if err != nil {
		return err
	}

	t.laddr = addr

	// start listening on that addr
	err = t.listen()
	if err != nil {
		return err
	}

	// build an encoding pool
	t.pool = event.NewEncodingPool(event.EncoderFactories[c.Encoding], event.DecoderFactories[c.Encoding], c.MaxDecoders)
	return nil
}

func (t *TCPProvider) ConfigStruct() interface{} {
	return &TCPConfig{
		Encoding:    event.ENCODING_TYPE_JSON,
		MaxDecoders: runtime.NumCPU(),
	}
}

// start accepting connections and consume each of them as they come in
func (t *TCPProvider) Start(dst chan *event.Event) {

	// listen for ever
	for {
		c, err := t.listener.AcceptTCP()
		if err != nil {
			log.Println(err)
		} else {
			// consume the connection
			t.consume(c, dst)
		}
	}
}

func (t *TCPProvider) consume(conn *net.TCPConn, dst chan *event.Event) {
	buff := make([]byte, 1024*200)
	var size_buff = make([]byte, 8)
	var e *event.Event
	var nextEventSize uint64
	var n int
	var err error
	for {

		// read the size of the next event
		n, err = conn.Read(size_buff)
		if err != nil {
			logrus.Error(err)
			conn.Close()
			return
		}

		if n != 8 {
			logrus.Errorf("tcp-provider: Expecting 8byte 64bit unsigned int. Only got %d bytes", n)
			conn.Close()
			return
		}

		nextEventSize, _ = binary.Uvarint(size_buff)
		logrus.Debugf("Next event from tcp provider is %d bytes", nextEventSize)

		// read the next event
		n, err = conn.Read(buff[:int(nextEventSize)])
		if err != nil {
			logrus.Error(err)
			conn.Close()
			return
		}

		logrus.Debugf("New event from tcp provider: %s", string(buff[:int(nextEventSize)]))

		t.pool.Decode(func(d event.Decoder) {
			e, err = d.Decode(buff[:nextEventSize])
		})

		if err != nil {
			log.Println(err)
		} else {
			dst <- e
		}
	}
}

func (t *TCPProvider) listen() error {
	l, err := net.ListenTCP("tcp", t.laddr)
	if err != nil {
		return err
	}

	t.listener = l
	return nil
}
