package main

import (
    "code.google.com/p/gcfg"
)

type Config struct {
        Gerrit struct {
            User     string
            Password string
            Host     string
            CI       string
        }
        Robot struct {
            Interval string
            Project  []string
            Profile  string
        }
}

func ReadConfig(file string) (*Config, error) {
    var config Config
    err := gcfg.ReadFileInto(&config, file)
    if err != nil {
        return nil, err
    }

    return &config, nil
}
