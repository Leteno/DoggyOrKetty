package saving

import (
       "testing"
)

func TestSaving(t *testing.T) {
     testKey := "hello"
     testValue := "world"
     waitForInit()
     value := GetPref(testKey)
     t.Logf("%s: %s", testKey, value)

     PutPref(testKey, testValue)
     t.Logf("%s: %s", testKey, GetPref(testKey))
     refreshPref()
     OnClose()
     t.Logf("after refresh: %s: %s", testKey, GetPref(testKey))
}