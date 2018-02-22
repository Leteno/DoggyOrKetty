package saving

import (
       "sync"
       "log"
       _"time"
       "io/ioutil"
)

var prefSavingWg sync.WaitGroup
var prefInitWg sync.WaitGroup
var mutex sync.Mutex

var localMap map[string]string

var delimCh byte = '\n'
var delimStr = "\n"

func init() {
     localMap = make(map[string]string)
     refreshPref()
}

var prefPath = ".setting"

func refreshPref() {
     prefInitWg.Add(1)
     go func() {
	defer prefInitWg.Done()
     	mutex.Lock()
	defer mutex.Unlock()

     	for key := range localMap {
	    delete(localMap, key)
	}
	bytes, err := ioutil.ReadFile(prefPath)
	if err != nil {
	   log.Print(err)
	}
	var key string
	var value string
	var start = 0
	var isKey = true
	for i, ch := range bytes {
	    if ch == delimCh {
	       // TODO this method is ugly, maybe use queue can beautify it
	       if isKey {
	       	  key = string(bytes[start:i])
	       } else {
	       	 value = string(bytes[start:i])
		 localMap[key] = value
	       }
	       isKey = !isKey
	       start = i + 1
	    } else if  i == len(bytes) - 1 {
	       if isKey {
	       	  key = string(bytes[start:len(bytes)])
	       } else {
	       	 value = string(bytes[start:len(bytes)])
		 localMap[key] = value
	       }
	    }
	}
     }()
}

func PutPref(key string, value string) {
     waitForInit()
     localMap[key] = value
}

func GetPref(key string) string {
     waitForInit()
     return localMap[key]
}

var savingQueueCount = 0
func saveToFile() {
     if savingQueueCount > 10 {
     	// drop it
     	return
     }
     savingQueueCount++
     prefSavingWg.Add(1)
     mutex.Lock()
     defer mutex.Unlock()
     defer prefSavingWg.Done()
     // do saving thing in here
     var result string
     for key, value := range localMap {
	 result += key + delimStr + value + delimStr
     }
     err := ioutil.WriteFile(prefPath, []byte(result), 0666)
     if err != nil {
	log.Print(err)
     }
     savingQueueCount--
}

func waitForInit() {
     prefInitWg.Wait()
}

func onPrefClose() {
     log.Printf("pref save on close")
     saveToFile()
}