package proxy

import "io"

func transport(rw1, rw2 io.ReadWriter) error {
	errc := make(chan error, 1)
	go func() {
		buf := lPool.Get().([]byte)
		defer lPool.Put(buf)

		_, err := io.CopyBuffer(rw1, rw2, buf)
		errc <- err
	}()

	go func() {
		buf := lPool.Get().([]byte)
		defer lPool.Put(buf)

		_, err := io.CopyBuffer(rw2, rw1, buf)
		errc <- err
	}()

	err := <-errc
	if err != nil && err == io.EOF {
		err = nil
	}
	return err
}
