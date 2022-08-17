package fizzbuzz

import "sync"

// RequestCounter keeps track of the most frequent requests received by the server
type RequestCounter interface {
	NewRequest(request any)
	MostFrequentRequest() ([]any, int)
}

// RequestCounterImpl implements RequestCounter interface
type RequestCounterImpl struct {
	counter             map[any]int
	highestCount        int
	mostFrequentRequest []any
	mutext              sync.RWMutex
}

// NewRequestCounter initiates new instance of RequestCounter
func NewRequestCounter() RequestCounter {
	return &RequestCounterImpl{
		counter:      make(map[any]int),
		highestCount: -1,
	}
}

// NewRequest registers new request received by the server
func (requestCounter *RequestCounterImpl) NewRequest(request any) {
	requestCounter.mutext.Lock()
	defer requestCounter.mutext.Unlock()

	requestCounter.counter[request]++
	if requestCounter.counter[request] == requestCounter.highestCount {
		requestCounter.mostFrequentRequest = append(requestCounter.mostFrequentRequest, request)
	} else if requestCounter.counter[request] > requestCounter.highestCount {
		requestCounter.highestCount = requestCounter.counter[request]
		requestCounter.mostFrequentRequest = []any{request}
	}
}

// MostFrequentRequest returns the most frequent requests
func (requestCounter *RequestCounterImpl) MostFrequentRequest() ([]any, int) {
	requestCounter.mutext.RLock()
	defer requestCounter.mutext.RUnlock()

	return requestCounter.mostFrequentRequest, requestCounter.highestCount
}
