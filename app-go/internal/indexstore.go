package internal

import "sync"

// Data structure that stores a document number and the number of time a word/term appears in the document
type DocFreqPair struct {
	DocumentNumber int64
	WordFrequency  int64
}

type IndexStore struct {
	// TO-DO declare data structure that keeps track of the DocumentMap
	DocumentMap map[int64]string

	// TO-DO declare data structures that keeps track of the TermInvertedIndex
	TermInvertedIndex map[string][]DocFreqPair

	// TO-DO declare two locks, one for the DocumentMap and one for the TermInvertedIndex
	DocMapMutex    sync.Mutex
	TermIndexMutex sync.Mutex
}

func NewIndexStore() *IndexStore {
	return &IndexStore{
		// TO-DO initialize the DocumentMap and TermInvertedIndex members and the locks
		DocumentMap:       make(map[int64]string),
		TermInvertedIndex: make(map[string][]DocFreqPair),
	}
}

func (self *IndexStore) PutDocument(documentPath string) int64 {
	var documentNumber int64
	// TO-DO assign a unique number to the document path and return the number
	// IMPORTANT! you need to make sure that only one thread at a time can access this method
	self.DocMapMutex.Lock()

	documentNumber = int64(len(self.DocumentMap))
	self.DocumentMap[documentNumber] = documentPath

	self.DocMapMutex.Unlock()

	return documentNumber
}

func (self *IndexStore) GetDocument(documentNumber int64) string {
	var documentPath string
	// TO-DO retrieve the document path that has the given document number
	self.DocMapMutex.Lock()
	documentPath = self.DocumentMap[documentNumber]
	self.DocMapMutex.Unlock()

	return documentPath
}

func (self *IndexStore) UpdateIndex(documentNumber int64, wordFrequencies map[string]int64) {
	// TO-DO update the TermInvertedIndex with the word frequencies of the specified document
	// IMPORTANT! you need to make sure that only one thread at a time can access this method
	self.TermIndexMutex.Lock()
	for key, value := range wordFrequencies {
		pair := DocFreqPair{documentNumber, value}

		self.TermInvertedIndex[key] = append(self.TermInvertedIndex[key], pair)
	}
	self.TermIndexMutex.Unlock()
}

func (self *IndexStore) LookupIndex(term string) []DocFreqPair {
	var results []DocFreqPair
	// TO-DO return the document and frequency pairs for the specified term
	self.TermIndexMutex.Lock()
	for _, value := range self.TermInvertedIndex[term] {
		results = append(results, value)
	}
	self.TermIndexMutex.Unlock()

	return results
}
