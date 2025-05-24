package main

func crypt(textCh, keyCh <-chan byte, result chan<- byte) {
	// Use defer to ensure result channel is closed when function exits
	defer close(result)

	var text, key byte
	var textOk, keyOk bool

	for {
		text, textOk = <-textCh
		if !textOk {
			// Text channel is closed, we're done
			return
		}

		key, keyOk = <-keyCh
		if !keyOk {
			// Key channel is closed, we're done
			return
		}

		result <- text ^ key
	}
}
