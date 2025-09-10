package main

import "fmt"

type Config struct {
    Timeout int
    Mode    string
}

type configManager struct {
    config     Config
    setChan    chan Config
    getChan    chan chan Config
}

func newConfigManager(initial Config) *configManager {
    cm := &configManager{
        config:  initial,
        setChan: make(chan Config),
        getChan: make(chan chan Config),
    }
    go cm.run()
    return cm
}

func (cm *configManager) run() {
    for {
        select {
        case newConfig := <-cm.setChan:
            cm.config = newConfig
        case respChan := <-cm.getChan:
            respChan <- cm.config
        }
    }
}

func (cm *configManager) Set(config Config) {
    cm.setChan <- config
}

func (cm *configManager) Get() Config {
    respChan := make(chan Config)
    cm.getChan <- respChan
    return <-respChan
}

func main() {
    manager := newConfigManager(Config{Timeout: 30, Mode: "production"})
    manager.Set(Config{Timeout: 20, Mode: "debug"}) // меняешь 

    config := manager.Get()
    fmt.Println("задержка и режим", config.Timeout, config.Mode)
}