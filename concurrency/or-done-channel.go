package concurrency

func Process(c <-chan any, done <-chan struct{}) {
	for val := range orDone(c, done) {
		_ = val // todo: use val
	}
}

func orDone(c <-chan any, done <-chan struct{}) <-chan any {
	valStream := make(chan any)
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}

				select {
				case <-done:
					return
				case valStream <- v:
				}
			}
		}
	}()

	return valStream
}
